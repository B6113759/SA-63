package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/B6113759/app/ent"
	"github.com/B6113759/app/ent/relative"
	"github.com/gin-gonic/gin"
)

// RelativeController defines the struct for the relative controller
type RelativeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateRelative handles POST requests for adding relative entities
// @Summary Create relative
// @Description Create relative
// @ID create-relative
// @Accept   json
// @Produce  json
// @Param relative body ent.Relative true "Relative entity"
// @Success 200 {object} ent.Relative
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /relatives [post]
func (ctl *RelativeController) CreateRelative(c *gin.Context) {
	obj := ent.Relative{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "relative binding failed",
		})
		return
	}

	u, err := ctl.client.Relative.
		Create().
		SetRelativeName(obj.RelativeName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetRelative handles GET requests to retrieve a relative entity
// @Summary Get a relative entity by ID
// @Description get relative by ID
// @ID get-relative
// @Produce  json
// @Param id path int true "Relative ID"
// @Success 200 {object} ent.Relative
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /relatives/{id} [get]
func (ctl *RelativeController) GetRelative(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Relative.
		Query().
		Where(relative.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListRelative handles request to get a list of relative entities
// @Summary List relative entities
// @Description list relative entities
// @ID list-relative
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Relative
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /relatives [get]
func (ctl *RelativeController) ListRelative(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	relatives, err := ctl.client.Relative.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, relatives)
}

// DeleteRelative handles DELETE requests to delete a relative entity
// @Summary Delete a relative entity by ID
// @Description get relative by ID
// @ID delete-relative
// @Produce  json
// @Param id path int true "Relative ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /relatives/{id} [delete]
func (ctl *RelativeController) DeleteRelative(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Relative.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateRelative handles PUT requests to update a relative entity
// @Summary Update a relative entity by ID
// @Description update relative by ID
// @ID update-relative
// @Accept   json
// @Produce  json
// @Param id path int true "Relative ID"
// @Param relative body ent.Relative true "Relative entity"
// @Success 200 {object} ent.Relative
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /relatives/{id} [put]
func (ctl *RelativeController) UpdateRelative(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Relative{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "relative binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.Relative.
		UpdateOneID(int(id)).
		SetRelativeName(obj.RelativeName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewRelativeController creates and registers handles for the relative controller
func NewRelativeController(router gin.IRouter, client *ent.Client) *RelativeController {
	uc := &RelativeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitRelativeController registers routes to the main engine
func (ctl *RelativeController) register() {
	relatives := ctl.router.Group("/relatives")

	relatives.GET("", ctl.ListRelative)

	// CRUD
	relatives.POST("", ctl.CreateRelative)
	relatives.GET(":id", ctl.GetRelative)
	relatives.PUT(":id", ctl.UpdateRelative)
	relatives.DELETE(":id", ctl.DeleteRelative)
}
