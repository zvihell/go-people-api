package controllers

import (
	"go-people-api/internal/models"
	"go-people-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService service.UserService
}

func NewHandler(userService service.UserService) *Handler {
	return &Handler{
		userService: userService}
}

func (h *Handler) InitRouter() *gin.Engine {
	r := gin.New()

	r.POST("/people", h.createPeople)
	r.GET("/people", h.getByName)
	r.DELETE("/people/:id", h.deletePeople)
	r.PUT("/people/:id", h.updatePeople)

	return r

}

func (h *Handler) createPeople(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	richUser := getRichorDieTrying(user)
	err := h.userService.Create(richUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, richUser)

}

func (h *Handler) getByName(c *gin.Context) {
	name := c.Query("name")
	users, err := h.userService.Get(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)

}

func (h *Handler) deletePeople(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.userService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "OK")

}

func (h *Handler) updatePeople(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id param")
		return
	}
	var todo models.User
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.userService.Update(id, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, "OK")

}
