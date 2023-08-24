# Go Cars API
This api was built as a working example to demonstrate how to build a simple API in Go.
Keep in mind it is not a complete example and is not production ready and most likely has some bugs and tech debt.
This is a beta version and has not been peer reviewed.
I have just completed the code and have not fully reviewed it myself so there may be some obvious bugs or mistakes.

## Running
You can run the api using docker-compose or go run main.go
Make sure you have the .env file which is included in the repo. You should never include .env files in your repo, but I have included it here since it is a demo.
It's best to avoid env vars for sensitive data and use secrets or vault.
See the Makefile for shell commands.
`make compose all` will run the api, postgres, and redis.


## Testing
To run the tests you will need to have docker and docker-compose installed.


### Go
This API was built using the latest version of Go and uses Go modules.
It uses Gin as the web framework and Gorm as the ORM.
It uses Postgres as the database and Redis as the cache.
The postgres database is used to store the cars and the redis cache is used to store tokens for authentication.
It uses middleware for authentication.
There isn't any custom logic in services and the unit tests are integration tests.



## TODO
* Finish users api. The user api hasn't been implemented so there is no way to create a token. The token creation is handled in tests, but there is no endpoint that you can retrieve a token.
* Implement secrets or vault for sensitive data.
