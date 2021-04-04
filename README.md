# Flaber-Auth

# Migrations test
sql-migrate up -env="development" -dryrun
sql-migrate down -env="development" -dryrun

# Migrations  postgresql
sql-migrate up -env="development_pgsql"  
sql-migrate down -env="development_pgsql" -limit=0

# Migrations  mssql
sql-migrate up -env="development_mssql"  
sql-migrate down -env="development_mssql" -limit=0

# Generate gql
go run github.com/99designs/gqlgen generate

# Generate Dataloader
go run github.com/vektah/dataloaden CreditsByUserIDsLoader string '[]*gitlab.com/e-sign/api-gql/pkg/nxm/credits.Credit'

# cross compilation
GOOS=linux  GOARCH=amd64 go build
GOOS=windows  GOARCH=amd64 go build
