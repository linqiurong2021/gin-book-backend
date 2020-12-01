package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-book-backend/models"
	"github.com/linqiurong2021/gin-book-backend/routers"
	"github.com/linqiurong2021/gin-book-frontend/config"
	"github.com/linqiurong2021/gin-book-frontend/mysql"
	"github.com/linqiurong2021/gin-book-frontend/validator"
)

func main() {
	r := gin.Default()

	// 加载配置文件(这里可以使用默认的配置文件)
	if len(os.Args) > 1 {
		if err := config.Init(os.Args[1]); err != nil {
			fmt.Printf("load config from file falure !, err:%v\n", err)
			return
		}
	} else {
		if err := config.Init("./config/config.ini"); err != nil {
			fmt.Printf("load config from file falure !, err:%v\n", err)
			return
		}
		fmt.Printf("\n\n#### load config from config/config.ini ! ####\n\n")
	}

	if err := mysql.InitMySQL(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 开启校验转换
	if err := validator.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	// 绑定数据表
	mysql.DB.AutoMigrate(&models.Cate{})
	// 注册路由
	routers.RegisterRouter(r)
	r.Run(fmt.Sprintf(":%#v", config.Conf.Port))
}
