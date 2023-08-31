package models

import (
	"net/http"
)

type Segment struct {
	SegmentID   int    `json:"id"`
	SegmentName string `json:"segmentName"`
}
type SegmentList struct {
	Segments []Segment `json:"Segments"`
}

func (i *Segment) Bind(r *http.Request) error {

	return nil
}
func (*SegmentList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Segment) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
