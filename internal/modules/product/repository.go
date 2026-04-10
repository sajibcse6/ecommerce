package product

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, p *Product) error {
	query := `INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id`

	return r.db.QueryRow(ctx, query, p.Name, p.Description, p.Price).Scan(&p.ID)
}

func (r *Repository) GetAll(ctx context.Context) ([]Product, error) {
	query := `SELECT id, name, description, price FROM products`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product

		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*Product, error) {
	var p Product

	query := `SELECT id, name, description, price FROM products WHERE id=$1`

	err := r.db.QueryRow(ctx, query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price)

	if err != nil {
		return nil, err
	}

	return &p, nil
}
