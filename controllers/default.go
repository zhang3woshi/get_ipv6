package controllers

import (
	"fmt"
	"os/exec"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

// Get router ipv6 addr
// @Title Get
// @Description get ipv6 addr
// @Success 200 {string} success
// @router / [get]
func (c *MainController) Get() {
	cmd := "ip -6 addr show dev pppoe-wan | grep 'inet6' | grep 'global' | awk '{print $2}' | cut -d/ -f1"
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	ipv6 := strings.TrimSpace(string(out))
	c.Ctx.WriteString(ipv6)
}
