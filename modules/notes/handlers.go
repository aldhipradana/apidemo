package notes

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

// CRUD API for notes
func (h *apiHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/notes/create", h.Create)
	r.GET("/notes/:id", h.GetById)
	r.POST("/notes/list", h.List)
	r.PUT("/notes/:id", h.Update)
	r.DELETE("/notes/:id", h.Delete)
}

// CreateNotes godoc
// @Tags Notes
// @Summary Create notes
// @Description Create notes
// @Accept  json
// @Produce  json
// @Param notes body schemas.NotesDTO true "Notes"
// @Success 200 {object} schemas.Notes
// @Router /notes/create [post]
func (h *apiHandler) Create(c *gin.Context) {
	var note schemas.Notes
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.Create(&note); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, note)
}

// GetNotesById godoc
// @Tags Notes
// @Summary Get notes by ID
// @Description Get notes by ID
// @Accept  json
// @Produce  json
// @Param id path string true "Notes ID"
// @Success 200 {object} schemas.Notes
// @Router /notes/{id} [get]
func (h *apiHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	uid := uuid.FromStringOrNil(id)

	note, err := h.uc.GetById(uid)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, note)
}

// ListNotes godoc
// @Tags Notes
// @Summary List notes
// @Description List notes
// @Accept  json
// @Produce  json
// @Success 200 {array} schemas.Notes
// @Router /notes/list [post]
func (h *apiHandler) List(c *gin.Context) {
	notes, err := h.uc.List()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, notes)
}

// UpdateNotes godoc
// @Tags Notes
// @Summary Update notes
// @Description Update notes
// @Accept  json
// @Produce  json
// @Param id path string true "Notes ID"
// @Param notes body schemas.NotesDTO true "Notes"
// @Success 200 {object} schemas.Notes
// @Router /notes/{id} [put]
func (h *apiHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	uid := uuid.FromStringOrNil(id)

	var note schemas.Notes
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	note.ID = uid

	if err := h.uc.Update(&note); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, note)
}

// DeleteNotes godoc
// @Tags Notes
// @Summary Delete notes
// @Description Delete notes
// @Accept  json
// @Produce  json
// @Param id path string true "Notes ID"
// @Success 200 {object} string
// @Router /notes/{id} [delete]
func (h *apiHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	uid := uuid.FromStringOrNil(id)

	if err := h.uc.Delete(uid); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}
