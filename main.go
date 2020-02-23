package main

import (
	_ "github.com/UESTC-ACM/acm-training/boot"
	_ "github.com/UESTC-ACM/acm-training/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
