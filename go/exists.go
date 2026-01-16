package diccionario

import (
	"net/http"
	"regexp"
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
	// Get the word from the URL parameter and trim spaces
	word := strings.TrimSpace(c.Param("word"))
	if word == "" {
		c.String(http.StatusBadRequest, "word cannot be empty")
		return
	}

	// word is that capproprtiate or not if not 400input error it should be only alph and no numerics
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`)

	if !isAlpha.MatchString(word) {
		c.String(http.StatusBadRequest, "input must contain only alphabets")
		return	
	}

	// conver word to lowercase for case insensitive comparison
	word = strings.ToLower(word)
	wordlist, err := s.w.GetWords()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	resp := ExistsResponse{Exists: false}

	for _, w := range wordlist {
		// case insensitive comparison
		w = strings.ToLower(w)
		if w == word {
			resp.Exists = true
			break
		}
	}

	c.JSON(http.StatusOK, resp)
}

