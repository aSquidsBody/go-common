package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v9"
)

type RedisClient interface {
	Set(uint, interface{}) error
	Get(uint, interface{}) error
}

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(host, port, password string) (rc RedisClient, err error) {
	connStr := fmt.Sprintf("%s:%s", host, port)
	defer func() {
		if recover() != nil {
			err = fmt.Errorf("Could not connect ot redis server %s", connStr)
		}
	}()

	rc = &redisClient{
		redis.NewClient(&redis.Options{
			Addr:     connStr,
			Password: password,
			DB:       0,
		})}
	return
}

func (rc *redisClient) Set(ID uint, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Error marshalling value for redis ", err)
		return err
	}
	err = rc.client.Set(context.Background(), fmt.Sprint(ID), bytes, 0).Err()
	if err != nil {
		fmt.Printf("Error saving in redis. ID = %d, value = %+v\n", ID, value)
		return err
	}
	return nil
}

func (rc *redisClient) Get(ID uint, v interface{}) error {
	bytes, err := rc.client.Get(context.Background(), fmt.Sprint(ID)).Bytes()
	if err == redis.Nil {
		fmt.Println(fmt.Errorf("ID not found in redis. ID = %d", ID))
		return fmt.Errorf("ID = %d does not exist in redis", ID)
	}
	if err != nil {
		fmt.Printf("Encountered error reading from redis. ID = %d, error = %e\n", ID, err)
		return err
	}

	err = json.Unmarshal(bytes, v)
	if err != nil {
		fmt.Printf("Error unmarshalling object in redis. ID = %d, err = %e\n", ID, err)
		return err
	}
	return nil
}
