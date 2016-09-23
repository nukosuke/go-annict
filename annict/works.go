package annict

import (
	"net/http"
	_ "time"
)

type WorksService service

type Work struct {
	Id             int64  `json:"id"`
	Title          string `json:"title"`
	TitleKana      string `json:"title_kana"`
	Media          string `json:"media"`
	MediaText      string `json:"media_text"`
	SeasonName     string `json:"season_name"`
	SeasonNameText string `json:"season_name_text"`

	// TODO: time.TimeのunmarshalがデフォルトでRFC3339なのでparseできない
	// 自前で実装する必要がある
	ReleasedOn      string `json:"released_on"`
	ReleasedOnAbout string `json:"released_on_about"`

	OfficialSiteUrl string `json:"official_site_url"`
	WikipediaUrl    string `json:"wikipedia_url"`
	TwitterUsername string `json:"twitter_username"`
	TwitterHashtag  string `json:"twitter_hashtag"`
	EpisodesCount   int    `json:"episodes_count"`
	WatchersCount   int64  `json:"watchers_count"`
}

type WorkList struct {
	Works []Work `json:"works"`
	*Pagination
}

type WorksListOptions struct {
	Fields            []string
	FilterIds         []int64
	FilterSeason      string
	FilterTitle       string
	Page              int64
	PerPage           int64
	SortId            string
	SortSeason        string
	SortWatchersCount string
}

func (s *WorksService) List(opt *WorksListOptions) (*WorkList, *http.Response, error) {
	u, err := addOptions("v1/works", opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	works := new(*WorkList)
	resp, err := s.client.Do(req, works)
	if err != nil {
		return nil, resp, err
	}

	return *works, resp, err
}
