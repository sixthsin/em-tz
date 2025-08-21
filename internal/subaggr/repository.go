package subaggr

import "sub-aggregator/pkg/db"

type Repository struct {
	Database *db.Db
}

func NewRepository(database *db.Db) *Repository {
	return &Repository{
		Database: database,
	}
}
