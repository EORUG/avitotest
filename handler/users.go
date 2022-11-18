package handler

import (
	"net/http"
	"strconv"

	"github.com/EORUG/avitotest/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func users(router chi.Router) {
	router.Route("/{id}", func(router chi.Router) {
		router.Get("/", getUserCash)
		router.Put("/", UserCashAccrual)
	})
}
func getUserCash(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}
	user, err := dbInstance.GetUserById(userID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &user); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

type UserCashRequest struct {
	Cash int `json:"cash"`
}

func (i *UserCashRequest) Bind(r *http.Request) error {
	return nil
}
func UserCashAccrual(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	userCashRequest := &UserCashRequest{}
	if err := render.Bind(r, userCashRequest); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	user, err := dbInstance.GetUserById(userID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	user.Cash += userCashRequest.Cash
	user, err = dbInstance.UpdateUser(userID, user)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &user); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	return
}
