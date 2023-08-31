package db

import (
	"database/sql"
	"github.com/EORUG/avitotest/models"
)

func (db Database) AddSegment(segment *models.Segment) (*models.Segment, error) {
	var id int
	query := `INSERT INTO SEGMENTS (segmentName) VALUES ($1) RETURNING segmentID`
	err := db.Conn.QueryRow(query, segment.SegmentName).Scan(&id)
	if err != nil {
		return nil, err
	}
	segment.SegmentID = id
	return segment, nil
}

func (db Database) DeleteSegment(name string) error {
	_, err := db.Conn.Exec("DELETE FROM SEGMENTS WHERE segmentName like $1", name)
	return err
}

func (db Database) GetSegmentByName(segmentName string) (int, error) {
	segment := models.Segment{}
	query := `SELECT * FROM SEGMENTS WHERE segmentName LIKE $1;`
	row := db.Conn.QueryRow(query, segmentName)
	switch err := row.Scan(&segment.SegmentID, &segment.SegmentName); err {
	case sql.ErrNoRows:
		return segment.SegmentID, ErrNoMatch
	default:
		return segment.SegmentID, err
	}
}
