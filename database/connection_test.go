package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDBAddress(t *testing.T) {
	os.Setenv("BRIZO_MYSQL_USER", "foo")
	os.Setenv("BRIZO_MYSQL_PASSWORD", "bar")
	os.Setenv("BRIZO_MYSQL_HOST", "db.example.com")
	os.Setenv("BRIZO_MYSQL_PORT", "3307")
	os.Setenv("BRIZO_MYSQL_DATABASE", "somedb")

	expected := "foo:bar@(db.example.com:3307)/somedb?charset=utf8&parseTime=True&loc=Local"
	actual := getDBAddress()
	assert.Equal(t, expected, actual)
}

func TestGetDBAddressUsesDefaultPort(t *testing.T) {
	os.Setenv("BRIZO_MYSQL_USER", "foo")
	os.Setenv("BRIZO_MYSQL_PASSWORD", "bar")
	os.Setenv("BRIZO_MYSQL_HOST", "db.example.com")
	os.Setenv("BRIZO_MYSQL_PORT", "")
	os.Setenv("BRIZO_MYSQL_DATABASE", "somedb")

	expected := "foo:bar@(db.example.com:3306)/somedb?charset=utf8&parseTime=True&loc=Local"
	actual := getDBAddress()
	assert.Equal(t, expected, actual)
}
