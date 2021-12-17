package connection

import (
	"fmt"
	"time"
)

type MySql struct {
	db string
}

func (m *MySql) Connect() error {
	conn := "Connect to MySql..."
	fmt.Println(conn)
	m.db = conn
	return nil
}

func (m *MySql) GetNow() (*time.Time, error) {
	t := time.Now()
	return &t, nil
}

func (m *MySql) Close() error {
	fmt.Println("Connection closed...")
	return nil
}
