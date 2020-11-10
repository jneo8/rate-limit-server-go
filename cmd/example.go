package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jneo8/rate-limit-server-go/repository/tokenbucket"
	"log"
	"strconv"
)

func main() {
	r := gin.Default()
	tokenbucket := tokenbucket.New(60)

	go func() {
		if err := tokenbucket.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	r.GET("/", func(c *gin.Context) {
		ip := c.ClientIP()
		n := tokenbucket.Get(ip)
		if n >= 60 {
			c.String(200, "error")
		} else {
			c.String(200, strconv.Itoa(n))
		}
	})
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
