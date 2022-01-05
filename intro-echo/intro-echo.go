package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
)

type Account struct {
	ID          int
	Username    string
	Displayname string
	Location    string
}

type Article struct {
	ID      int    `json:"id" form:"id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

type Response struct {
	Status  bool
	Code    int
	Message string
	Data    []Article
}

var data = []Article{
	{1, "Tahanan lepas", "tahanan lapas, kabur dengan mobdin"},
	{2, "Tahun baru sepi", "tahun baru disurabaya dikabarkan sepi"},
	{3, "Tahun baru sepi1", "1tahun baru disurabaya dikabarkan sepi"},
}

var arrAccount = []Account{}

func delete(id int) {
	fmt.Println(data)
	fmt.Println("===============")
	// id yang di delete 2
	newArr := data[:id-1]
	for i := id; i < len(data); i++ {
		newArr = append(newArr, data[i])
	}
	data = newArr
	// map [int] string
	// 1 : jerry | 2 : arif | 3 : Andrew --> remove 2 (data arif hilang)
	// jika akses tinggal map [1] -> jerry

	// array []user (user : id, nama, alamat)
	// 0 : id 1, | 1 : id 2 | 2 : id 3 --> remove id : 2
	// 0 : id 1 | 1 : id 3
	fmt.Println(data)
}

func GetArticlesbyID(c echo.Context) error {
	sid := c.Param("id")
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println(sid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ID seharusnya angka")
	}
	resData := []Article{}
	resData = append(resData, data[id-1])
	resp := Response{true, http.StatusOK, "Berhasil Get Article by ID", resData}

	return c.JSON(http.StatusOK, resp)
}

func GetByQuery(c echo.Context) error {
	parameter := c.QueryParam("content")
	parameter2 := c.QueryParam("title")

	return c.JSON(http.StatusOK, parameter+" "+parameter2)
}

func BindDataParameter(c echo.Context) error {
	article := Article{}
	c.Bind(&article)

	if strings.Contains(article.Content, "halo") {

	}

	return c.JSON(http.StatusOK, article)
}

func GetFromDB(c echo.Context) error {
	db, err := sql.Open("mysql", "root:@/be5db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM account")

	// qry1 := "Insert into account values (6, 'ela10', 'ela', 'indonesia')"

	// results2, err := db.Exec(qry1)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "There is some error")
	}

	for results.Next() {
		tmp := Account{}
		if err := results.Scan(&tmp.ID, &tmp.Username, &tmp.Displayname, &tmp.Location); err != nil {
			return c.JSON(http.StatusInternalServerError, "There is something wrong")
		}
		arrAccount = append(arrAccount, tmp)
	}

	return c.JSON(http.StatusOK, arrAccount)
}

func main() {
	// delete(2)
	// delete(3)
	e := echo.New()

	e.GET("/articles", func(c echo.Context) error {
		if len(data) > 0 {
			return c.JSON(http.StatusOK, data)
		}
		return c.JSON(http.StatusInternalServerError, "Ada masalah bro")
	})

	e.GET("/articles/:id", GetArticlesbyID)
	e.GET("/article", GetByQuery)
	e.POST("/article", BindDataParameter)
	e.GET("/accounts", GetFromDB)

	e.Start(":8080")
}
