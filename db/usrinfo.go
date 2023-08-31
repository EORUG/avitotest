package db

import (
	"github.com/EORUG/avitotest/models"
)

func (db Database) GetUsrById(userId int) (models.Usr, error) {
	type row struct {
		UserID    int     `json:"userID"`
		SegmentID int     `json:"segmentID"`
		TTL       *string `json:"TTL"`
	}

	user := models.Usr{}
	query := `SELECT * FROM user_segments WHERE userID = $1;`
	rows, err := db.Conn.Query(query, userId)
	defer rows.Close()
	if err != nil {
		return models.Usr{}, err
	}
	user.ID = userId
	for rows.Next() {
		el := row{}
		err := rows.Scan(&el.UserID, &el.SegmentID, &el.TTL)
		if err != nil {
			return models.Usr{}, err
		}
		user.SegmentID = append(user.SegmentID, el.SegmentID)

	}
	return user, err
}
