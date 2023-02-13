package main

import (
	"douyin-lite/comm"
	"douyin-lite/router"
	"github.com/gin-gonic/gin"
)

func main() {

	comm.DB = comm.InitDB()

	r := gin.Default()
	r = router.CollectRoute(r)

	panic(r.Run())

}