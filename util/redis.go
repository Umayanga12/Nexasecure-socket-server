package util

import (
	"context"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
)

func init() {
	// Parse REDIS_DB (default to 0 on error)
	dbNum := 0
	if dbStr := os.Getenv("REDIS_DB"); dbStr != "" {
		if n, err := strconv.Atoi(dbStr); err == nil {
			dbNum = n
		}
	}

	// Build Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNum,
	})
}

// AddWallet stores conAddr, reqPubKey, and authPubKey in Redis.
func AddWallet(conAddr, reqPubKey, authPubKey string) error {
	// Store the wallet data in a hash
	walletKey := "wallet:" + conAddr
	err := redisClient.HSet(ctx, walletKey, map[string]interface{}{
		"reqpubkey":  reqPubKey,
		"authpubkey": authPubKey,
		"conaddr":    conAddr,
	}).Err()
	if err != nil {
		return err
	}

	// Create reverse mappings for quick lookup
	if err := redisClient.Set(ctx, "reqpubkey:"+reqPubKey, conAddr, 0).Err(); err != nil {
		return err
	}
	if err := redisClient.Set(ctx, "authpubkey:"+authPubKey, conAddr, 0).Err(); err != nil {
		return err
	}

	return nil
}

// DeleteWallet removes the wallet data and associated reverse mappings.
func DeleteWallet(conAddr string) error {
	// Retrieve the wallet data to delete reverse mappings
	walletKey := "wallet:" + conAddr
	reqPubKey, _ := redisClient.HGet(ctx, walletKey, "reqpubkey").Result()
	authPubKey, _ := redisClient.HGet(ctx, walletKey, "authpubkey").Result()

	// Delete the wallet hash
	if err := redisClient.Del(ctx, walletKey).Err(); err != nil {
		return err
	}

	// Delete reverse mappings
	if reqPubKey != "" {
		redisClient.Del(ctx, "reqpubkey:"+reqPubKey)
	}
	if authPubKey != "" {
		redisClient.Del(ctx, "authpubkey:"+authPubKey)
	}

	return nil
}

// GetConAddrByReqPubKey retrieves the conAddr associated with the given reqPubKey.
func GetConAddrByReqPubKey(reqPubKey string) (string, error) {
	return redisClient.Get(ctx, "reqpubkey:"+reqPubKey).Result()
}

// GetConAddrByAuthPubKey retrieves the conAddr associated with the given authPubKey.
func GetConAddrByAuthPubKey(authPubKey string) (string, error) {
	return redisClient.Get(ctx, "authpubkey:"+authPubKey).Result()
}
