migrateup:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/crime_report?sslmode=disable" -verbose up 
migratedown:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/crime_report?sslmode=disable" -verbose down
db-create:
	psql -U postgres -d crime_report -f db/crime_reporter.sql
db-drop:
	psql -U postgres -f db/drop_all.sql
