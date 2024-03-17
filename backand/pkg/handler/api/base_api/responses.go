package baseApi

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Response401(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, NewErrorResponse(err))
}

func Response403(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, NewErrorResponse(err))
}

func Response404(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, NewErrorResponse(err))
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{Error: err.Error()}
}

func GetPathID(c *gin.Context) (int, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.New("id должен быть числом!")
	}
	return id, err
}
