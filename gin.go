package main

import (
	"dds_core_server/config"
	"dds_core_server/controller"
	"dds_core_server/kafka"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 初始化配置
	conf := config.ConfigInit()

	// 初始化kafka
	err := kafka.InitKafkaProducer(conf.Kafka.Addr)
	if err != nil {
		log.Fatalf("kafka producer init err: %v", err)
		return
	}
	defer kafka.Producer.Close()

	// 初始化gin框架
	r := gin.Default()
	// dds启动时的载入接口
	r.POST("/send_info", controller.SendInfo)

	r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}
