package main

import (
	"fmt"
	_ "get_ipv6/routers"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func testFunc() {
	testURL := "http://:xxxxxxxxx@localhost:8080/v1/ipv6"
	parsedURL, err := url.Parse(testURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("User:", parsedURL.User.Username())
	if pwd, b := parsedURL.User.Password(); b {
		fmt.Println("Password:", pwd)
	}
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("RawQuery:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)
}

func main() {
	testFunc()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers",
			"Content-Type", "Access-Token", "Accept", "x-requested-with", "Domain"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowOrigins:  []string{"*"},
	}))
	go beego.Run()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	fmt.Printf("Get signal to exit, bye!\n")
}
