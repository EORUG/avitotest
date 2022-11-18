package db

import (
	"database/sql"

	"github.com/EORUG/avitotest/models"
)

func (db Database) AddPurchase(purchase *models.Purchase) (*models.Purchase, error) {
	var id int
	var createdAt string
	query := `INSERT INTO purchases (userid, serviseid, orderid) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, purchase.Userid, purchase.Serviseid, purchase.Orderid).Scan(&id, &createdAt)
	if err != nil {
		return nil, err
	}
	purchase.ID = id
	purchase.CreatedAt = createdAt
	return purchase, nil
}
func (db Database) GetPurchaseById(purchaseId int) (models.Purchase, error) {
	purchase := models.Purchase{}
	query := `SELECT * FROM purchases WHERE id = $1;`
	row := db.Conn.QueryRow(query, purchaseId)
	switch err := row.Scan(&purchase.ID, &purchase.Userid, &purchase.Serviseid, &purchase.Orderid, &purchase.Paid, &purchase.CreatedAt); err {
	case sql.ErrNoRows:
		return purchase, ErrNoMatch
	default:
		return purchase, err
	}
}
func (db Database) GetPurchaseByUserServiseOrder(userId int, orderId int, serviseId int) (models.Purchase, error) {
	purchase := models.Purchase{}
	query := `SELECT * FROM purchases WHERE userid = $1 AND serviseid = $2 AND orderid = $3 AND paid= false;`
	row := db.Conn.QueryRow(query, userId, serviseId, orderId)
	switch err := row.Scan(&purchase.ID, &purchase.Userid, &purchase.Serviseid, &purchase.Orderid, &purchase.Paid, &purchase.CreatedAt); err {
	case sql.ErrNoRows:
		return purchase, ErrNoMatch
	default:
		return purchase, err
	}
}
func (db Database) UpdatePurchase(purchaseId int, purchaseData models.Purchase) (models.Purchase, error) {
	query := `UPDATE purchases SET paid=$1 WHERE id=$2 RETURNING id, paid;`
	err := db.Conn.QueryRow(query, purchaseData.Paid, purchaseId).Scan(&purchaseData.ID, &purchaseData.Paid)
	if err != nil {
		if err == sql.ErrNoRows {
			return purchaseData, ErrNoMatch
		}
		return purchaseData, err
	}
	return purchaseData, nil
}
