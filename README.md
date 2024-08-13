# template-golang

## Running

- Server: `go run ./cmd/server`
- Playground: `go run ./cmd/playground`

## Database Configuration

To generate new schemas for the database run the following command:

```bash
cd internal/ && go run -mod=mod entgo.io/ent/cmd/ent new User
```

and then all the ORM boilerplate can be generated with:

```bash
go generate ./ent
```