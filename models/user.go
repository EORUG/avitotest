package models
import (
    "fmt"
    "net/http"
)
type User struct {
    ID int `json:"id"`
    Cash int `json:"cash"`
    Reserve int `json:"reserve"`
}
type UserList struct {
    User []User `json:"Users"`
}
func (i *User) Bind(r *http.Request) error {
    if i.Cash == "" {
        return fmt.Errorf("name is a required field")
    }
    return nil
}
func (*UserList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (*User) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}