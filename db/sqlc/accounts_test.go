package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require" // testify is a library containing a set of packages that provide many tools for testifying that your code will behave as you intend.
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    "tom", //randomly generated?
		Balance:  100,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}
