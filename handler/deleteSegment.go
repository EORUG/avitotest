package handler

import (
	"github.com/EORUG/avitotest/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func DeleteSegment(router chi.Router) {
	router.Route("/DeleteSegment", func(router chi.Router) {
		router.Post("/", DeleteSegments)
	})
}

type DeleteSegmentRequest struct {
	Name string `json:"name"`
}

func (i *DeleteSegmentRequest) Bind(r *http.Request) error {

	return nil
}
func (*DeleteSegmentRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func DeleteSegments(w http.ResponseWriter, r *http.Request) {
	deleteSegmentRequest := &DeleteSegmentRequest{}
	if err := render.Bind(r, deleteSegmentRequest); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	err := dbInstance.DeleteSegment(deleteSegmentRequest.Name)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, &models.Status{
		Success: true,
	}); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

}
