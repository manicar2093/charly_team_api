# Health Records API

## Conventions

### What to consider when create a new handler

This new handler must call `config.Start()` to avoid problems on DB connection

### Testing

All test run using `.env` that you create from `.env.example`. It is need to create the database `health_records_db_test` and after that just run the make command `make test` and all test should run without any problem.

Into the package ```testfunc``` there is a func to start environment for testing purposes.

### Error handling

You can create a new struct which implements HandableErrors interfaz to be able to build an response.

### Handlers Logger

All handler have to log its incomings and errors. This should be done using `logger` package

## Deployment

This is done through `serverless` and with the command `make` which compile and active `serverless cli` command.

## Considerations

On Cognito PreTokekenGen must be configurated manually. [More info](https://stackoverflow.com/questions/54530537/serverless-framework-cognito-userpool-pre-token-generator)
