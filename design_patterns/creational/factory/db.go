package factory

// The base interface which is implemented by other object
type DB interface {
	Connect()
	ExecuteQuery(query string)
}
