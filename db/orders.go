package db
import (
    "database/sql"
    "github.com/EORUG/avitotest/models"
)
func (db Database) GetOrderById(orderId int) (models.Order, error) {
    order := models.Order{}
    query := `SELECT * FROM orders WHERE id = $1;`
    row := db.Conn.QueryRow(query, orderId)
    switch err := row.Scan(&order.ID); err {
    case sql.ErrNoRows:
        return order, ErrNoMatch
    default:
        return order, err
    }
}