package gsheet

import (
	"context"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/itimofeev/vhustle/modules/util"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"time"
)

const readRange = "Contests!A2:K"

const (
	colTitle = iota
	colDate
	colDateStr
	colCityName
	colForumUrl
	colVkLink
	colPreregLink
	colCommonInfo
	colResultsLink
	colVideosLink
	colPhotosLink
)

func fetchContests() (c []ContestG, err error) {
	util.ContestGLog.Debug("Started to fetch contests from gsheet")

	values, err := getValues()
	if err != nil {
		util.ContestGLog.WithError(err).Error("Got error while fetching contests from gsheet")
		return
	}
	if len(values) == 0 {
		err = errors.New("Values is empty")
		return
	}

	for _, row := range values {
		c = append(c, ContestG{
			Title:      rowVal(row, colTitle),
			Date:       parseGDate(rowVal(row, colDate)),
			DateStr:    rowVal(row, colDateStr),
			CityName:   rowVal(row, colCityName),
			ForumURL:   rowVal(row, colForumUrl),
			VkLink:     rowVal(row, colVkLink),
			PreregLink: rowVal(row, colPreregLink),
			CommonInfo: rowVal(row, colCommonInfo),
			ResultLink: rowVal(row, colResultsLink),
			VideoLink:  rowVal(row, colVideosLink),
			PhotoLink:  rowVal(row, colPhotosLink),
		})
	}
	util.ContestGLog.WithField("rowsFetched", len(values)).Debug("Finished fetching contests from gsheet")

	return
}

func rowVal(row []interface{}, col int) string {
	if len(row) <= col {
		return ""
	}
	return row[col].(string)
}

func parseGDate(date string) time.Time {
	t, err := time.Parse("02/01/2006", date)
	if err != nil {
		util.AnyLog.WithField("date", date).WithError(err).Error("Fail in parse date from G sheet")
	}

	return t
}

const gSheetJSON = `
{
  "type": "service_account",
  "project_id": "vhustle-151505",
  "private_key_id": "acd6c7760750b485e944dc80f6002ea9f6cc88bd",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEwAIBADANBgkqhkiG9w0BAQEFAASCBKowggSmAgEAAoIBAQDFJtjj//U3ElnQ\nFx7A49nt4oMmt7WK/0MaVz8EQ0spZGHtCNJXyVXX8G7DpUZBoDRZPzSO+fL071yA\nqx/KaiA5+Y3IXyHGG5CKtiVKCOtO0N+AYumMqnRV4Gavl1qCrSB30tz2+zY2Z2l4\nQd0rp8M+VJ0ytVb9kwTaeC/JwLxYs/Z2rEjU5A1sZnaXEKKtcov+bfIh2GAWB31Z\n/DMAGqxP4ZDW3G0xbeZtjq4ho3iKS3kKgfArv6dz38jqxifR2xqhnAUvHRueF3RZ\n9sWWW5wrLEW4vB9ufGJY4j1s8N7/OBWgL9a8YAqYHiLah1/1Dp5Bj18JxdHBaVHn\ninYdNr7ZAgMBAAECggEBAIH6PY37puga8hlt1LmovnnGF19ESK0N42iPUp113Cy6\n4JDMexijRTQrcGsOIIaNn1WjhPwqL5Jp6Ftv9nKViw+NxnrutS6N57p7oZPw02nP\n7ToQfBdgHXisjCaBq4txpnE5FLLEJhayEOfWzIDGhsMmN8lBostkzRNXn3Hs3n+/\nY9RRiFH4fMof+9toy+gXg3ktyE5Pqh3QyeXG4fA+6K9icdTeK24dH4fycEF/usST\nviFJXbcP2qPhL82SRDic4metYogtveR0ZciBPUpY+7eCHwsrM09lK9lMinqxCfD2\ncDJjFwutiWUh1r8Ysf+ZOhHfQtbVRJBewGaPdWZW3j0CgYEA+w+jaL0ZVD+BXnBx\n6WgUhtRPRGJdVMjwORvC77fqzYf0ejlnTZWSXDI84Sx6o/mJNarsjnag7CvbMHor\nbuQ36NBNdy2bn93cA2vjYu6z1nKbw5FiBP6Xtw8V0/DnNA3qsJXYw1l3qnIjakJv\nmoM4tttT9XXTvPyoyKcyOi8qCZMCgYEAyQe3tJnH5geXuGatR4w+wbdvR3DnZz9y\nPX29mT8d5PDxoy7bLeLlByc9cPSFCfuvBUeREv9qbkIiry/u1yDn1EzwPWe9TSrr\n3Iu0ENPgr+uLOdsGbZWjan75Bb4FYtwEgTt/dS3mRxf4q/fopkKEd5tDtYYIfjAN\nJbXHQBuxqWMCgYEA+mQEY6eFJYMYsWmQEmtdXYNNczRvROoKu8o2RwK0yTt41pV9\ns+Ei0ZTBJwpHXla3Q7EusH8by6+JsfWGb9ho8mcde2kfNvf8P+VQKRFMhupS81+B\n1N1dzpLbAD/ZNw9SK7+nKl2GfZXMQGP2DrIk1Co2uC5FeMy8QTKPY3w6fsUCgYEA\noG7GTx1DCPMaRBG9TBJCqzqHVk5mfmGF/EjzHO/gHaukTATC6pXfDZxlTlX2Lodr\noB0DTFQGTkP9hi4MwCcEnMtiFr9JteIBJZtgcuWEtSCXle6T7LS1G/KFLe3+Qm0w\nMyqWh5+/3RDmZeeNBdKkAQgqOx3ifUI/2858W08+s90CgYEA4Xj7WC/Uts6qG/Ld\nQ/aDevK6A3oPynj5OpSbtSOmM0TWO8OTZ1kjWt+3sImmvp99UbDtnErLJYu5+YaP\nzGAT1vWbDb0xno2JkwXPjZKe72+8waTh/5SQDReCnpdzSYuzszIDlLi39REMQFS9\n1Waii5Gs2VWqJm97O/tUvmr+2ME=\n-----END PRIVATE KEY-----\n",
  "client_email": "vhustle-sync@vhustle-151505.iam.gserviceaccount.com",
  "client_id": "114910332783718693157",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://accounts.google.com/o/oauth2/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/vhustle-sync%40vhustle-151505.iam.gserviceaccount.com"
}
`

func getValues() (values [][]interface{}, err error) {
	ctx := context.Background()

	config, err := google.JWTConfigFromJSON([]byte(gSheetJSON), "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		err = errors.New(fmt.Sprintf("Unable to parse client secret file to config: %v", err))
		return
	}

	client := config.Client(ctx)

	srv, err := sheets.New(client)
	if err != nil {
		err = errors.New(fmt.Sprintf("Unable to retrieve Sheets Client %v", err))
		return
	}

	spreadsheetId := "13L13i-0VH_8NH508rVJm6ohcChtAb4XuuZz5umkCT2Y"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		err = errors.New(fmt.Sprintf("Unable to retrieve data from sheet. %v", err))
		return
	}
	values = resp.Values
	return
}
