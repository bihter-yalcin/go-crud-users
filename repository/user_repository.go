package repository

import (
	models "api/model"
	"database/sql"
	"errors"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	rows, err := repo.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Mail); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := repo.DB.QueryRow("SELECT id, name, mail FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Mail)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	err := repo.DB.QueryRow("INSERT INTO users (name, mail) VALUES ($1, $2) RETURNING id", user.Name, user.Mail).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) UpdateUser(user *models.User) error {
	_, err := repo.DB.Exec("UPDATE users SET name = $1, mail = $2 WHERE id = $3", user.Name, user.Mail, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) DeleteUser(id int) error {
	_, err := repo.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
