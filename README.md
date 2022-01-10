# The LWMN Go Backend Developer Test
## Installation
1. Clone this repository.
```
$ git clone https://github.com/techid2000/lmwn-go-test.git
```
2. Add `.env` file in the root of repository with these environment variables
```
REDIS_URI=redis:6379
COVID_CASES_API_URL=http://static.wongnai.com/devinterview/covid-cases.json
CACHE_EXPIRATION_SEC=10
GIN_MODE=debug
```
3. Run docker compose
```
$ docker-compose up
```
## Testing
Testing is done for the service `GenerateCovidSummary`.
```
$ go test -v ./...
```
