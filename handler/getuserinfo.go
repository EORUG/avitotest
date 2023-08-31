package handler

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func GetUserInfo(router chi.Router) {
	router.Route("/GetUserInfo", func(router chi.Router) {
		router.Post("/", GetUserInfos)
	})
}

type GetUserInfoRequest struct {
	UsrID int `json:"id"`
}

func (i *GetUserInfoRequest) Bind(r *http.Request) error {

	return nil
}
func (*GetUserInfoRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func GetUserInfos(w http.ResponseWriter, r *http.Request) {
	getuserinfoRequest := &GetUserInfoRequest{}
	if err := render.Bind(r, getuserinfoRequest); err != nil {
		render.Render(w, r, ErrBadRequest)
		fmt.Println(err.Error())
		return
	}
	usersegments, err := dbInstance.GetUsrById(getuserinfoRequest.UsrID)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		fmt.Println(err.Error())
		return
	}
	if err := render.Render(w, r, &usersegments); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		fmt.Println(err.Error())
		return
	}

}
