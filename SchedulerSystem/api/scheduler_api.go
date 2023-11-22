package api

import (
	"atse/scheduler_system/broker"
	"atse/scheduler_system/database"
	"atse/scheduler_system/dto"
	"atse/scheduler_system/env"
	"atse/scheduler_system/logger"
	"atse/scheduler_system/scheduler"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	s := v1.Group("/schedule")
	s.GET("", a.getAllSchedules)
	s.GET("/:schedule_id", a.getSchedule)
	s.POST("", a.postSchedule)
	p := v1.Group("/production_line")
	p.GET("", a.getAllProductionLines)
	p.GET("/:production_line_id", a.getProductionLine)
	p.POST("", a.postProductionLine)

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

func (a *Api) getAllSchedules(context *gin.Context) {
	var schedules = make([]dto.Schedule, 0)
	a.database.GetDB().Find(&schedules)
	context.JSON(http.StatusOK, schedules)
}

func (a *Api) getSchedule(context *gin.Context) {
	id := context.Param("schedule_id")
	schedule := dto.Schedule{}
	a.database.GetDB().First(&schedule, id)
	context.JSON(http.StatusOK, schedule)
}

func (a *Api) postSchedule(context *gin.Context) {
	schedule := dto.Schedule{}
	err := context.BindJSON(&schedule)
	if err != nil {
		context.JSON(http.StatusBadRequest, "{\"message\": \"missing schedule\"}")
	}
	a.database.GetDB().Create(&schedule)
	context.JSON(http.StatusCreated, "{\"message\": \"schedule created\"}")
}

func (a *Api) getAllProductionLines(context *gin.Context) {
	var productionLine = make([]dto.ProductionLine, 0)
	a.database.GetDB().Find(&productionLine)
	context.JSON(http.StatusOK, productionLine)
}

func (a *Api) getProductionLine(context *gin.Context) {
	id := context.Param("production_line_id")
	productionLine := dto.ProductionLine{}
	a.database.GetDB().First(&productionLine, id)
	context.JSON(http.StatusOK, productionLine)
}

func (a *Api) postProductionLine(context *gin.Context) {
	productionLine := dto.ProductionLine{}
	err := context.BindJSON(&productionLine)
	if err != nil {
		context.JSON(http.StatusBadRequest, "{\"message\": \"missing production line\"}")
	}
	a.database.GetDB().Create(&productionLine)
	context.JSON(http.StatusCreated, "{\"message\": \"production line created\"}")
}
