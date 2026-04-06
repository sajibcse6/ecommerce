package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`

	return r.db.QueryRow(ctx, query, user.Name, user.Email).Scan(&user.ID)
}

func (r *Repository) GetAll(ctx context.Context) ([]User, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, email FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*User, error) {
	var u User

	query := `SELECT id, name, email FROM users WHERE id=$1`

	err := r.db.QueryRow(ctx, query, id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}

	return &u, nil
}