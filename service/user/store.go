package user

import (
	"database/sql"
	"fmt"

	"github.com/absakran01/ecom/types"
)

type Store struct {
	db *sql.DB
}

// CreateUser implements types.UserStore.
func (s *Store) CreateUser(user *types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, user.Password)
	return err
}

// GetUserByID implements types.UserStore.
func (s *Store) GetUserByID(id int) (*types.User, error) {
	row, err := s.db.Query("SELECT user FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	 user, err := scanRowIntoUser(row)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("user with id %d not found", id)
	}
	return user, nil
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user with email %s not found", email)
	}
	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	u := new(types.User)
	err := rows.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}
