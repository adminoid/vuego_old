package store

import (
	"fmt"
	"github.com/adminoid/vuego/internal/config"
	"strings"
	"testing"
)

func TestStore(t *testing.T) (*Store, func(...string)) {
	t.Helper()

	cnf := config.NewConfig(".env.test")
	s := New(cnf)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(
					fmt.Sprintf(
						"TRUNCATE %s CASCADE",
						strings.Join(tables, ", ")));
			err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}

}
