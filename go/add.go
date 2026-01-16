package diccionario

import (
	"net/http"
	"regexp"
	"strings"

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

	// check given input word is valid or not -> only alphabets allowed else 400
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`)
	if !isAlpha.MatchString(req.Word) {
		c.String(http.StatusBadRequest, "input must contain only alphabets")
		return
	}

	// check if word exist in word list or not -> if exists return 409
	wordlist, err := s.w.GetWords()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	for _, w := range wordlist {
		if strings.ToLower(w) == strings.ToLower(req.Word) {
			c.String(http.StatusConflict, "word already exists")
			return
		}
	}

	// if not exists add the word to word list and return 204
	err = s.w.AddWord(req.Word)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusNoContent)

}

// This endpoint adds a new word to the word list. This endpoint needs to be implemented still.

// Expected functionality:

// It returns a 204 upon success.
// It returns a 409 if the word already exists in the word list.
// It returns other status codes as appropriate (4XXs for input errors, 5XXs for internal server errors)
// The newly added word should persist for the life of the running Docker container.
// A word is considered a single string of unbroken alpha characters (no numbers or special characters)