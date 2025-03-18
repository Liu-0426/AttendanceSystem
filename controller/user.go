package controller

import (
    "github.com/gin-gonic/gin"
    "easyBackend/model"
    "net/http"
    "strconv"
    
)

func GetAllUsers(c *gin.Context) {
	users, err := model.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetClockInsByUserID(c *gin.Context) {
    userIDStr := c.Param("id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    clockIns, err := model.GetClockInsByUserID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve clock-ins"})
        return
    }

    c.JSON(http.StatusOK, clockIns)
}