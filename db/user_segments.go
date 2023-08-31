package db

import "github.com/EORUG/avitotest/models"

func (db Database) AddUserSegment(usrsegment *models.UsrSegment) error {
	query := `INSERT INTO USER_SEGMENTS (userID,segmentID,TTL) VALUES ($1,$2,$3)`
	_, err := db.Conn.Exec(query, usrsegment.UsrID, usrsegment.SegmentID, usrsegment.TTL)
	return err
}

func (db Database) DeleteUserSegment(usrid int, segmentid int) error {
	query := `DELETE FROM USER_SEGMENTS  WHERE userID =$1 AND segmentID=$2 `
	_, err := db.Conn.Exec(query, usrid, segmentid)
	return err
}

func (db Database) TTLControl() error {
	query := `DELETE FROM USER_SEGMENTS  WHERE TTL <=now()`
	_, err := db.Conn.Exec(query)
	return err
}
