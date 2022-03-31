package main

import (
	"fmt"
	"net/http"
	_ "strconv"

	"github.com/gin-gonic/gin"
)

type ErrorRResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) CreateEmployee(c *gin.Context) {

	var employee Employee

	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorRResponse{
			Message: err.Error(),
		})
		return
	}
	h.storage.Insert(&employee)
	c.JSON(http.StatusOK, map[string]int{
		"id": employee.Id,
	})
}
