package users

import (
	"context"
	"template-golang/internal/ent"
	"template-golang/internal/ent/user"
)

var db *ent.Client

func SetupRepository(client *ent.Client) {
	db = client
}

func QueryById(id int) (*ent.User, error) {
	u, err := db.User.
		Query().
		Where(user.ID(id)).
		Only(context.Background())

	if err != nil {
		return nil, err
	}

	return u, nil
}
