package unit

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRepoUser_Create(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// userRepo :=
}
