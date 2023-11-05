package user

import (
	"strconv"

	"github.com/aldhipradana/apidemo/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type apiHandler struct {
	uc Usecases
}

func NewHandler(db *gorm.DB) *apiHandler {
	return &apiHandler{Usecases{db}}
}

func (h *apiHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/user/create", h.Create)
	r.GET("/user/:id", h.GetById)
	r.POST("/user/list", h.List)
	r.PUT("/user/:id", h.Update)
	r.DELETE("/user/:id", h.Delete)
}

// CreateUser godoc
// @Summary Create user
// @Description Create user
// @Accept  json
// @Produce  json
// @Param user body schemas.User true "User"
// @Success 200 {object} schemas.User
// @Router /user/create [post]
func (h *apiHandler) Create(c *gin.Context) {
	var user schemas.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.Create(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

// GetUserById godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} schemas.User
// @Router /user/{id} [get]
func (h *apiHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	user, err := h.uc.GetById(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

// ListUsers godoc
// @Summary List users
// @Description List users
// @Accept  json
// @Produce  json
// @Success 200 {array} schemas.User
// @Router /user/list [post]
func (h *apiHandler) List(c *gin.Context) {
	users, err := h.uc.List()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body schemas.User true "User"
// @Success 200 {object} schemas.User
// @Router /user/{id} [put]
func (h *apiHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var user schemas.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	uid, e2 := strconv.ParseUint(id, 10, 64)
	if e2 != nil {
		c.JSON(400, gin.H{"error": e2.Error()})
		return
	}

	user.ID = uint(uid)
	if err := h.uc.Update(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} string
// @Router /user/{id} [delete]
func (h *apiHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	uid, e2 := strconv.ParseUint(id, 10, 64)
	if e2 != nil {
		c.JSON(400, gin.H{"error": e2.Error()})
		return
	}

	if err := h.uc.Delete(uint(uid)); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "OK"})
}
