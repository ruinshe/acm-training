package main

import (
	"github.com/gogf/gf/frame/g"
	_ "github.com/ruinshe/acm-training/boot"
	_ "github.com/ruinshe/acm-training/router"
)

func main() {
	g.Server().Run()
}
