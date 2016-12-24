package gsheet

import (
	"time"
)

var contestInsert = []string{"title", "date", "date_str", "city_name", "forum_url", "vk_link", "prereg_link", "common_info", "results_link", "videos_link", "photos_link", "update_date", "last_sync_date"}

type ContestDb struct {
	ID       int64     `db:"id"`
	Title    string    `db:"title"`
	Date     time.Time `db:"date"`
	DateStr  string    `db:"date_str"`
	CityName string    `db:"city_name"`
	ForumURL string    `db:"forum_url"`

	VkLink     string `db:"vk_link"`
	PreregLink string `db:"prereg_link"`

	CommonInfo string `db:"common_info"`

	ResultsLink string `db:"results_link"`
	VideosLink  string `db:"videos_link"`
	PhotosLink  string `db:"photos_link"`

	UpdateDate   time.Time `db:"update_date"`
	LastSyncDate time.Time `db:"last_sync_date"`
}

type ContestG struct {
	Title    string
	Date     time.Time
	DateStr  string
	CityName string
	ForumURL string

	VkLink     string
	PreregLink string

	CommonInfo string

	ResultLink string
	VideoLink  string
	PhotoLink  string
}

type TextLink struct {
	Text string `json:"text"`
	Link string `json:"link"`
}

type ContestDto struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	Date     time.Time `json:"date"`
	DateStr  string    `json:"date_str"`
	CityName string    `json:"city_name"`
	ForumURL string    `json:"forum_url"`

	VkLink     string `json:"vk_link"`
	PreregLink string `json:"prereg_link"`

	CommonInfo string `json:"common_info"`

	ResultsLink []TextLink `json:"results_link"`
	VideosLink  []TextLink `json:"videos_link"`
	PhotosLink  []TextLink `json:"photos_link"`

	UpdateDate   time.Time `json:"update_date"`
	LastSyncDate time.Time `json:"last_sync_date"`
}
