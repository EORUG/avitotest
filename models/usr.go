package models

import (
	"net/http"
)

type Usr struct {
	ID        int   `json:"id"`
	SegmentID []int `json:"segmentID"`
}
type UsrList struct {
	Usrs []Usr `json:"Usrs"`
}

func (i *Usr) Bind(r *http.Request) error {

	return nil
}
func (*UsrList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Usr) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
