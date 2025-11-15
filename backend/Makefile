login_db:
	psql $(DB_DSN)

login_test_db:
	psql $(TEST_DB_DSN)

migrate_create:
	migrate create -ext sql -dir migrations -seq $(name)

migrate_up:
	migrate -database $(DB_DSN) -path migrations up

migrate_down:
	migrate -database $(DB_DSN) -path migrations down

migrate_force:
	migrate -database $(DB_DSN) -path migrations force $(v)