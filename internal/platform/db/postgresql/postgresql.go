package postgresql

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/saromanov/go-blog/internal/blog"
	st "github.com/saromanov/go-blog/internal/platform/db"
	"github.com/saromanov/go-blog/internal/user"
)

// storage implements db handling with Postgesql
type storage struct {
	db *gorm.DB
}

// New provides init for postgesql storage
func New(s *st.Config) (st.Storage, error) {
	if s == nil {
		return nil, errors.New("config is not defined")
	}
	args := "dbname=goblog"
	if s.Name != "" && s.Password != "" && s.User != "" {
		args += fmt.Sprintf(" user=%s dbname=%s password=%s", s.User, s.Name, s.Password)
	}
	db, err := gorm.Open("postgres", args)
	if err != nil {
		return nil, fmt.Errorf("unable to open db: %v", err)
	}
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&blog.Blog{})
	return &storage{
		db: db,
	}, nil
}

// Insert provides inserting of data
func (s *storage) Insert(m interface{}) error {
	err := s.db.Create(m).Error
	if err != nil {
		return fmt.Errorf("storage: unable to insert data: %v", err)
	}
	return nil
}

// Insert provides finding data
func (s *storage) Search(sr interface{}) ([]interface{}, error) {
	return nil, nil
}

// Close provides closing of db
func (s *storage) Close() error {
	return s.db.Close()
}
