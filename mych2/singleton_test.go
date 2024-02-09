package mych2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetInstance(t *testing.T) {
	counterOne := GetInstance()
	require.NotNil(t, counterOne)
	require.Equal(t, 0, counterOne.GetCounter())

	counterOne.AddOne()
	require.Equal(t, 1, counterOne.GetCounter())

	counterTwo := GetInstance()
	require.NotNil(t, counterTwo)
	require.Equal(t, 1, counterTwo.GetCounter())

	counterTwo.AddOne()
	require.Equal(t, 2, counterTwo.GetCounter())
}
