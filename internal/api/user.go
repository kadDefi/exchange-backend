package api

import (
	"exchange-backend/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (b *Backend) queryUserNFT(c *gin.Context) {
	query := domain.QueryNFTItemArg{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rows, err := b.service.QueryUserNFT(c, &query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rows": rows,
	})
}
