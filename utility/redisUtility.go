package utility

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "", // no password set
	DB:       15, // use default DB
})

func Add2DB(key string) bool {
	now := time.Now()
	sec := now.Unix()
	err := rdb.Set(ctx, key, sec, 0).Err()
	if err != nil {
		Log("ERROR", "add ", key, "to redis failed!!!")
		return false
	}
	Log("INFO", "add ", key, " done successfully")

	return true
}

func RemoveFromDB(key string) bool {
	err := rdb.Del(ctx, "key")
	if err != nil {
		Log("ERROR", "removed ", key, "to redis failed!!!")
		return false
	}
	Log("INFO", "Remove ", key, " done successfully")
	return true
}

func IsExist(key string) int64 {
	value, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		Log("ERROR", "Check existence ", key, "to redis failed!!!")
		return 0
	}
	Log("INFO", "Existence ", key, " done successfully")
	valueRet, _ := strconv.ParseInt(value, 10, 64)
	return valueRet
}

func CheckRelay(key string) bool {
	result := false
	value := IsExist(key)
	if value == 0 {
		Log("INFO", "key : ", key, " is not exist in DB so relay .")
	}

	now := time.Now()
	sec := now.Unix()
	diff := sec - value
	if diff > 15552000 {
		Log("INFO", "key : ", key, " is exist in DB but time will be updated!!!")
		Add2DB(key)
		result = true
	} else {
		Log("WARNING", "key : ", key, " is exist in DB and time is near")
	}

	return result
}
