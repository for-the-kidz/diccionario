package diccionario

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddRequest represents the request body for the /add endpoint.
type AddRequest struct {
	// Word is the word to add to the word list.
	Word string `json:"word"`
}

// Add a new word to the word list.
func (s *Server) Add(c *gin.Context) {
	var req AddRequest
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// implement your logic here
}
