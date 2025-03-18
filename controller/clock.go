package controller

import (
    "github.com/gin-gonic/gin"
    "easyBackend/model"
    "net/http"
)

func ClockIn(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
        return
    }

    hasClockedIn, err := model.HasClockedIn(userID.(int))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check clock-in status"})
        return
    }
    if hasClockedIn {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User has already clocked in today"})
        return
    }

    err = model.SaveClockIn(userID.(int))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save clock-in"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Clock-in successful"})
}

func ClockOut(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
        return
    }

    hasClockedOut, err := model.HasClockedOut(userID.(int))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check clock-out status"})
        return
    }
    if hasClockedOut {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User has already clocked out today"})
        return
    }

    err = model.SaveClockOut(userID.(int))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save clock-out"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Clock-out successful"})
}

func GetAllClockList(c *gin.Context) {
    role, exists := c.Get("role")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User role not found"})
        return
    }

    var clockList []model.Attendance
    var err error

    if role == "admin" {
        clockList, err = model.GetAllClockList()
    } else {
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
            return
        }
        clockList, err = model.GetClockInsByUserID(userID.(int))
    }

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve clock-list"})
        return
    }

    c.JSON(http.StatusOK, clockList)
}

func GetTodayClockInCount(c *gin.Context) {
    count, err := model.GetTodayClockInCount()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve today's clock-in count"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"count": count})
}