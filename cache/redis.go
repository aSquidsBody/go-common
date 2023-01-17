package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aSquidsBody/go-common/logs"
	"github.com/go-redis/redis/v9"
)

type RedisClient interface {
	Set(string, interface{}) error
	Get(string, interface{}) error
	Delete(string) error
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

func (rc *redisClient) Set(ID string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Error marshalling value for redis ", err)
		return err
	}
	err = rc.client.Set(context.Background(), ID, bytes, 0).Err()
	if err != nil {
		fmt.Printf("Error saving in redis. ID = %s, value = %+v\n", ID, value)
		return err
	}
	return nil
}

func (rc *redisClient) Get(ID string, v interface{}) error {
	bytes, err := rc.client.Get(context.Background(), ID).Bytes()
	if err == redis.Nil {
		fmt.Println(fmt.Errorf("ID not found in redis. ID = %s", ID))
		return redis.Nil
	}
	if err != nil {
		fmt.Printf("Encountered error reading from redis. ID = %s, error = %e\n", ID, err)
		return err
	}

	err = json.Unmarshal(bytes, v)
	if err != nil {
		fmt.Printf("Error unmarshalling object in redis. ID = %s, err = %e\n", ID, err)
		return err
	}
	return nil
}

func (rc *redisClient) Delete(ID string) error {
	_, err := rc.client.Del(context.Background(), ID).Result()
	if err != nil {
		logs.Errorf(err, "Could not delete %s in redis", ID)
		return err
	}

	return nil
}
