package index

import (
	context "context"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/my/app/controller/db"
)

type Controller struct {
}
type IndexList struct {
	ID         int64  `json:"id"`
	Time       string `json:"time"`
	Content    string `json:"content"`
	Picture    string `json:"picture"`
	Main_title string `json:"main_Title"`
}

// GET /
func (c *Controller) Index(ctx context.Context) (postList []IndexList, err error) {
	var item = IndexList{}
	db, err := db.DB()
	if err != nil {
		return nil, err
	}
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100)
	result, _ := db.Query("select id, time, content, picture, main_title  from data_pool limit " + strconv.Itoa(randomNum) + ", 10")
	for result.Next() {
		result.Scan(&item.ID, &item.Time, &item.Content, &item.Picture, &item.Main_title)
		postList = append(postList, item)
	}
	return postList, nil

}
