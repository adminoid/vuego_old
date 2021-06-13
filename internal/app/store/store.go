package store

import (
	"fmt"
	"github.com/adminoid/vuego/internal/config"
)

type Store struct {
	DatabaseUrl string
}

func New(config config.Config) *Store {
	DatabaseUrl := fmt.Sprintf("user=%s dbname=%s sslmode=disable", config.DbUser, config.DbName)
	fmt.Println(DatabaseUrl + "\n")
	return &Store{
		DatabaseUrl: DatabaseUrl,
	}
}

func (s *Store) Open() error {
	return nil
}

func (s *Store) Close() {

}
