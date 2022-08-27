package library

type (
	Errors     string
	ErrorsList map[Errors]string
)

const (
	FAIL_LOAD_ENV Errors = "The file (.env) hasn't been loaded."
)

var Erro = ErrorsList{
	FAIL_LOAD_ENV: "The file (.env) hasn't been loaded.",
}
