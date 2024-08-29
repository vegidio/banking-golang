# banking-golang

## Running

- Playground: `go run ./cmd/playground`
- API Gateway: `go run ./cmd/apigateway`

## Database Configuration

To generate new schemas for the database run the following command:

```bash
cd internal/<project> && go run -mod=mod entgo.io/ent/cmd/ent new User
```

and then all the ORM boilerplate can be generated with:

```bash
go generate ./ent
```

## Project Architecture

This project uses the MVCS pattern (sans-V) to structure the codebase, where the **Controller** is responsible for handling the requests, doing the initial conversions and validations; the **Service** is responsible for the business logic, and the **Repository** is responsible for the database operations.

### Nomenclature

The functions in each layer should follow the following nomenclature to minimize confusion and naming conflicts:

- Service: `Create...`, `Retrieve...`, `Update...`, `Delete...`
- Repository: `Query...`, `Mutate...`
