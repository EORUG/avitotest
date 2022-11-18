package db
import (
    "database/sql"
    "github.com/EORUG/avitotest/models"
)
func (db Database) GetItemById(serviseId int) (models.Servise, error) {
    servise := models.Servise{}
    query := `SELECT * FROM servises WHERE id = $1;`
    row := db.Conn.QueryRow(query, serviseId)
    switch err := row.Scan(&servise.ID, &servise.Name, &servise.Cost); err {
    case sql.ErrNoRows:
        return servise, ErrNoMatch
    default:
        return servise, err
    }
}