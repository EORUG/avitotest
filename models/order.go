package models
import (
    "fmt"
    "net/http"
)
type Order struct {
    ID int `json:"id"`
    Userid string `json:"userid"`
    Serviseid string `json:"servisesid"`
    CreatedAt string `json:"created_at"`
}
type OrderList struct {
    Orders []Order `json:"Orders"`
}
func (i *Order) Bind(r *http.Request) error {
    if i.Userid == "" {
        return fmt.Errorf("name is a required field")
    }
    return nil
}
func (*OrderList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (*Order) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}