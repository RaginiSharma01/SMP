package utils

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

func CheckOTPLimit(ctx context.Context, rdb *redis.Client, email string) error {

	countKey := "otp_count:" + email
	blockKey := "otp_block:" + email

	// check if blocked
	blocked, _ := rdb.Exists(ctx, blockKey).Result()
	if blocked == 1 {
		return errors.New("too many OTP requests. try again after 1 hour")
	}

	// increment count
	count, err := rdb.Incr(ctx, countKey).Result()
	if err != nil {
		return err
	}

	// expire counter after 10 minutes
	rdb.Expire(ctx, countKey, 10*time.Minute)

	// if more than 3 requests → block
	if count > 3 {
		rdb.Set(ctx, blockKey, "blocked", time.Hour)
		return errors.New("too many OTP requests. user blocked for 1 hour")
	}

	return nil
}