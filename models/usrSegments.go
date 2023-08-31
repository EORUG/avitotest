package models

import (
	"net/http"
)

type UsrSegment struct {
	UsrID     int     `json:"usrID"`
	SegmentID int     `json:"segmentID"`
	TTL       *string `json:"TTL"`
}
type UsrSegmentList struct {
	UsrSegments []UsrSegment `json:"UsrSegments"`
}

func (i *UsrSegment) Bind(r *http.Request) error {

	return nil
}
func (*UsrSegmentList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*UsrSegment) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
