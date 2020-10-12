package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/B6113759/app/ent"
	"github.com/B6113759/app/ent/coolroom"
	"github.com/B6113759/app/ent/patient"
	"github.com/B6113759/app/ent/relative"
	"github.com/B6113759/app/ent/user"
	"github.com/gin-gonic/gin"
)

// DeceasedReceiveController defines the struct for the deceasedreceive controller
type DeceasedReceiveController struct {
	client *ent.Client
	router gin.IRouter
}

// DeceasedReceive defines the struct for the deceasedreceive
type DeceasedReceive struct {
	RelativeID int
	CoolroomID int
	UserID     int
	PatientID  int
	Deathtime  string
}

// CreateDeceasedReceive handles POST requests for adding deceasedreceive entities
// @Summary Create deceasedreceive
// @Description Create deceasedreceive
// @ID create-deceasedreceive
// @Accept   json
// @Produce  json
// @Param deceasedreceive body DeceasedReceive true "DeceasedReceive entity"
// @Success 200 {object} ent.DeceasedReceive
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /deceasedreceives [post]
func (ctl *DeceasedReceiveController) CreateDeceasedReceive(c *gin.Context) {
	obj := DeceasedReceive{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "deceasedreceive binding failed",
		})
		return
	}

	cr, err := ctl.client.Coolroom.
		Query().
		Where(coolroom.IDEQ(int(obj.CoolroomID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "coolroom not found",
		})
		return
	}

	r, err := ctl.client.Relative.
		Query().
		Where(relative.IDEQ(int(obj.RelativeID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "relative not found",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.UserID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.PatientID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	times, err := time.Parse(time.RFC3339, obj.Deathtime)

	dr, err := ctl.client.DeceasedReceive.
		Create().
		SetDoctor(u).
		SetPatient(p).
		SetCoolroom(cr).
		SetRelative(r).
		SetDeathTime(times).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, dr)
}

// ListDeceasedReceive handles request to get a list of deceasedreceive entities
// @Summary List deceasedreceive entities
// @Description list deceasedreceive entities
// @ID list-deceasedreceive
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.DeceasedReceive
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /deceasedreceives [get]
func (ctl *DeceasedReceiveController) ListDeceasedReceive(c *gin.Context) {
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

	deceasedreceives, err := ctl.client.DeceasedReceive.
		Query().
		WithCoolroom().
		WithDoctor().
		WithRelative().
		WithPatient().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, deceasedreceives)
}

// DeleteDeceasedReceive handles DELETE requests to delete a deceasedreceive entity
// @Summary Delete a deceasedreceive entity by ID
// @Description get deceasedreceive by ID
// @ID delete-deceasedreceive
// @Produce  json
// @Param id path int true "DeceasedReceive ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /deceasedreceives/{id} [delete]
func (ctl *DeceasedReceiveController) DeleteDeceasedReceive(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.DeceasedReceive.
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

// NewDeceasedReceiveController creates and registers handles for the deceasedreceive controller
func NewDeceasedReceiveController(router gin.IRouter, client *ent.Client) *DeceasedReceiveController {
	drc := &DeceasedReceiveController{
		client: client,
		router: router,
	}
	drc.register()
	return drc
}

// InitDeceasedReceiveController registers routes to the main engine
func (ctl *DeceasedReceiveController) register() {
	deceasedreceives := ctl.router.Group("/deceasedreceives")

	deceasedreceives.GET("", ctl.ListDeceasedReceive)

	// CRUD
	deceasedreceives.POST("", ctl.CreateDeceasedReceive)
	deceasedreceives.DELETE(":id", ctl.DeleteDeceasedReceive)
}
