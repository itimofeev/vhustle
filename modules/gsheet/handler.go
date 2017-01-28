package gsheet

import (
	"encoding/json"
	"github.com/itimofeev/vhustle/modules/util"
	"net/http"
)

func HandleGetContestList(c *util.Context) {
	contests, err := listContests(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, convertContestToDto(contests))
}

func convertContestToDto(contestDbs []ContestDb) []ContestDto {
	res := make([]ContestDto, 0, len(contestDbs))

	for _, db := range contestDbs {
		res = append(res, ContestDto{
			ID:           db.ID,
			Title:        db.Title,
			Date:         db.Date,
			DateStr:      db.DateStr,
			CityName:     db.CityName,
			ForumURL:     db.ForumURL,
			VkLink:       db.VkLink,
			PreregLink:   db.PreregLink,
			CommonInfo:   db.CommonInfo,
			ResultsLink:  convertLinksDbToDto(db.ResultsLink),
			VideosLink:   convertLinksDbToDto(db.VideosLink),
			PhotosLink:   convertLinksDbToDto(db.PhotosLink),
			AvatarFile:   db.AvatarFile,
			UpdateDate:   db.UpdateDate,
			LastSyncDate: db.LastSyncDate,
		})
	}

	return res
}

func convertLinksDbToDto(jsonLinks string) []TextLink {
	res := make([]TextLink, 0)
	err := json.Unmarshal([]byte(jsonLinks), &res)
	if err != nil {
		util.ContestGLog.WithError(err).WithField("jsonLinks", jsonLinks).Error("Unable to unmarshall json")
	}
	return res
}
