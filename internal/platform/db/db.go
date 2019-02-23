package db

// Storage defines interface for storage handling
type Storage interface {
	Insert(interface{}) error
	Search(interface{}) ([]interface{}, error)
	Close() error
}

// Config defines data for db configuration
type Config struct {
	Name     string
	Password string
	User     string
}
