package operation

type RegisterStatusCode uint16

const (
	RegisterSuccess RegisterStatusCode = iota
	RegisterServerError
	RegisterDupeUsername
	RegisterDupeEmail
)
