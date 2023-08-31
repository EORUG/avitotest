package handler

import (
	"math/rand"
	"net/http"

	"github.com/EORUG/avitotest/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func CreateSegment(router chi.Router) {
	router.Route("/CreateSegment", func(router chi.Router) {
		router.Post("/", CreateSegments)
	})
}

type CreateSegmentRequest struct {
	Name    string `json:"name"`
	Persent *int   `json:"persent"`
}

func (i *CreateSegmentRequest) Bind(r *http.Request) error {

	return nil
}
func (*CreateSegmentRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func CreateSegments(w http.ResponseWriter, r *http.Request) {
	createSegmentRequest := &CreateSegmentRequest{}
	if err := render.Bind(r, createSegmentRequest); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	segment, err := dbInstance.AddSegment(&models.Segment{
		SegmentName: createSegmentRequest.Name,
	})
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if (createSegmentRequest.Persent == nil) || (*createSegmentRequest.Persent == 0) {
		if err := render.Render(w, r, segment); err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
	}
	userIDs, err := dbInstance.GetAllUserIds()
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	for _, userID := range userIDs {
		if rand.Intn(100) > *createSegmentRequest.Persent {
			continue
		}
		toadd := models.UsrSegment{SegmentID: segment.SegmentID, UsrID: int(userID)}
		err = dbInstance.AddUserSegment(&toadd)
		if err != nil {
			render.Render(w, r, ErrorRenderer(err))
			return
		}
	}
	if err := render.Render(w, r, segment); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

}
