package factory

import (
	"fmt"
	"strings"
)

type Factory struct{}

func NewDBFactory(dbType string, conn string) DB {
	switch strings.ToLower(dbType) {
	case "mysql":
		return NewMySQL(conn)

	case "mongodb":
		return NewMongoDB(conn)

	default:
		fmt.Println("Unknown type of DB:", dbType)
		return nil
	}
}
