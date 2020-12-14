package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	entitiesKey = "entities"
	entitiesTTL = 5 * time.Minute
)

type redisClient struct {
	*redis.Client
	TTL time.Duration
}

type RedisClient interface {
	SetEntities(ctx context.Context, entities []*Entity) error
	GetEntities(ctx context.Context) ([]*Entity, error)
}

func (cli *redisClient) GetEntities(ctx context.Context) ([]*Entity, error) {
	var res []*Entity
	data, err := cli.Get(ctx, entitiesKey).Bytes()
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, &res)
}

func (cli *redisClient) SetEntities(ctx context.Context, entities []*Entity) error {
	data, err := json.Marshal(&entities)
	if err != nil {
		return err
	}
	return cli.Set(ctx, entitiesKey, data, entitiesTTL).Err()
}

func NewRedisClient(host, port string, ttl time.Duration) (RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, fmt.Errorf("try to ping to redis: %w", err)
	}

	c := &redisClient{
		TTL:    ttl,
		Client: client,
	}

	return c, nil
}

func (c *redisClient) Close() error {
	return c.Client.Close()
}
