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

// QueryById retrieves a user by their id from the database. It returns a pointer to the ent.User object and an error if
// the query fails.
//
// Parameters:
// - id: The user id to be queried.
//
// Returns:
// - *ent.User: A pointer to the user object if found.
// - error: An error object if the query fails or the user is not found.
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

// QueryUserByEmail retrieves a user by their email address from the database. It returns a pointer to the ent.User
// object and an error if the query fails.
//
// Parameters:
// - email: The email address of the user to be queried.
//
// Returns:
// - *ent.User: A pointer to the user object if found.
// - error: An error object if the query fails or the user is not found.
func QueryUserByEmail(email string) (*ent.User, error) {
	u, err := db.User.
		Query().
		Where(user.Email(email)).
		Only(context.Background())

	if err != nil {
		return nil, err
	}

	return u, nil
}
