package services

import (
	"database/sql"

	"appstack.xyz/keeper_rest_api/exceptions"
	"appstack.xyz/keeper_rest_api/models"
)

type UserService struct {
	MysqlDB *sql.DB
}

const table = "users_tbl"

func NewUserService(db *sql.DB) *UserService {
	return &UserService{MysqlDB: db}
}

func (service *UserService) FetchByID(userID string) (*models.UserEntity, error) {

	user := &models.UserEntity{}

	const query = `SELECT id, mobile 
    FROM users_tbl WHERE id = ?`
	//log.Println("get by mobile ", user)

	stmt, err := service.MysqlDB.Prepare(query)
	if err != nil {
		return nil, exceptions.ErrorServer
	}
	defer stmt.Close()
	row := stmt.QueryRow(userID)
	err = row.Scan(&user.Id, user.Mobile)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exceptions.ErrUserNotFound
		}

	}

	return user, nil
}

func (service *UserService) InsertUser(user *models.UserEntity) (*models.UserEntity, error) {
	const query = `INSERT INTO users_tbl 
    (id, mobile) 
    VALUES (?,?)`

	stmt, err := service.MysqlDB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, exceptions.ErrorServer
	}

	_, err = stmt.Exec(user.Id, user.Mobile)
	if err != nil {
		return nil, exceptions.ErrorServer
	}
	return user, nil
}
