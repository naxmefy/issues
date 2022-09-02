package data

type Source interface {
	Fetch(args []string) ([]*Issue, error)
}
