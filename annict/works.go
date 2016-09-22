package annict

import (
	"net/http"
	"time"
)

type WorksService service

type Work struct {
	Id              int64     `json:"id"`
	Title           string    `json:"title"`
	TitleKana       string    `json:"title_kana"`
	Media           string    `json:"media"`
	MediaText       string    `json:"media_text"`
	SeasonName      string    `json:"season_name"`
	SeasonNameText  string    `json:"season_name_text"`
	ReleasedOn      time.Time `json:"released_on"`
	ReleasedOnAbout time.Time `json:"released_on_about"`
	OfficialSiteUrl string    `json:"official_site_url"`
	WikipediaUrl    string    `json:"wikipedia_url"`
	TwitterUsername string    `json:"twitter_username"`
	TwitterHashtag  string    `json:"twitter_hashtag"`
	EpisodesCount   int       `json:"episodes_count"`
	WatchersCount   int       `json:"watchers_count"`
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

func (s *WorksService) List(opt *WorksListOptions) ([]Work, *http.Response, error) {
	u, err := addOptions("works", opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	works := new([]Work)
	resp, err := s.client.Do(req, works)
	if err != nil {
		return nil, resp, err
	}

	return *works, resp, err
}
