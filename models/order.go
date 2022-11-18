package models
import (
    
    "net/http"
)
type Order struct {
    ID int `json:"id"`
    Userid string `json:"userid"`
    Serviseid string `json:"serviseid"`
    Paid bool `json:"paid"`
    CreatedAt string `json:"created_at"`
}
type OrderList struct {
    Orders []Order `json:"Orders"`
}
func (i *Order) Bind(r *http.Request) error {
    
    return nil
}
func (*OrderList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (*Order) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}