package server

import (
	"github.com/gin-gonic/gin"
)

func HandlerAddMovie(ctx *gin.Context) {
    var movies Movies

    if err := ctx.BindJSON(&movies); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }
    // TODO: here add to DB
    ctx.JSON(200, gin.H{"message": "Movie added successfully"})
}
