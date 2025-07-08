package product

import (
	"database/sql"

	"github.com/absakran01/ecom/types"
)

type Store struct {
	db *sql.DB
}


func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// CreateProduct inserts a new product into the database.
func (s *Store) CreateProduct(product *types.CreateProductPayLoad) error {
	_, err := s.db.Exec("INSERT INTO products (name, description, image, quantity, price) VALUES (?, ?, ?, ?, ?)",
		product.Name, product.Description, product.Image,
		product.Quantity, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetProductByID(id int) (*types.Product, error) {
	row := s.db.QueryRow("SELECT id, name, description, image, quantity, price, createdAt FROM products WHERE id = ?", id)
	product, err := scanRowIntoProduct(row)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT id, name, description, image, quantity, price, createdAt FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*types.Product
	for rows.Next() {
		product, err := scanRowIntoProducts(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func scanRowIntoProduct(row *sql.Row) (*types.Product, error) {
	product := &types.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Image, &product.Quantity, &product.Price, &product.CreatedAt)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func scanRowIntoProducts(rows *sql.Rows) (*types.Product, error) {
	product := &types.Product{}
	err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Image, &product.Quantity, &product.Price, &product.CreatedAt)
	if err != nil {
		return nil, err
	}
	return product, nil
}
