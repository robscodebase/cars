# Go Cars API
* This api was built as a working example to demonstrate how to build a simple API in Go.
* It is not a complete example and is not production ready and most likely has some bugs and tech debt.
* This is a beta version and has not been peer reviewed.
* I just completed the code and have not fully reviewed it myself so there may be some obvious bugs or mistakes.

## Running
* You can run the api using docker-compose or `go run main.go`.
* Make sure you have the .env file which is included in the repo. You should never include .env files in your repo, but I have included it here since it is a demo.
* It's best to avoid env vars for sensitive data and use secrets or vault.
* See the Makefile for shell commands.
* `make compose-all` run postgres, redis, and tests.
* `make compose-deps` run postgres and redis only.
* `make compose-down` stop and remove containers.
* `make delete-db-volume` compose-down will not remove the db volume so data will persist. This command will remove the volume.


## Testing
* `make compose-all` run postgres, redis, and tests.
* Postgres and Redis aren't mocked out, all tests are integration tests.
* To run the tests you will need to have docker-compose installed.
* It runs tests in parallel and uses redis and postgres containers.
* Logs will be printed to the console and the test results will be printed to the console.
* You should expect to see output similar to below:
`cars_1      | PASS
  cars_1      | ok  	cars/integration_tests	1.519s
  cars_1      | === RUN   TestAuthenticate
  cars_1      | === RUN   TestAuthenticate/missing_Authorization_header
  cars_1      | === RUN   TestAuthenticate/incorrect_Authorization_format
  cars_1      | === RUN   TestAuthenticate/invalid_token
  cars_1      | === RUN   TestAuthenticate/valid_token
  cars_1      | --- PASS: TestAuthenticate (0.00s)
  cars_1      |     --- PASS: TestAuthenticate/missing_Authorization_header (0.00s)
  cars_1      |     --- PASS: TestAuthenticate/incorrect_Authorization_format (0.00s)
  cars_1      |     --- PASS: TestAuthenticate/invalid_token (0.00s)
  cars_1      |     --- PASS: TestAuthenticate/valid_token (0.00s)
  cars_1      | PASS
  cars_1      | ok  	cars/middleware	0.013s
  cars_cars_1 exited with code 0`

# Trouble Shooting
* If for some reason there is a failure
* Check the following variables in the .env file.
1. POSTGRES_HOST
2. REDIS_HOST
* Since we are using docker-compose it is possible that docker-compose is using a different name for the containers.
* Run `docker ps` to see the running containers and check the names. Update the .env file as needed.
* If you want to run it locally, you can run `make compose-deps` and then run `go run main.go` but you will need to update the .env file to use localhost for the hosts.

### Go
* This API uses the latest version of Go and Go modules.
* It uses Gin as the web framework and Gorm as the ORM.
* It uses Postgres as the database and Redis as the cache.
* The postgres database is used to store the cars and the redis cache is used to store tokens for authentication.
* It uses middleware for authentication.
* There isn't any custom logic except for the authentication middleware and the unit tests are integration tests.

### Bash
1. wait-for is used to wait for the postgres and redis containers to be ready before running the tests using docker-compose.
2. entrypoint.sh is used to run the tests and wait for the postgres and redis containers to be ready before running the tests using docker-compose.

### Docker
* The Dockerfile is used to build the image for the api, it is not used for tests, but it would be used to build the image to run in a production environment.
* Notice in the Dockerfile the bash commands to ensure success and print the reason for any errors. Here we check to ensure the .env file is present before continuing and print a message if it is not. `f="./.env" && if test -f ${f}; then echo "found ${f}"; else echo "couldn't find ${f}"; fi`

### Docker Compose
* Docker compose is only used to run the tests. Typically, you wouldn't run docker-compose in a production environment.
* Docker compose runs Redis, Postgres, and the Cars api.
* When the docker container for the cars api starts it runs the entrypoint.sh script which waits for postgres to be ready then runs the tests.

### Deployment
* This api is not ready for deployment, but if it were, you would build the image using the Dockerfile and then deploy the image to a container orchestration platform such as Kubernetes.
* The docker-compose command can be run in Github Actions, Jenkins, AWS Codebuild, or any other CI/CD platform.
* To deploy to the cloud using Jenkins you could use the Makefile.
* To deploy to the cloud using Codebuild you would need a buildspec.yml with build commands.
* Other deployment considerations include a load balancer, autoscaling, and monitoring.
* An application like this would be set behind a load balancer and would be monitored using Cloudwatch.
* Autoscaling would be used to scale the application based on load.
* The application would be monitored for CPU, RAM, and disk usage.
* Alarms would be created based on thresholds of metrics including HTTP Status codes for example if 500 Status increase by a certain threshold it would trigger an alarm.

## TODO
* Implement a users api or some method to generate tokens, as of now, tokens are created in tests only.
* Implement secrets or vault for sensitive data.
* Update comments and documentation.
