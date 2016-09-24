package annict

import (
	"net/http"
)

type MeWorksService service

type MeWorksListOptions struct {
	Fields            []string `url:"fields,omitempty"`
	FilterIds         []int64  `url:"filter_ids,omitempty"`
	FilterSeason      string   `url:"filter_season,omitempty"`
	FilterTitle       string   `url:"filter_title,omitempty"`
	FilterStatus      string   `url:"filter_status,omitempty"`
	Page              int64    `url:"page,omitempty"`
	PerPage           int64    `url:"per_page,omitempty"`
	SortId            string   `url:"sort_id,omitempty"`
	SortSeason        string   `url:"sort_season,omitempty"`
	SortWatchersCount string   `url:"sort_watchers_count,omitempty"`
}

func (s *MeWorksService) List(opt *MeWorksListOptions) (*WorkList, *http.Response, error) {
	u, err := addOptions("v1/me/works", opt)
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
