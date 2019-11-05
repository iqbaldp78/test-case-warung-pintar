package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
)

//PingServer is controllers for request to server
func PingServer(c *gin.Context) {
	msg := c.Query("message")
	if msg == "" {
		msg = "pong"
	}
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

//PingClient is controllers for request from client to PingServer
func PingClient(c *gin.Context) {
	var response struct {
		Msg string `json:"message"`
	}

	msg := c.Query("message")

	// Create a Resty Client
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"message": msg,
		}).
		SetHeader("Accept", "application/json").
		Get("http://localhost/ping/server")

	//raise error when cant connect to server
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	//raise error when cant unmarshall body response
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println("response", response)
	c.JSON(http.StatusOK, gin.H{"message": response.Msg})
}

//PingWs is controllers for request to server via websocket
func PingWs(c *gin.Context) {
	var wsupgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer conn.Close()
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		log.Printf("recv: %s", msg)
		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
