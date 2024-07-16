// INSTALL PACKAGES
$ go mod tidy

// AUTO RELOAD (Air)
$ go install github.com/air-verse/air@latest

// MIGRATION
{Install}
$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

// SECRET KEY
{generate}
$ node -e "console.log(require('crypto').randomBytes(32).toString('hex'))"

{generate migration files}
$ migrate create -ext sql -dir migrations [file name]

{migrate up file}
$ migrate -database "${POSTGRES_DB_URL}" -path EXAMPLE_PATH up