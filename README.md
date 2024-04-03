## Download Migrate tool

 The migrate tool is used for database migrations.

https://dev.to/techschoolguru/how-to-write-run-database-migration-in-golang-5h6g
https://github.com/golang-migrate/migrate

## Extra Work!! 
For you to create a migration, you need to create the kodboris database

Connect to the postgres database server through the docker container exec
and run the init.sql script in db/init.sql. You can confirm this by connecting to the 
kodboris database with the postgres user.

migrate create -ext sql -dir db/migration -seq  init-schema

## Copy db schema into the migrate migration file

`Copy db schema from db/shcema.sql into *****_init-shecma.up.sql

## Initialize a migration file
`make migrate_int`
Inject your sql data

## Auto-Migration
Auto-migration works on app start-up. Migration of schemas to the database is handled automatically

`migrate -path db/migration -database "postgresql://username:password@localhost:5432/database?sslmode=disable" -verbose up`

## Post Data to database
You can use postman for this.
Navigate to the resource url `http://localhost:3000/member`
`{"first_name": "example", "last_name": "example"}`
