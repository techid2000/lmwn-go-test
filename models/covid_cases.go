package models

import (
	"bytes"
	"context"
	"encoding/json"
	"lmwn-go-test/covid-19-summary/types"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func GetCovidCases() (*types.CovidCases, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	cachedCases, err := rdb.Get(ctx, "covid-cases").Result()
	if err == redis.Nil {
		resp, err := http.Get("http://static.wongnai.com/devinterview/covid-cases.json")
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		bodyStr := buf.String()
		err = rdb.Set(ctx, "covid-cases", bodyStr, 10*time.Second).Err()
		if err != nil {
			return nil, err
		}

		cases := new(types.CovidCases)
		json.Unmarshal([]byte(bodyStr), cases)

		return cases, nil
	} else if err != nil {
		return nil, err
	}

	cases := new(types.CovidCases)
	json.Unmarshal([]byte(cachedCases), cases)

	return cases, nil
}
