# template-golang

## Running

- Playground: `go run ./cmd/playground`
- API Gateway: `go run ./cmd/apigateway`

## Database Configuration

To generate new schemas for the database run the following command:

```bash
cd internal/ && go run -mod=mod entgo.io/ent/cmd/ent new User
```

and then all the ORM boilerplate can be generated with:

```bash
go generate ./ent
```