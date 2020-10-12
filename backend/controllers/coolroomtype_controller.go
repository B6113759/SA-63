package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/B6113759/app/ent"
	"github.com/B6113759/app/ent/coolroomtype"
	"github.com/gin-gonic/gin"
)

// CoolroomTypeController defines the struct for the coolroomtype controller
type CoolroomTypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateCoolroomType handles POST requests for adding coolroomtype entities
// @Summary Create coolroomtype
// @Description Create coolroomtype
// @ID create-coolroomtype
// @Accept   json
// @Produce  json
// @Param coolroomtype body ent.CoolroomType true "CoolroomType entity"
// @Success 200 {object} ent.CoolroomType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coolroomtypes [post]
func (ctl *CoolroomTypeController) CreateCoolroomType(c *gin.Context) {
	obj := ent.CoolroomType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "coolroomtype binding failed",
		})
		return
	}

	u, err := ctl.client.CoolroomType.
		Create().
		SetCoolroomtypeName(obj.CoolroomtypeName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetCoolroomType handles GET requests to retrieve a coolroomtype entity
// @Summary Get a coolroomtype entity by ID
// @Description get coolroomtype by ID
// @ID get-coolroomtype
// @Produce  json./main
// @Param id path int true "CoolroomType ID"
// @Success 200 {object} ent.CoolroomType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coolroomtypes/{id} [get]
func (ctl *CoolroomTypeController) GetCoolroomType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.CoolroomType.
		Query().
		Where(coolroomtype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListCoolroomType handles request to get a list of coolroomtype entities
// @Summary List coolroomtype entities
// @Description list coolroomtype entities
// @ID list-coolroomtype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.CoolroomType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coolroomtypes [get]
func (ctl *CoolroomTypeController) ListCoolroomType(c *gin.Context) {
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

	coolroomtypes, err := ctl.client.CoolroomType.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, coolroomtypes)
}

// DeleteCoolroomType handles DELETE requests to delete a coolroomtype entity
// @Summary Delete a coolroomtype entity by ID
// @Description get coolroomtype by ID
// @ID delete-coolroomtype
// @Produce  json
// @Param id path int true "CoolroomType ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coolroomtypes/{id} [delete]
func (ctl *CoolroomTypeController) DeleteCoolroomType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.CoolroomType.
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

// UpdateCoolroomType handles PUT requests to update a coolroomtype entity
// @Summary Update a coolroomtype entity by ID
// @Description update coolroomtype by ID
// @ID update-coolroomtype
// @Accept   json
// @Produce  json
// @Param id path int true "CoolroomType ID"
// @Param coolroomtype body ent.CoolroomType true "CoolroomType entity"
// @Success 200 {object} ent.CoolroomType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coolroomtypes/{id} [put]
func (ctl *CoolroomTypeController) UpdateCoolroomType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.CoolroomType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "coolroomtype binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.CoolroomType.
		UpdateOneID(int(id)).
		SetCoolroomtypeName(obj.CoolroomtypeName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewCoolroomTypeController creates and registers handles for the coolroomtype controller
func NewCoolroomTypeController(router gin.IRouter, client *ent.Client) *CoolroomTypeController {
	uc := &CoolroomTypeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitCoolroomTypeController registers routes to the main engine
func (ctl *CoolroomTypeController) register() {
	coolroomtypes := ctl.router.Group("/coolroomtypes")

	coolroomtypes.GET("", ctl.ListCoolroomType)

	// CRUD
	coolroomtypes.POST("", ctl.CreateCoolroomType)
	coolroomtypes.GET(":id", ctl.GetCoolroomType)
	coolroomtypes.PUT(":id", ctl.UpdateCoolroomType)
	coolroomtypes.DELETE(":id", ctl.DeleteCoolroomType)
}
