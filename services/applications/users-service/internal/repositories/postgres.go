package repositories

import (
	"database/sql"
	"fmt"

	"github.com/TineoC/users-service/internal/handlers"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetAll() ([]handlers.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]handlers.User, 0)
	for rows.Next() {
		var u handlers.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *PostgresUserRepository) Create(user handlers.User) (handlers.User, error) {
	if exists, err := r.exists("id", user.ID); err != nil {
		return handlers.User{}, err
	} else if exists {
		return handlers.User{}, fmt.Errorf("user with id %s already exists", user.ID)
	}

	if exists, err := r.exists("email", user.Email); err != nil {
		return handlers.User{}, err
	} else if exists {
		return handlers.User{}, fmt.Errorf("user with email %s already exists", user.Email)
	}

	_, err := r.db.Exec(
		"INSERT INTO users (id, name, email) VALUES ($1, $2, $3)",
		user.ID, user.Name, user.Email,
	)
	if err != nil {
		return handlers.User{}, err
	}

	return user, nil
}

// small helper to check existence
func (r *PostgresUserRepository) exists(field string, value any) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM users WHERE %s = $1)", field)
	err := r.db.QueryRow(query, value).Scan(&exists)
	return exists, err
}
