package models

import (
	"net/http"
)

type Status struct {
	Success bool `json:"success"`
}
type StatusList struct {
	Status []Status `json:"status"`
}

func (i *Status) Bind(r *http.Request) error {

	return nil
}
func (*StatusList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Status) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
