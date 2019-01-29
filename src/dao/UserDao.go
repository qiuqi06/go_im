package dao

import (
	"github.com/garyburd/redigo/redis"
	"model"
	"encoding/json"
)
type UserDao struct {
	pool *redis.Pool

}
func(this *UserDao) getUserById(conn redis.Conn,id int )(user *model.User,err error) {
	res, err := redis.String(conn.Do("HGET", "users", id))
	if err != nil {
		if err==redis.ErrNil{
			println("err")
		}
	}
	user= &model.User{}
	json.Unmarshal([]byte(res),user)
	return
}