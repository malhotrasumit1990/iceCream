package dataAccess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayToString(t *testing.T) {

	testArray := []string{"Sumit", "Malhotra"}
	str := arrayToString(testArray)
	assert.Equal(t, "Sumit , Malhotra", str)

}

func TestStringToArray(t *testing.T) {

	str := "Sumit,Malhotra"
	arr := stringToArray(str)
	assert.Equal(t, "Sumit", arr[0])
	assert.Equal(t, "Malhotra", arr[1])

}
