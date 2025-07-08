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
	panic("unimplemented")
}

// GetUserByID implements types.UserStore.
func (s *Store) GetUserByID(id int) (*types.User, error) {
	panic("unimplemented")
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
