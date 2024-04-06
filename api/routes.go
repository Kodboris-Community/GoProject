package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	db "kodboris/db/sqlc"
	"net/http"
)

type probe struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type KodborisMembers struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type KodborisMembersResponse struct {
	FirstName string       `json:"first_Name"`
	LastName  string       `json:"last_Name"`
	Comment   string       `json:"comment"`
	CreatedAt sql.NullTime `json:"created_At"`
	Status    string       `json:"status"`
}

func (server *Server) homeProbe(ctx *gin.Context) {
	// Set status code to 200
	response := probe{
		Status:  "200",
		Message: "Welcome to Kodeboris Community",
	}
	ctx.JSON(http.StatusOK, response)

}

func (server *Server) createMember(ctx *gin.Context) {
	var r KodborisMembers

	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c := db.CreateMemberParams{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Status:    "active",
	}
	fmt.Printf("get here: %v\n", r.FirstName+r.LastName)

	// Create member record into the db
	result, err := server.db.CreateMember(ctx, c)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// Return JSON response
	a := KodborisMembersResponse{
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Status:    string(result.Status),
		CreatedAt: result.CreatedAt,
	}
	fmt.Printf("Member %v has been created", a.FirstName+a.LastName)
	ctx.JSON(http.StatusOK, a)
}
