DB_URL=postgres://postgres:1918@localhost:5432/oy_3

cl:
	clear

run:
	go run main.go

migrate_file:
	migrate create -ext sql -dir migrations -seq oy_3

migrate_up:
	migrate -path migrations -databse ${DB_URL} -verbose up

migrate_down:
	migrate -path migrations -databse ${DB_URL} -verbose down 

migrate_force:
	migrate -path migrations -databse ${DB_URL} -verbose force 1