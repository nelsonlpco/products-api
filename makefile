run-migrate:
	migrate create -ext sql -dir src/database/migrations -seq create_product_table
