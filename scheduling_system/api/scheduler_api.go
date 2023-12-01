package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/danielbahrami/se07-atsa/scheduling_system/broker"
	"github.com/danielbahrami/se07-atsa/scheduling_system/database"
	"github.com/danielbahrami/se07-atsa/scheduling_system/dto"
	"github.com/danielbahrami/se07-atsa/scheduling_system/env"
	"github.com/danielbahrami/se07-atsa/scheduling_system/logger"
	"github.com/danielbahrami/se07-atsa/scheduling_system/scheduler"
	"github.com/gin-gonic/gin"
)

type Api struct {
	scheduler scheduler.IScheduler
	database  database.IDatabase
	broker    broker.IBroker
}

type Builder struct {
	scheduler scheduler.IScheduler
	database  database.IDatabase
	broker    broker.IBroker
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Broker(broker broker.IBroker) *Builder {
	b.broker = broker
	return b
}

func (b *Builder) Scheduler(scheduler scheduler.IScheduler) *Builder {
	b.scheduler = scheduler
	return b
}

func (b *Builder) Database(database database.IDatabase) *Builder {
	b.database = database
	return b
}

func (b *Builder) Build() Api {
	return Api{
		scheduler: b.scheduler,
		database:  b.database,
		broker:    b.broker,
	}
}

func (a *Api) Start() {
	log.Println("Initializing ...")
	a.broker.Connect()
	logger.SetBroker(a.broker)

	logger.GetLogger().Log("starting")
	a.database.Connect()

	router := gin.Default()
	router.Use(CORSMiddleware())

	v1 := router.Group("/api/v1")
	r := v1.Group("/robot")
	r.GET("", a.getAllRobots)
	r.GET("/:id", a.getRobot)
	r.GET("/:id/:signal", a.signalRobot)

	a.broker.Subscribe("topic/robot/new", func(m string) {
		robot := &dto.Robot{
			ID:    m,
			State: "IDLE",
		}
		a.database.GetDB().Create(robot)
		a.broker.Subscribe(fmt.Sprintf("topic/%s", m), func(m string) {
			robot.State = m
			a.database.GetDB().Save(robot)
		})
	})
	a.broker.Subscribe("topic/production/gpu/completed", func(m string) {
		log.Println(m)
		//FIXME: does not update the amount
		_, id, _ := strings.Cut(strings.Split(m, ",")[0], "=")
		robot := &dto.Robot{}
		a.database.GetDB().First(&robot, "ID = ?", id)
		robot.ProductsProduces++
		a.database.GetDB().Save(&robot)
	})

	port := env.Get("API_PORT")
	router.Run(fmt.Sprintf("0.0.0.0:%s", port))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (a *Api) getAllRobots(context *gin.Context) {
	robots := make([]dto.Robot, 0)
	a.database.GetDB().Find(&robots)
	context.JSON(http.StatusOK, robots)
}

func (a *Api) getRobot(context *gin.Context) {
	id := context.Param("id")
	robot := dto.Robot{}
	a.database.GetDB().First(&robot, "ID = ?", id)
	context.JSON(http.StatusOK, robot)
}

func (a *Api) signalRobot(context *gin.Context) {
	robotId := context.Param("id")
	signal := context.Param("signal")
	a.broker.Message(fmt.Sprintf("topic/%s/signal", robotId), signal)
	context.Status(http.StatusOK)
}
