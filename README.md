# lightban.com

Lightweight Kanban.
Stupid simple.

## Deps 

- Go
- Node / NPM

## Clone
```
git clone git@github.com:rdcx/lightban.git
```
### Create .env 
```
JWT_SECRET="secret"
DB_DSN="root:lban@tcp(128.0.0.1:3306)/lban?charset=utf8mb4&parseTime=True&loc=Local"
```
### Run API locally
```
go run cmd/migration/main.go
go run cmd/api/main.go
```
### Run Frontend Locally
```
cd app; npm run dev;
```
### Build API
```
go build -o apicmd cmd/api/main.go
```
### Build migration to run elsewhere
```
go build -o mig cmd/migration/main.go
```
### Build frontend
```
cd app; npm run build;
```
### Run with TLS
```
PORT=443 ./apicmd
```
