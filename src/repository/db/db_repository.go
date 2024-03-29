package db

import (
	"github.com/williamliang705/bookstore-oauth-api/src/clients/cassandra"
	"github.com/williamliang705/bookstore-oauth-api/src/domain/access_token"
	"github.com/williamliang705/bookstore-oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return nil, errors.NewInternalServeError("database connection not implemented yet")
}
