package main

import (
    "github.com/gin-gonic/gin"
    "easyBackend/controller"
    "easyBackend/middleware"
)



func main() {
    r := gin.Default()
    r.Use(middleware.CORSConfig())

    r.POST("/login", controller.Login)
	r.POST("/register", controller.RegisterHandler)

    protected := r.Group("/api", middleware.JWTMiddleware())
	{
        protected.GET("/clockins/:id", controller.GetClockInsByUserID) //check out employee's clock in/out records
        protected.POST("/clockin", controller.ClockIn) 
        protected.POST("/clockout", controller.ClockOut)
        protected.GET("/users", controller.GetAllUsers) //get all employees info
        protected.GET("/clocklist", controller.GetAllClockList) //get all employees clock in/out records
        protected.GET("/todayClockin", controller.GetTodayClockInCount) //get today's clock in count
        // protected.GET("/userSalary/:id", controller.GetSalaryByUserID) //get employee's salary
	}
    r.Run(":7777") 
}
