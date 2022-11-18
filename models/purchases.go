package models
import (
    
    "net/http"
)
type Purchase struct {
    ID int `json:"id"`
    Userid int `json:"userid"`
    Serviseid int `json:"servisesid"`
	Orderid int `json:"orderid"`
    Paid bool `json:"paid"`
    CreatedAt string `json:"created_at"`
}
type PurchaseList struct {
    Purchases []Purchase `json:"Purchases"`
}
func (i *Purchase) Bind(r *http.Request) error {
    
    return nil
}
func (*PurchaseList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (*Purchase) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}