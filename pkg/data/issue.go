package data

type Issue interface {
	Save() error
	Update() error

	CanBeRemoved() (bool, error)
	Remove() error

	CanBeCommented() (bool, error)
	Comments() ([]*Comment, error)

	CanBeAssigned() (bool, error)
	Assign() (*User, error)
}
