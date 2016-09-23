package annict

import (
	"net/http"
)

type EpisodesService service

type EpisodeProps struct {
	Id           int64  `json:"id"`
	Number       string `json:"number"`
	NumberText   string `json:"number_text"`
	SortNumber   int64  `json:"sort_number"`
	Title        string `json:"title"`
	RecordsCount int64  `json:"recourds_count"`
}

type Episode struct {
	*EpisodeProps
	Work        Work         `json:"work"`
	PrevEpisode EpisodeProps `json:"prev_episode"`
	NextEpisode EpisodeProps `json:"next_episode"`
}

type EpisodeList struct {
	Episodes []Episode `json:"episodes"`
	*Pagination
}

type EpisodesListOptions struct {
}

func (s *EpisodesService) List(opt *EpisodesListOptions) (*EpisodeList, *http.Response, error) {
	u, err := addOptions("v1/episodes", opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	episodes := new(*EpisodeList)
	resp, err := s.client.Do(req, episodes)
	if err != nil {
		return nil, resp, err
	}

	return *episodes, resp, err
}
