package models

import (
	"bytes"
	"context"
	"encoding/json"
	"lmwn-go-test/covid-19-summary/logging"
	"lmwn-go-test/covid-19-summary/types"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func GetCovidCases() (*types.CovidCases, error) {
	logger := logging.GetInfoLogger()
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URI"),
	})

	cachedCases, err := rdb.Get(ctx, "covid-cases").Result()
	if err == redis.Nil {
		logger.Println("Cache miss.")
		resp, err := http.Get(os.Getenv("COVID_CASES_API_URL"))
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		bodyStr := buf.String()
		expire, err := strconv.Atoi(os.Getenv("CACHE_EXPIRATION_SEC"))
		if err != nil {
			return nil, err
		}
		err = rdb.Set(ctx, "covid-cases", bodyStr, time.Duration(expire)*time.Second).Err()
		if err != nil {
			return nil, err
		}

		cases := new(types.CovidCases)
		json.Unmarshal([]byte(bodyStr), cases)

		return cases, nil
	} else if err != nil {
		return nil, err
	}

	logger.Println("Cache hit.")

	cases := new(types.CovidCases)
	json.Unmarshal([]byte(cachedCases), cases)
	return cases, nil
}
