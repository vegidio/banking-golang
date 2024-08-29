package users

import (
	"banking/internal/apigateway/ent"
	"banking/internal/apigateway/ent/user"
	"context"
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
