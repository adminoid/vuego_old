package store

import "github.com/adminoid/vuego/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
		u.Name, u.Email).Scan(&u.Id); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(u *model.User) (*model.User, error) {
	return nil, nil
}
