# Repositories

We must consider create the needed errors for the internal use of this package and all repositories must follow the next conventions:

- If a register is not found repository must return an NotFoundError from `errors.go`
- All method meant to save or update data must be into a transaction and receive a instance of `context.Context`

These conventions can increase as many errors are found to be handled
