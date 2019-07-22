package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-db/utils"
)

func main() {
	utils.InitDB()
	//utils.GetItem("article", 1, "*", "id")
	//utils.GetItems("article",  "*", "id>0")
	//var rs = utils.ModifyItem("article",1,map[string]string{"title":"okok！！！"},"id")
	//utils.CreateItem("user", map[string]string{"mobile":"1111","password":"4444","nickname":"嘎嘎嘎","status":"1"})
	//utils.DeleteItem("user","3","id")
	//var rs,_ = utils.Query("select * from article")
	//fmt.Println(rs)
	var e = echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/test", func(context echo.Context) error {
		//var results,_ = utils.GetItem("article",1,"*","id")
		var results, _ = utils.Query("select * from article where id = 2")
		fmt.Println(results)
		return utils.Ok(context, results)
		return nil
	})
	e.Logger.Fatal(e.Start(":9787"))
}
