MYSQL_URL=root:superSecret@tcp(127.0.0.1:3306)/belajarfast

migrate-create:
	@migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up:
	@migrate -database "mysql://$(MYSQL_URL)" -path scripts/migrations up

migrate-down:
	@migrate -database "mysql://$(MYSQL_URL)" -path scripts/migrations down
