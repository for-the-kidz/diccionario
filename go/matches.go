package diccionario

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

// MatchesResponse is the response sent back for the prefix matches endpoint.
type MatchesResponse struct {
	// Matches contains the list of words that match the given prefix.
	Matches []string `json:"matches"`
}

// Matches returns a list of words that matched the given prefix.
// It performs case insensitive matching to the words in the wordlist.
func (s *Server) Matches(c *gin.Context) {
	prefix := strings.TrimSpace(c.Param("prefix"))

	// Validate prefix is not empty
	if prefix == "" {
		c.String(http.StatusBadRequest, "prefix cannot be empty")
		return
	}

	// Validate prefix contains only alphabets
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`)
	if !isAlpha.MatchString(prefix) {
		c.String(http.StatusBadRequest, "prefix must contain only alphabets")
		return
	}

	// Get wordlist
	wordlist, err := s.w.GetWords()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Convert prefix to lowercase for case-insensitive matching
	prefixLower := strings.ToLower(prefix)
	resp := MatchesResponse{Matches: make([]string, 0)}

	// Find all words that match the prefix (case-insensitive)
	for _, w := range wordlist {
		if strings.HasPrefix(strings.ToLower(w), prefixLower) {
			resp.Matches = append(resp.Matches, w)
		}
	}

	c.JSON(http.StatusOK, resp)
}
