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
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := repo.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	err := repo.DB.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) UpdateUser(user *models.User) error {
	result, err := repo.DB.Exec(
		"UPDATE users SET name = $1, email = $2 WHERE id = $3",
		user.Name, user.Email, user.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected, user not found")
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
