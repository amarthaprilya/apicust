package repository

import (
	"customer/models"
	"database/sql"
)

type CustomerRepository struct {
	DB *sql.DB
}

func (r *CustomerRepository) GetAll() ([]models.Customer, error) {
	rows, err := r.DB.Query("SELECT id, name, email, phone, created_at, updated_at FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var c models.Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (r *CustomerRepository) GetByID(id int) (*models.Customer, error) {
	var c models.Customer
	err := r.DB.QueryRow("SELECT id, name, email, phone, created_at, updated_at FROM customers WHERE id = ?", id).
		Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CustomerRepository) Create(c *models.Customer) error {
	result, err := r.DB.Exec("INSERT INTO customers (name, email, phone, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())",
		c.Name, c.Email, c.Phone)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	return r.DB.QueryRow("SELECT id, name, email, phone, created_at, updated_at FROM customers WHERE id = ?", id).
		Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.CreatedAt, &c.UpdatedAt)
}

func (r *CustomerRepository) Update(id int, c *models.Customer) error {
	_, err := r.DB.Exec("UPDATE customers SET name=?, email=?, phone=?, updated_at=NOW() WHERE id=?",
		c.Name, c.Email, c.Phone, id)
	if err != nil {
		return err
	}
	return r.DB.QueryRow("SELECT id, name, email, phone, created_at, updated_at FROM customers WHERE id = ?", id).
		Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.CreatedAt, &c.UpdatedAt)
}

func (r *CustomerRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM customers WHERE id=?", id)
	return err
}
