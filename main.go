package main

import (
	"fmt"
	"net/http"

	"github.com/evanxzj/go-gin-example/routers"

	"github.com/evanxzj/go-gin-example/pkg/setting"
)

func main() {
	r := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
