package factory

import "fmt"

type MySQL struct {
	conn string
}

func NewMySQL(conn string) DB {
	return &MySQL{conn: conn}
}

// Connect implements DB.
func (m *MySQL) Connect() {
	fmt.Println("Connecting to MySQL server at", m.conn)
}

// ExecuteQuery implements DB.
func (m *MySQL) ExecuteQuery(query string) {
	fmt.Println("Executing", query)
}

// ------

type MongoDB struct {
	conn string
}

func NewMongoDB(conn string) DB {
	return &MongoDB{conn: conn}
}

// Connect implements DB.
func (m *MongoDB) Connect() {
	fmt.Println("Connecting to MongoDB server at", m.conn)
}

// ExecuteQuery implements DB.
func (m *MongoDB) ExecuteQuery(query string) {
	fmt.Println("Executing", query)
}
