package annict

import (
	"net/http"
)

type MeRecordsService service

type MeRecordsListOptions struct {
	EpisodeId     int64   `url:"episode_id"`
	Comment       string  `url:"comment,omitempty"`
	Rating        float32 `url:"rating,omitempty"`
	ShareTwitter  bool    `url:"share_twitter,omitempty"`
	ShareFacebook bool    `url:"share_facebook,omitempty"`
}

func (s *MeRecordsService) Create(opt *MeRecordsListOptions) (*http.Response, error) {
	u, err := addOptions("v1/me/records", opt)
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
