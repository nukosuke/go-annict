package annict

import (
	"net/http"
	"time"
)

type RecordsService service

type Record struct {
	Id         int64     `json:"id"`
	Comment    string    `json:"comment"`
	Rating     float32   `json:"rating"`
	IsModified bool      `json:"is_modified"`
	LikesCount int64     `json:"likes_count"`
	CreatedAt  time.Time `json:"created_at"`

	User struct {
		Id           int64     `json:"id"`
		Username     string    `json:"username"`
		Name         string    `json:"name"`
		Description  string    `json:"description"`
		Url          string    `json:"url"`
		RecordsCount int64     `json:"recourds_count"`
		CreatedAt    time.Time `json:"created_at"`
	} `json:"user"`

	Work Work `json:"work"`

	Episode struct {
		Id           int64  `json:"id"`
		Number       string `json:"number"`
		NumberText   string `json:"number_text"`
		SortNumber   int64  `json:"sort_number"`
		Title        string `json:"title"`
		RecordsCount int64  `json:"records_count"`
	} `json:"episode"`
}

type RecordList struct {
	Records []Record `json:"records"`
	*Pagination
}

type RecordsListOptions struct {
}

func (s *RecordsService) List(opt *RecordsListOptions) (*RecordList, *http.Response, error) {
	u, err := addOptions("v1/records", opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	records := new(*RecordList)
	resp, err := s.client.Do(req, records)
	if err != nil {
		return nil, resp, err
	}

	return *records, resp, err
}
