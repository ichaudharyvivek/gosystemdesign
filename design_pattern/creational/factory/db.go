package factory

type DB interface {
	Connect()
	ExecuteQuery(query string)
}
