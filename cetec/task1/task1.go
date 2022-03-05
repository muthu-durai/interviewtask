package main

import (
	"database/sql"
	"fmt"
	_ "time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

var dbIp string
var dbUserName string
var dbPassword string

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(cors.Default())
	r.GET("/person/:person_id/info", func(c *gin.Context) {
		person_id := c.Param("person_id")
		type person struct {
			Name         string `json:"name"`
			Phone_number string `json:"phone_number"`
			City         string `json:"city"`
			State        string `json:"state"`
			Street1      string `json:"street1"`
			Street2      string `json:"street2"`
			Zip_code     string `json:"zip_code"`
		}
		var v person
		sqlRaw := fmt.Sprintf("select (select name from cetec.person as a  where a.id=ua.id) as name,(select number from cetec.phone as b where b.person_id=ua.person_id )as number,(select city  from cetec.address as u where u.id=ua.address_id) as city,(select state from cetec.address as u where u.id=ua.address_id) as state,(select street1 from cetec.address as u where u.id=ua.address_id) as street1,(select street2 from cetec.address  as u where u.id= ua.address_id) as street2,(select zip_code from cetec.address as u where u.id=ua.address_id) as zip_code  from cetec.address_join as ua where person_id = '%s'", person_id)
		db, err := sql.Open("mysql", dbUserName+":"+dbPassword+"@tcp("+dbIp+")/cetec")

		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		results, err := db.Query(sqlRaw)
		fmt.Println(sqlRaw)
		if err != nil {
			panic(err.Error())
		}

		for results.Next() {
			err = results.Scan(&v.Name, &v.Phone_number, &v.City, &v.State, &v.Street1, &v.Street2, &v.Zip_code)
			if err != nil {
				panic(err.Error())
			}
			c.JSON(200, v)
		}
		fmt.Println(v)
		c.JSON(200, v)
	})
	return r
}
func main() {
	dbIp = "localhost"
	dbUserName = "root"
	dbPassword = "durai9087"
	r := setupRouter()
	r.Run(":8091")

}
