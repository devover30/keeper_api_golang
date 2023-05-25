package services

import (
	"database/sql"
	"log"

	"time"

	"appstack.xyz/keeper_rest_api/exceptions"
	"appstack.xyz/keeper_rest_api/models"
	"github.com/google/uuid"
)

type CredentialService struct {
	MysqlDB *sql.DB
}

const tableName = "credentials_tbl"

func NewCredentialService(db *sql.DB) *CredentialService {
	return &CredentialService{MysqlDB: db}
}

func (service *CredentialService) PersistCredential(cred *models.CredentialRequestDTO, userReq *models.UserEntity) (*models.CredentialEntity, error) {

	userService := NewUserService(service.MysqlDB)

	user, err := userService.FetchByID(userReq.Id)

	if err != nil {
		if err != exceptions.ErrUserNotFound {
			return nil, err
		}
		user, err = userService.InsertUser(userReq)
		if err != nil {
			return nil, err
		}
	}

	current_time := time.Now().Local()
	created_at := current_time.Format("2006-01-02")
	modified_at := current_time.Format("2006-01-02")

	credentialEntity := &models.CredentialEntity{
		Id:           uuid.New().String(),
		PlatformName: cred.PlatformName,
		UserName:     cred.UserName,
		Password:     cred.Password,
		CreatedAt:    created_at,
		ModifiedAt:   modified_at,
	}

	const query = `INSERT INTO credentials_tbl 
    (id,platform_name,username,password,created_at,modified_at,user_cl) 
    VALUES (?,?,?,?,?,?,?)`

	tx, err := service.MysqlDB.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()
	_, err = stmt.Exec(credentialEntity.Id,
		credentialEntity.PlatformName,
		credentialEntity.UserName,
		credentialEntity.Password,
		credentialEntity.CreatedAt,
		credentialEntity.ModifiedAt,
		user.Id,
	)
	if err != nil {
		return nil, exceptions.ErrorServer
	}

	tx.Commit()
	return credentialEntity, nil

}

func (service *CredentialService) AcquireCredentials(userReq *models.UserEntity) ([]models.CredentialEntity, error) {
	userService := NewUserService(service.MysqlDB)

	user, err := userService.FetchByID(userReq.Id)

	if err != nil {
		log.Println("acquire user: ", err.Error())
		return nil, err
	}

	const query = `SELECT id,platform_name,username,password,created_at,modified_at 
	FROM credentials_tbl 
	WHERE user_cl = ?`

	rows, err := service.MysqlDB.Query(query, user.Id)
	if err != nil {
		return nil, exceptions.ErrorServer
	}
	defer rows.Close()

	var credentials []models.CredentialEntity

	for rows.Next() {
		var credential models.CredentialEntity

		if err := rows.Scan(&credential.Id,
			&credential.PlatformName,
			&credential.UserName,
			&credential.Password,
			&credential.CreatedAt,
			&credential.ModifiedAt); err != nil {
			return credentials, exceptions.ErrorServer
		}
		credentials = append(credentials, credential)
	}

	return credentials, nil

}
