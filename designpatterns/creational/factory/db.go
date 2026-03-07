package factory

import "fmt"

// The base interface which is implemented by other object
type DB interface {
	Connect()
	ExecuteQuery(query string)
}

// -- Concrete impl
type MySQL struct {
	conn string
}

func NewMySQL(conn string) *MySQL {
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

// -- Concrete impl
type MongoDB struct {
	conn string
}

func NewMongoDB(conn string) *MongoDB {
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
