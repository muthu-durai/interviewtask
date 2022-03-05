package main

import (
	"fmt"
	_ "time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(cors.Default())
	r.POST("/person/create", func(c *gin.Context) {

		type User struct {
			Name         string `json:"name"`
			Phone_number string `json:"phone_number"`
			City         string `json:"city"`
			State        string `json:"state"`
			Street1      string `json:"street1"`
			Street2      string `json:"street2"`
			Zip_code     string `json:"zip_code"`
		}
		var user1 User
		err := c.ShouldBindJSON(&user1)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Print(&user1)
		c.JSON(200, user1)

	})
	return r
}
func main() {

	r := setupRouter()
	r.Run(":8091")

}
