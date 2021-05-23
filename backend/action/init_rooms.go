package action

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/timfuhrmann/spotify-rooms/backend/entity"
	"strconv"
	"time"
)

func InitRooms(rdb *redis.Client) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	size := rdb.HLen(ctx, entity.RoomsKey)
	if size.Val() > 0 {
		return nil
	}

	for i := 0; i < 5; i++ {
		id := strconv.Itoa(i + 1)
		room := entity.Room{
			Id: id,
			Name: "Worldwide #" + id,
			Live: true,
		}

		r, err := room.MarshalRoom()
		if err != nil {
			return err
		}

		if err = rdb.HSet(ctx, entity.RoomsKey, id, r).Err(); err != nil {
			return err
		}
	}

	return nil
}
