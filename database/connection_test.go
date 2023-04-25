package database_test

import (
	"context"
	"github.com/dozheiny/backend-template/database"
	"testing"
)

func TestConnection(t *testing.T) {
	t.Parallel()

	if err := database.Connection(context.Background()); err != nil {
		t.Errorf("Got Error on connect to database: %s", err.Error())
	}
}
