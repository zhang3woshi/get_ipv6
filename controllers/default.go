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
	var ipv6 string
	cmd := "ip -6 addr show dev pppoe-wan | grep 'inet6' | grep 'global' | awk '{print $2}' | cut -d/ -f1"
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// ipv6 := strings.TrimSpace(string(out))
	// c.Ctx.WriteString(ipv6)
	// 将输出按行分割成数组
	addresses := strings.Split(strings.TrimSpace(string(out)), "\n")

	// 过滤出长度为 128 位的 IPv6 地址
	for _, addr := range addresses {
		if strings.Count(addr, ":") == 7 { // IPv6 地址应包含 7 个冒号
			ipv6 = addr
			break
		}
	}

	// 将过滤后的地址返回给客户端
	c.Ctx.WriteString(ipv6)
}
