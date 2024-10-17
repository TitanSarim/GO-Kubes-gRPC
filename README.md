migration cmd command

```
migrate -path db/migration -database "postgresql://postgres:12345@localhost:5432/simple_bank?sslmode=disable" -verbose up
```

Type ORM is SQLC

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Clean Up the Dependencies

```
go mod tidy
```
