package store

import "database/sql"

type Store struct {
	config *Config
}

func New(config *Config) *Store{
	return &Store{
		config: config,
	}
}


func (s *Store) Open() error{
	db, err := sql.Open("postgress", s.config.DatabaseURL)
	if err != nil{
		return err
	}
	return nil
}

func (s *Store) Close(){

}