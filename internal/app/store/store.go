package store

import (
	"database/sql"
	"fmt"
	"github.com/adminoid/vuego/internal/config"
	_ "github.com/lib/pq"
)

type Store struct {
	DatabaseUrl string
	db *sql.DB
	userRepository *UserRepository
}

func New(config config.Config) *Store {
	DatabaseUrl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost,
		config.DbUser,
		config.DbPwd,
		config.DbName)
	return &Store{
		DatabaseUrl: DatabaseUrl,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.DatabaseUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() error {
	err := s.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
