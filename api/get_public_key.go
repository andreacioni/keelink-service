package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getPublicKey(c *gin.Context) {
	entry, found := getEntryFromSessionID(c)

	if !found {
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": entry.PublicKey})
}
