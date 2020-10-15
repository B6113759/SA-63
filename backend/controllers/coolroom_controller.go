package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/B6113759/app/ent"
	"github.com/B6113759/app/ent/coolroom"
	"github.com/B6113759/app/ent/coolroomtype"
	"github.com/gin-gonic/gin"
)

// CoolroomController defines the struct for the coolroom controller
type CoolroomController struct {
	client *ent.Client
	router gin.IRouter
}

// Coolroom defines the struct for the coolroom
type Coolroom struct {
	COOLROOMNAME   string
	COOMROOMTYPEID int
}

// CreateCoolroom handles POST requests for adding coolroom entities
// @Summary Create coolroom
// @Description Create coolroom
// @ID create-coolroom
// @Accept   json
// @Produce  json
// @Param coolroom body Coolroom true "Coolroom entity"
// @Success 200 {object} ent.Coolroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coolrooms [post]
func (ctl *CoolroomController) CreateCoolroom(c *gin.Context) {
	obj := Coolroom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "coolroom binding failed",
		})
		return
	}

	cr, err := ctl.client.Coolroom.
		Create().
		SetCoolroomName(obj.COOLROOMNAME).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	ct, err := ctl.client.CoolroomType.
		UpdateOneID(int(obj.COOMROOMTYPEID)).
		AddCoolrooms(cr).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving edge failed",
		})
		return
	}

	c.JSON(200, ct)
}

// GetCoolroom handles GET requests to retrieve a coolroom entity
// @Summary Get a coolroom entity by ID
// @Description get coolroom by ID
// @ID get-coolroom
// @Produce  json
// @Param id path int true "Coolroom ID"
// @Success 200 {object} ent.Coolroom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coolrooms/{id} [get]
func (ctl *CoolroomController) GetCoolroom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Coolroom.
		Query().
		WithCoolroomtype().
		Where(coolroom.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListCoolroom handles request to get a list of coolroom entities
// @Summary List coolroom entities
// @Description list coolroom entities
// @ID list-coolroom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Coolroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coolrooms [get]
func (ctl *CoolroomController) ListCoolroom(c *gin.Context) {
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

	coolrooms, err := ctl.client.Coolroom.
		Query().
		WithCoolroomtype().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, coolrooms)
}

// ListCoolroomByCoolroomType handles request to get a list of coolroom entities by coolroomtype
// @Summary List coolroom entities by coolroomtype
// @Description list coolroom entities by coolroomtype
// @ID list-coolroom-by-coolroomtype
// @Produce json
// @Param typeid  query int false "Typeid"
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Coolroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coolroomswithcoolroomtype [get]
func (ctl *CoolroomController) ListCoolroomByCoolroomType(c *gin.Context) {
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

	typeidQuery := c.Query("typeid")
	typeid := 0
	if typeidQuery != "" {
		typeid64, err := strconv.ParseInt(typeidQuery, 10, 64)
		if err == nil {
			typeid = int(typeid64)
		}
	}

	coolrooms, err := ctl.client.Coolroom.
		Query().
		Where(coolroom.HasCoolroomtypeWith(coolroomtype.IDEQ(typeid))).
		WithCoolroomtype().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, coolrooms)
}

// DeleteCoolroom handles DELETE requests to delete a coolroom entity
// @Summary Delete a coolroom entity by ID
// @Description get coolroom by ID
// @ID delete-coolroom
// @Produce  json
// @Param id path int true "Coolroom ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coolrooms/{id} [delete]
func (ctl *CoolroomController) DeleteCoolroom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Coolroom.
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

// UpdateCoolroom handles PUT requests to update a coolroom entity
// @Summary Update a coolroom entity by ID
// @Description update coolroom by ID
// @ID update-coolroom
// @Accept   json
// @Produce  json
// @Param id path int true "Coolroom ID"
// @Param coolroom body ent.Coolroom true "Coolroom entity"
// @Success 200 {object} ent.Coolroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /coolrooms/{id} [put]
func (ctl *CoolroomController) UpdateCoolroom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := Coolroom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "coolroom binding failed",
		})
		return
	}
	ct, err := ctl.client.CoolroomType.
		Query().
		Where(coolroomtype.IDEQ(int(obj.COOMROOMTYPEID))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	cr, err := ctl.client.Coolroom.
		UpdateOneID(int(id)).
		SetCoolroomName(obj.COOLROOMNAME).
		SetCoolroomtype(ct).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "update video failed",
		})
		return
	}

	c.JSON(200, cr)
}

// NewCoolroomController creates and registers handles for the coolroom controller
func NewCoolroomController(router gin.IRouter, client *ent.Client) *CoolroomController {
	cc := &CoolroomController{
		client: client,
		router: router,
	}
	cc.register()
	return cc
}

// InitCoolroomController registers routes to the main engine
func (ctl *CoolroomController) register() {
	coolrooms := ctl.router.Group("/coolrooms")
	coolroomswithcoolroomtype := ctl.router.Group("/coolroomswithcoolroomtype")

	coolrooms.GET("", ctl.ListCoolroom)

	// CRUD
	coolrooms.POST("", ctl.CreateCoolroom)
	coolrooms.GET(":id", ctl.GetCoolroom)
	coolroomswithcoolroomtype.GET("", ctl.ListCoolroomByCoolroomType)
	coolrooms.PUT(":id", ctl.UpdateCoolroom)
	coolrooms.DELETE(":id", ctl.DeleteCoolroom)
}
