## Utility Commands
So that, I do not have to search google for commands every time :p

```shell
docker exec -it postgres14 psql -U root
docker logs postgres14
```
```shell
brew install golang-migrate
migrate create -ext sql -dir db/migration -seq init_schema
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/surelink-db?sslmode=disable" -verbose up  
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/surelink-db?sslmode=disable" -verbose down
```
```shell
brew install sqlc
sqlc init
sqlc generate
```
```shell
go test -v -cover ./...
go run main.go
```