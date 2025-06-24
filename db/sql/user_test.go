package db

import (
	"context"
	"fmt"
	"github.com/ismail118/booklending/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUser(t *testing.T) {
	for i := 1; i < 10; i++ {
		hashedPassword, err := util.HashPassword("password")
		require.NoError(t, err)
		require.NotEmpty(t, hashedPassword)

		name := RandomName()
		arg := CreateUserParams{
			Name:           name,
			Email:          fmt.Sprintf("%s@gmail.com", name),
			HashedPassword: hashedPassword,
		}

		newId, err := testQuerier.CreateUser(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, newId)
	}
}
