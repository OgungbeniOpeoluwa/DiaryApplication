package controller

import (
	"Diary/src/dto/request"
	"Diary/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var diary service.DiaryService = service.NewDiaryServiceImpl()

func Runner() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", func(context *gin.Context) {
		var register request.RegisterRequest
		err := context.BindJSON(&register)
		if err != nil {
			return
		}
		diaryMessage, errors := diary.CreateDiary(register.Username, register.Password)
		if errors != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": errors.Error(),
			})
		} else {
			context.JSON(http.StatusAccepted, gin.H{
				"message": diaryMessage.Username(),
			})
		}

	})
	r.POST("/entry", func(context *gin.Context) {
		var newEntry request.AddEntryRequest
		err := context.BindJSON(&newEntry)
		if err != nil {
			return
		}
		diary.AddEntry(newEntry.Username, newEntry.Title, newEntry.Body)

		context.JSON(http.StatusAccepted, gin.H{
			"message": "Entry Created",
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
