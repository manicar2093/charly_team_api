init_dev_env:
	@ echo "Installing all NPM packages need"
	@ npm install
	@ echo "Initializing dev db..."
	@ docker-compose -f dev_kit.yml up -d
	@ echo "Creating DB..."
	@ knex migrate:latest
	@ echo "Running seeds..."
	@ knex seed:run
	@ echo "DONE! :D"

db_rollback:
	@ knex migrate:rollback --all

db_fill:
	@ knex migrate:latest
	@ knex seed:run

db_testing_fill:
	@ knex migrate:latest --env testing
	@ knex seed:run --env testing

db_testing_rollback:
	@ rm -r testing.db

mocking:
	@ mockery --all

test:
	@ - make db_testing_fill
	@ - go test ./... -v
	@ - make db_testing_rollback

coverage:
	@ make db_testing_fill
	@ go test -cover ./...
	@ make db_testing_rollback

coverage_html:
	@ make db_testing_fill
	@ go test ./... -coverprofile=coverage.out
	@ go tool cover -html=coverage.out
	@ make db_testing_rollback

clean:
	@ rm -r ./dist