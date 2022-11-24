package main

import (
	"github.com/gin-gonic/gin"
)

import "net/http"

type UserUri struct {
	Name string `uri:"name" binding:"required"`
	Hoge string `uri:"hoge" binding:"required"`
}

func main() {
	persons := []Parson{
		{
			name:   "makiko",
			weight: 49,
			height: 156,
		},
		{
			name:   "haruhi",
			weight: 10,
			height: 80,
		},
		{
			name:   "nora",
			weight: 65,
			height: 175,
		},
	}

	engine := gin.Default()
	engine.GET("/:name/:hoge", func(c *gin.Context) {
		uri := UserUri{}
		c.BindUri(&uri)
		isExist := false
		for _, person := range persons {
			if person.name == uri.Name {
				isExist = true
				c.JSON(http.StatusOK, gin.H{
					"person1": map[string]any{
						"name":       uri.Name,
						"weight":     person.weight,
						"height":     person.height,
						"additional": uri.Hoge,
					},
				})
			}
		}
		if !isExist {
			c.JSON(http.StatusOK, gin.H{"error": "no such person", uri.Name: uri.Hoge})
		}
	})
	//engine.GET("/harupi", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"person2": map[string]string{
	//			"namae":  "harupi",
	//			"weight": "15",
	//			"height": "96",
	//		}})
	//})
	engine.Run(":3000")

}

type Parson struct {
	name   string
	weight int
	height int
}
