deploy:
	fly deploy

db:
	fly proxy 5432 -a bourgeoisie-oscars-2024-postgres

templ:
	templ generate --watch

templ-clean:
	rm **/*_templ.go

dev:
	gow run .

NAME = migration

create_migration:
	migrate create -ext sql -dir ./database/migrations ${NAME}

migrate:
	migrate -path ./database/migrations -database ${DATABASE_URL} up

migrate-reset:
	migrate -path ./database/migrations -database ${DATABASE_URL} drop

seed:
	go run ./database/seeds/seed.go

tailwind:
	tailwindcss -c tailwind/tailwind.config.js -i ./tailwind/src/styles.css -o ./static/css/styles.css --watch

.PHONY: deploy db templ dev create_migration migrate seed tailwind templ-clean
