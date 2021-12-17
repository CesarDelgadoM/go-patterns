package connection

import (
	"fmt"
	"time"
)

type Postgresql struct {
	db string
}

func (m *Postgresql) Connect() error {
	conn := "Connect to Postgresql..."
	fmt.Println(conn)
	m.db = conn
	return nil
}

func (m *Postgresql) GetNow() (*time.Time, error) {
	t := time.Now()
	return &t, nil
}

func (m *Postgresql) Close() error {
	fmt.Println("Connection closed...")
	return nil
}
