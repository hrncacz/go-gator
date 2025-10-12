up:
	@goose --dir ./sql/schema/ postgres postgres://postgres:postgres@localhost:5432/gator up

down:
	@goose --dir ./sql/schema/ postgres postgres://postgres:postgres@localhost:5432/gator down
