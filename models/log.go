package models

import (
	"net/http"
)

type Log struct {
	UsrID       int    `json:"usrID"`
	SegmentName string `json:"segmentName"`
	Operation   bool   `json:"operation"`
	Datetime    string `json:"datetime"`
}
type LogList struct {
	Logs []Log `json:"logs"`
}

func (i *Log) Bind(r *http.Request) error {

	return nil
}
func (*LogList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Log) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
