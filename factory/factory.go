package factory

import "github.com/cesardelgadom/go-patterns/factory/connection"

func Factory(typeDB int) connection.DBConnection {
	switch typeDB {
	case 1:
		return &connection.MySql{}
	case 2:
		return &connection.Postgresql{}
	default:
		return nil
	}
}
