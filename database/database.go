package database

import (
	"fmt"
	"gorm-mysql/models"
	"log"
	"os"
	"strings"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DBConn *gorm.DB
)

type dummyCacher struct {
	store *sync.Map
}

func (c *dummyCacher) init() {
	if c.store == nil {
		c.store = &sync.Map{}
	}
}

func (c *dummyCacher) Get(key string) interface{} {
	fmt.Printf("%v\n", key)
	if strings.HasPrefix(key, "SELECT * FROM `users`") {
		c.store.Delete(key)

	}
	c.init()
	val, ok := c.store.Load(key)
	if !ok {
		return nil
	}

	return val.(interface{})
}

func (c *dummyCacher) Store(key string, val interface{}) error {
	c.init()
	c.store.Store(key, val)
	return nil
}

// connectDb
func ConnectDb() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:julone520@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn2 := "root:julone520@tcp(127.0.0.1:3306)/test2?charset=utf8mb4&parseTime=True&loc=Local"
	/*
		NOTE:
		To handle time.Time correctly, you need to include parseTime as a parameter. (more parameters)
		To fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4. See this article for a detailed explanation
	*/
	//f,_ := os.OpenFile("./test.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0777)
	//mainL.SetOutput(f)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "mps-",
		},
	})

	db.Callback().Query().After("gorm:query").Register("my_plugin:after_query", AfterQuery)
	//cachesPlugin := &caches.Caches{Conf: &caches.Config{
	//	Easer:  true,
	//	Cacher: &dummyCacher{},
	//}}
	//_ = db.Use(cachesPlugin)
	//db = db.Debug()
	//db.Use(dbresolver.Register(dbresolver.Config{
	//	// use `db2` as sources, `db3`, `db4` as replicas
	//	Sources:  []gorm.Dialector{mysql.Open(dsn)},
	//	Replicas: []gorm.Dialector{mysql.Open(dsn),mysql.Open(dsn)},
	//	// sources/replicas load balancing policy
	//	Policy: dbresolver.RandomPolicy{},
	//	TraceResolverMode: true,
	//	// print sources/replicas mode in logger
	//}))

	//db.Use(prometheus.New(prometheus.Config{
	//	DBName:          "test", // 使用 `DBName` 作为指标 label
	//	RefreshInterval: 15,    // 指标刷新频率（默认为 15 秒）
	//	PushAddr:        "prometheus pusher address", // 如果配置了 `PushAddr`，则推送指标
	//	StartServer:     true,  // 启用一个 http 服务来暴露指标
	//	HTTPServerPort:  8080,  // 配置 http 服务监听端口，默认端口为 8080 （如果您配置了多个，只有第一个 `HTTPServerPort` 会被使用）
	//	MetricsCollector: []prometheus.MetricsCollector {
	//		&prometheus.MySQL{
	//			VariableNames: []string{"Threads_running"},
	//		},
	//	},  // 用户自定义指标
	//}))
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.AutoMigrate(&models.Book{})
	DBConn = db

}
