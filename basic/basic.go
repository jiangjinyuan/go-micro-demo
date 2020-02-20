package basic

import (
	"github.com/jiangjinyuan/go-micro-demo/basic/config"
	"github.com/jiangjinyuan/go-micro-demo/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
