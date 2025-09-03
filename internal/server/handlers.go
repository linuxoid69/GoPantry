package server

import (
	"github.com/gin-gonic/gin"
)

func HandlerAddMovie(ctx *gin.Context) {
    var movie Movie

    if err := ctx.BindJSON(&movie); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }
    // TODO: here add to DB
    ctx.JSON(200, gin.H{"message": "Movie added successfully"})
}
