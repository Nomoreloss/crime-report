# Crime Report

## Setup

```bash
# install golang-migrate cli
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.1.0

# setup db
make db-create

# run migrations
make migrateup
```
