build_deploy: compile_aws
	serverless deploy --aws-profile health_records_dev

compile_aws:
	./compiler.sh cmd/aws/lambda bin/aws/lambda/

db_rollback:
	@ knex migrate:rollback --all

db_fill:
	@ knex migrate:latest
	@ knex seed:run

mocking:
	@ mockery --all

test:
	@ - make db_fill
	@ - go test ./internal/... -v
	@ - make db_rollback

coverage:
	@ make db_fill
	@ go test -cover ./...
	@ make db_rollback

coverage_html:
	@ make db_fill
	@ go test ./... -coverprofile=coverage.out
	@ go tool cover -html=coverage.out
	@ make db_rollback

clean:
	@ - rm -r ./dist
	@ - rm -r coverage.out
	@ - rm -r ./bin
