package annict

import (
  "net/http"
)

type MeStatusesService service

type MeStatusesListOptions struct {
	WorkId int64  `url:"work_id"`
	Kind   string `url:"kind"`
}

func (s *MeStatusesService) Change(opt *MeStatusesListOptions) (*http.Response, error) {
	u, err := addOptions("v1/me/statuses", opt)
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
