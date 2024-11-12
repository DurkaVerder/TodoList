package postgres_test

import (
	"TodoList/internal/repository/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPostgresRepo(t *testing.T) {
	testRepo := postgres.NewPostgresRepo()

	err := testRepo.CheckConnect()

	assert.NoError(t, err)
}

func TestAddTask(t *testing.T) {
	// connection to fake db

}
