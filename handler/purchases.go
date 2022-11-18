package handler

import (
	"context"
	"net/http"

	"github.com/EORUG/avitotest/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func purchases(router chi.Router) {
	router.Route("/Purchase", func(router chi.Router) {
		router.Post("/", PurchaseReserve)
		router.Put("/", PurchaseRevenue)
	})
}

type PurchaseRequest struct {
	UserID    int `json:"userid"`
	Serviseid int `json:"servisesid"`
	Orderid   int `json:"orderid"`
	Cash      int `json:"cash"`
}

func (i *PurchaseRequest) Bind(r *http.Request) error {

	return nil
}
func (*PurchaseRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func PurchaseReserve(w http.ResponseWriter, r *http.Request) {
	purchaseRequest := &PurchaseRequest{}
	if err := render.Bind(r, purchaseRequest); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	user, err := dbInstance.GetUserById(purchaseRequest.UserID)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if user.Cash < purchaseRequest.Cash {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	tx, err := dbInstance.Conn.BeginTx(context.Background(), nil)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	defer tx.Rollback()
	purchase, err := dbInstance.AddPurchase(&models.Purchase{
		Userid:    purchaseRequest.UserID,
		Serviseid: purchaseRequest.Serviseid,
		Orderid:   purchaseRequest.Orderid,
	})
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	user.Cash -= purchaseRequest.Cash
	user.Reserve += purchaseRequest.Cash
	user, err = dbInstance.UpdateUser(user.ID, user)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err = tx.Commit(); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, purchase); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

}
func PurchaseRevenue(w http.ResponseWriter, r *http.Request) {
	purchaseRequest := &PurchaseRequest{}
	if err := render.Bind(r, purchaseRequest); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	user, err := dbInstance.GetUserById(purchaseRequest.UserID)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if user.Reserve < purchaseRequest.Cash {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	purchase, err := dbInstance.GetPurchaseByUserServiseOrder(purchaseRequest.UserID, purchaseRequest.Orderid, purchaseRequest.Serviseid)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	tx, err := dbInstance.Conn.BeginTx(context.Background(), nil)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	defer tx.Rollback()
	user.Reserve -= purchaseRequest.Cash
	purchase.Paid = true
	user, err = dbInstance.UpdateUser(user.ID, user)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	purchase, err = dbInstance.UpdatePurchase(purchase.ID, purchase)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err = tx.Commit(); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, &purchase); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
