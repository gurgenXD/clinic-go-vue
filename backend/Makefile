migrations-make:
	@migrate create -ext sql -dir ./migrations -seq test

migrations-up:
	@go run main.go migration up

migrations-down:
	@go run main.go migration down
