migrateup:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/crime_report?sslmode=disable" -verbose up 
migratedown:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/crime_report?sslmode=disable" -verbose down