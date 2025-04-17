package controllers

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
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

// Get iptv m3u8
// @Title Get iptv m3u8
// @Description get iptv m3u8
// @Success 200 {string} success
// @router /iptv [get]
func (c *MainController) GetIptv() {
	var iptv string
	// 使用系统函数打开iptv.txt文件
	filePath := "./iptv.txt"
	if runtime.GOOS == "linux" {
		filePath = "/zhangyu/iptv.txt"
	}
	if file, err := os.Open(filePath); err != nil {
		fmt.Println("open file failed:", err)
		return
	} else {
		defer file.Close()
		// 使用io.Reader读取文件内容
		if data, err := io.ReadAll(file); err != nil {
			iptv = fmt.Errorf("read file failed: %v", err).Error()
		} else {
			iptv = string(data)
		}
	}
	c.Ctx.WriteString(iptv)
}
