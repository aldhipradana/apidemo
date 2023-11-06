package user

import (
	"github.com/aldhipradana/apidemo/schemas"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
// @Tags User
// @Summary Create user
// @Description Create user
// @Accept  json
// @Produce  json
// @Param user body schemas.UserDTO true "User"
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
// @Tags User
// @Summary Get user by ID
// @Description Get user by ID
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} schemas.User
// @Router /user/{id} [get]
func (h *apiHandler) GetById(c *gin.Context) {
	id := uuid.FromStringOrNil(c.Param("id"))
	user, err := h.uc.GetById(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

// ListUsers godoc
// @Tags User
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
// @Tags User
// @Summary Update user
// @Description Update user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body schemas.UserDTO true "User"
// @Success 200 {object} schemas.User
// @Router /user/{id} [put]
func (h *apiHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var user schemas.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.ID = uuid.FromStringOrNil(id)
	if err := h.uc.Update(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

// DeleteUser godoc
// @Tags User
// @Summary Delete user
// @Description Delete user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} string
// @Router /user/{id} [delete]
func (h *apiHandler) Delete(c *gin.Context) {
	id := uuid.FromStringOrNil(c.Param("id"))

	if err := h.uc.Delete(id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "OK"})
}
