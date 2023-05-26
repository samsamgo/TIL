package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Create a new channel to send messages between clients.
	messages := make(chan string)

	// Handle the `/` route.
	r.GET("/", func(c *gin.Context) {
		// Get the user's name from the query string.
		name := c.Query("name")

		// If the user's name is empty, redirect them to the `/login` route.
		if name == "" {
			c.Redirect(302, "/login")
			return
		}

		// Send the user's name to the channel.
		messages <- name

		// Render the chat page.
		c.HTML(200, "chat.html", gin.H{
			"name":     name,
			"messages": messages,
		})
	})

	// Handle the `/login` route.
	r.GET("/login", func(c *gin.Context) {
		// Render the login page.
		c.HTML(200, "login.html")
	})

	// Handle the `/send` route.
	r.POST("/send", func(c *gin.Context) {
		// Get the message from the request body.
		message := c.PostForm("message")

		// Send the message to the channel.
		messages <- message
	})

	// Listen on port 8080.
	r.Run(":8080")
}
