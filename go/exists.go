package diccionario

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ExistsResponse is the response sent back for the exists endpoint.
type ExistsResponse struct {
	// Exists is true if the word exists; otherwise, false.
	Exists bool `json:"exists"`
}

// WordExists returns true if the word exists in the word list.
// It performs case insensitive matching to the words in the wordlist.
func (s *Server) WordExists(c *gin.Context) {
	word := c.Param("word")

	log.Println("checking if word exists:", word)

	wordlist, err := s.w.GetWords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiError{Err: err, Desc: "unable to retrieve word list"})
		return
	}

	resp := ExistsResponse{Exists: false}

	// Convert to lowercase for case-insensitive comparison
	wordLower := strings.ToLower(word)

	// TODO: use word exists function - use a better name
	for _, w := range wordlist {
		if strings.ToLower(w) == wordLower {
			resp.Exists = true
			break // Found exact match, no need to continue
		}
	}

	c.JSON(http.StatusOK, resp)
}
