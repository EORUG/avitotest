package db

import "github.com/EORUG/avitotest/models"

func (db Database) GetUsrlogs(MMYY string, usrID int) (models.LogList, error) {

	logout := models.LogList{}
	query := `SELECT * FROM LOG WHERE datatime > $1 AND userID = $2;`
	rows, err := db.Conn.Query(query, MMYY, usrID)
	defer rows.Close()
	if err != nil {
		return models.LogList{}, err
	}
	for rows.Next() {
		el := models.Log{}
		err := rows.Scan(&el)
		if err != nil {
			return models.LogList{}, err
		}
		logout.Logs = append(logout.Logs, el)

	}
	return logout, err
}
