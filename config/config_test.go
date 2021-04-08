package config

import (
	"fmt"
	"testing"
)

func TestGetDbConfig(t *testing.T) {
	dbConfig := GetDbConfig()
	fmt.Println(dbConfig)
}
