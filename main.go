package main

import (
	"log"
	"time"

	"github.com/jeonghoikun/hamjayoung.com/server"
	"github.com/jeonghoikun/hamjayoung.com/site"
	"github.com/jeonghoikun/hamjayoung.com/store"
)

func init() {
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}
	time.Local = loc
	site.Init()
	if err := store.Init(); err != nil {
		panic(err)
	}
}

func main() {
	s := server.New(site.Config.Port)
	log.Fatal(s.Run())
}
