package models
import (
    "fmt"
    "net/http"
)
type Servise struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Cost string `json:"cost"`
}
type ServiseList struct {
    Servises []Servise `json:"Servises"`
}
func (i *Servise) Bind(r *http.Request) error {
    if i.Name == "" {
        return fmt.Errorf("name is a required field")
    }
    return nil
}
func (*ServiseList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (*Servise) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}