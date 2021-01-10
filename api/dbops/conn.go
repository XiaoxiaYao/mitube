package dbops

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"time"
)

type User struct {
	Id     int64
	UserName   string
	Pwd string
}

type Video struct {
	Id int64
	AuthorId int64
	Name string
	CreatedAt time.Time
}

type Session struct {
	Id int64
	TTL int64
	UserId int64
}

var (
	db *pg.DB
	err error
)

func init() {
	db = pg.Connect(&pg.Options{
		User: "postgres",
		Password: "postgres",
		Database: "postgres",
	})
	err = createSchema(db)
	if err != nil {
		panic(err)
	}
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
		(*Video)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
