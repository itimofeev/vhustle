package gsheet

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/itimofeev/vhustle/modules/util"
	"strings"
	"time"
)

func UpdateContestsCron() func() {
	return func() {
		UpdateContests()
	}
}

func UpdateContests() {
	contestsG, err := fetchContests()
	if err != nil {
		util.ContestGLog.WithError(err).Error("Error while fetching contests")
	}
	updateDbInfo(contestsG)
}

func updateDbInfo(contestsG []ContestG) error {
	util.ContestGLog.WithField("contestGLen", len(contestsG)).Debug("Update contestsG in db")
	for _, contestG := range contestsG {
		contestDb, err := fetchByForumUrl(contestG.ForumURL)
		if err != nil {
			util.ContestGLog.WithField("forumUrl", contestG.ForumURL).WithError(err).Error("Got error while fetching from db")
			continue
		}
		if contestDb == nil {
			util.ContestGLog.WithField("contestDb", contestDb).Debug("Contest not exists in db")
			err := insertContestDb(contestG)
			if err != nil {
				util.ContestGLog.WithFields(logrus.Fields{
					"contestG": contestG,
				}).WithError(err).Debug("Not created in db")
			}
			continue
		}
		if areContestsEqual(contestG, *contestDb) {
			util.ContestGLog.WithFields(logrus.Fields{
				"contestG":  contestG,
				"contestDb": contestDb,
			}).Debug("Contests are equal, update only last sync date")

			err = updateLastSyncContest(contestDb)
			if err != nil {
				util.ContestGLog.WithFields(logrus.Fields{
					"contestG": contestG,
				}).WithError(err).Debug("Not updated last sync in db")
			}
			continue
		}

		util.ContestGLog.WithFields(logrus.Fields{
			"contestG":  contestG,
			"contestDb": contestDb,
		}).Debug("Contests are not equal, updating in db")

		err = updateContest(*contestDb, contestG)
		if err != nil {
			util.ContestGLog.WithFields(logrus.Fields{
				"contestG":  contestG,
				"contestDb": contestDb,
			}).WithError(err).Debug("Not updated in db")
		}
	}
	return nil
}

func fetchByForumUrl(forumUrl string) (ci *ContestDb, err error) {
	var contestDb []ContestDb
	err = util.DB.SQL(`
		SELECT
			c.*
		FROM
			contest c
		WHERE
			c.forum_url = $1
	`, forumUrl).QueryStructs(&contestDb)

	if err != nil {
		return
	}

	if len(contestDb) == 0 {
		return
	}

	if len(contestDb) > 1 {
		err = errors.New(fmt.Sprintf("Found more than one infos: %+v", contestDb))
		return
	}

	ci = &contestDb[0]
	return
}

func areContestsEqual(contestG ContestG, contestDb ContestDb) bool {
	converted := convertContestToDb(contestG)

	return contestDb.CityName == converted.CityName &&
		contestDb.Title == converted.Title &&
		contestDb.DateStr == converted.DateStr &&
		contestDb.CommonInfo == converted.CommonInfo &&
		contestDb.ForumURL == converted.ForumURL &&
		contestDb.PhotosLink == converted.PhotosLink &&
		contestDb.PreregLink == converted.PreregLink &&
		contestDb.ResultsLink == converted.ResultsLink &&
		contestDb.VideosLink == converted.VideosLink &&
		contestDb.VkLink == converted.VkLink
}

func convertContestToDb(contestG ContestG) ContestDb {
	return ContestDb{
		Title:       contestG.Title,
		Date:        contestG.Date,
		DateStr:     contestG.DateStr,
		CityName:    contestG.CityName,
		ForumURL:    contestG.ForumURL,
		VkLink:      contestG.VkLink,
		PreregLink:  contestG.PreregLink,
		CommonInfo:  contestG.CommonInfo,
		ResultsLink: convertLinksDb(contestG.ResultLink),
		VideosLink:  convertLinksDb(contestG.VideoLink),
		PhotosLink:  convertLinksDb(contestG.PhotoLink),
	}
}

func convertLinksDb(str string) string {
	textLinks, err := convertLinksDto(str)
	if err != nil {
		util.ContestGLog.WithField("str", str).WithError(err).Error("Unable to convert link to DTO")
		return ""
	}
	bytes, err := json.Marshal(textLinks)
	if err != nil {
		util.ContestGLog.WithField("str", str).
			WithField("textLinks", textLinks).
			WithError(err).
			Error("Unable to marshall textLinks")
		return ""
	}
	return string(bytes)
}

func convertLinksDto(str string) ([]TextLink, error) {
	res := make([]TextLink, 0)
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if !strings.Contains(line, ":") {
			return res, errors.New(fmt.Sprintf("Unable to parse line '%s'", line))
		}
		split := strings.SplitN(line, ":", 2)
		res = append(res, TextLink{
			Text: strings.TrimSpace(split[0]),
			Link: strings.TrimSpace(split[1]),
		})
	}
	return res, nil
}

func listContests(c *util.Context) (contests []ContestDb, err error) {
	contests = make([]ContestDb, 0)
	err = util.DB.SQL(`
	    SELECT
	        c.*
	    FROM
	        contest c
	    ORDER BY
	        c.date`,
	).QueryStructs(&contests)
	return
}

func insertContestDb(contestG ContestG) error {
	now := time.Now()
	contestDb := convertContestToDb(contestG)
	contestDb.UpdateDate = now
	contestDb.LastSyncDate = now
	_, err := util.DB.
		InsertInto("contest").
		Columns(contestInsert...).
		Record(contestDb).
		Exec()
	return err
}

func updateContest(existedContestDb ContestDb, contestG ContestG) error {
	now := time.Now()
	contestDb := convertContestToDb(contestG)
	contestDb.ID = existedContestDb.ID
	contestDb.UpdateDate = now
	contestDb.LastSyncDate = now

	_, err := util.DB.
		Update("contest").
		SetBlacklist(contestDb, "forum_url", "id").
		Where("forum_url = $1", contestDb.ForumURL).
		Exec()

	return err
}

func updateLastSyncContest(contestDb *ContestDb) error {
	_, err := util.DB.
		Update("contest").
		Set("last_sync_date", time.Now()).
		Where("forum_url = $1", contestDb.ForumURL).
		Exec()

	return err
}
