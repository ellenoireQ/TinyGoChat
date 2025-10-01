package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Body of JSON
type body struct {
	UID      string `json:"uid"`
	Messages string `json:"messages"`
	Lastime  string `json:"lastime"`
}

// Conversation JSON
type Conversation struct {
	RoomID string `json:"room_id"`
	Users  []body `json:"users"`
}

// Chat definition
// Maybe looks like
//
//		{
//	   "conversation_id": "abc123",
//	   "body": {
//	     "room_id": "room001",
//	     "users": [
//	       {
//	         "uid": "user001",
//	         "messages": "Hello World",
//	         "lastime": "23 Aug 2025"
//	       },
//	       {
//	         "uid": "user002",
//	         "messages": "Hi back!",
//	         "lastime": "23 Aug 2025"
//	       }
//	     ]
//	   }
//	 }
type ChatDefinition struct {
	Conversation_id string       `json:"conversation_id"`
	Body            Conversation `json:"body"`
}

// Main function right here
func main() {
	databaseChat := []ChatDefinition{}
	router := gin.Default()
	router.POST("/sendMess", func(c *gin.Context) {
		var newChat ChatDefinition

		if err := c.BindJSON(&newChat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		}
		databaseChat = append(databaseChat, newChat)
	})
	router.GET("/get", func(c *gin.Context) {
		if len(databaseChat) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "The data is null",
				"array":   databaseChat,
				"status":  http.StatusOK,
			})
		} else {
			c.JSON(http.StatusOK, databaseChat)
		}
	})
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the TinyGoChat API v1.0")
	})
	router.Run(":8080")
}
