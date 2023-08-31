package handler

import (
	"net/http"

	"github.com/EORUG/avitotest/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func ChangeSegment(router chi.Router) {
	router.Route("/ChangeSegment", func(router chi.Router) {
		router.Post("/", ChangeSegments)
	})
}

type ChangeSegmentRequest struct {
	ToAdd    []string `json:"toAdd"`
	ToDelete []string `json:"toDelete"`
	UserID   int      `json:"userId"`
	TTL      *string  `json:"TTL"`
}

func (i *ChangeSegmentRequest) Bind(r *http.Request) error {

	return nil
}
func (*ChangeSegmentRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ChangeSegments(w http.ResponseWriter, r *http.Request) {
	createSegmentRequest := &ChangeSegmentRequest{}
	if err := render.Bind(r, createSegmentRequest); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	for _, s := range createSegmentRequest.ToDelete {
		segid, err := dbInstance.GetSegmentByName(s)
		if err != nil {
			render.Render(w, r, ErrorRenderer(err))
			return
		}
		err = dbInstance.DeleteUserSegment(createSegmentRequest.UserID, segid)
		if err != nil {
			render.Render(w, r, ErrorRenderer(err))
			return
		}
	}
	for _, s := range createSegmentRequest.ToAdd {
		segid, err := dbInstance.GetSegmentByName(s)
		if err != nil {
			render.Render(w, r, ErrorRenderer(err))
			return
		}
		toadd := models.UsrSegment{SegmentID: segid, UsrID: createSegmentRequest.UserID, TTL: createSegmentRequest.TTL}
		err = dbInstance.AddUserSegment(&toadd)
		if err != nil {
			render.Render(w, r, ErrorRenderer(err))
			return
		}
	}
	if err := render.Render(w, r, &models.Status{
		Success: true,
	}); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
