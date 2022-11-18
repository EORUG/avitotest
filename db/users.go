package db
import (
    "database/sql"
    "github.com/EORUG/avitotest/models"
)
func (db Database) GetAllItems() (*models.UserList, error) {
    list := &models.UserList{}
    rows, err := db.Conn.Query("SELECT * FROM users ORDER BY ID DESC")
    if err != nil {
        return list, err
    }
    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.Cash, &user.Reserve)
        if err != nil {
            return list, err
        }
        list.Users = append(list.Users, user)
    }
    return list, nil
}
func (db Database) GetItemById(userId int) (models.user, error) {
    user := models.User{}
    query := `SELECT * FROM users WHERE id = $1;`
    row := db.Conn.QueryRow(query, userId)
    switch err := row.Scan(&user.ID, &user.Cash, &user.Reserve); err {
    case sql.ErrNoRows:
        return user, ErrNoMatch
    default:
        return user, err
    }
}
func (db Database) UpdateItem(userId int, userData models.User) (models.User, error) {
    user := models.User{}
    query := `UPDATE users SET cash=$1, reserve=$2 WHERE id=$3 RETURNING id, cash, reserve;`
    err := db.Conn.QueryRow(query, userData.Cash, userData.Reserve, userId).Scan(&user.ID, &user.Cash, &user.Reserve)
    if err != nil {
        if err == sql.ErrNoRows {
            return user, ErrNoMatch
        }
        return user, err
    }
    return user, nil
}