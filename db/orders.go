package db
import (
    "database/sql"
    "github.com/EORUG/avitotest/models"
)
func (db Database) AddItem(order *models.Order) error {
    var id int
    var userid int
    var servisesid int
    var createdAt string
    query := `INSERT INTO orders (userid, serviseid) VALUES ($1, $2) RETURNING id, created_at`
    err := db.Conn.QueryRow(query, order.Userid, order.Serviseid).Scan(&id, &createdAt)
    if err != nil {
        return err
    }
    order.ID = id
    order.CreatedAt = createdAt
    return nil
}
func (db Database) GetItemById(orderId int) (models.Order, error) {
    order := models.Order{}
    query := `SELECT * FROM orders WHERE id = $1;`
    row := db.Conn.QueryRow(query, orderId)
    switch err := row.Scan(&order.ID, &order.Userid, &order.Serviseid, &order.CreatedAt); err {
    case sql.ErrNoRows:
        return order, ErrNoMatch
    default:
        return order, err
    }
}
func (db Database) DeleteItem(orderId int) error {
    query := `DELETE FROM orders WHERE id = $1;`
    _, err := db.Conn.Exec(query, orderId)
    switch err {
    case sql.ErrNoRows:
        return ErrNoMatch
    default:
        return err
    }
}