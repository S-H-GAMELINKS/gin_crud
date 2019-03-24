package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// Create Todo Model
type Todo struct(
	gorm.Model
	text	string
	status	string
)

// DB Init
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")

	if err != nil {
		panic("Can not DB Open!")
	}

	db.AutoMigrate(&Todo{})
	defer db.Close()
}

// DB Insert
func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")

	if err != nil {
		panic("Can not Insert DB!")
	}

	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

// DB update
func dbUpdate(id int, text string, status string) {
	db, err := grom.Open("sqlite3", "test.sqlite3")

	if err ! nil {
		panic("Can not Update DB!")
	}

	var todo Todo
	db.First(&todo, id)

	todo.Text = text
	todo.Status = status

	db.save(&todo)
	db.Close()
}

// DB delete
func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")

	if err != nil {
		panic("Can not Delete DB!")
	}

	var todo Todo

	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

// DB all
func dbGetAll() {
	db, err := gorm,Open("sqlite3", "test.sqlite3")

	if err != nil {
		panic("Can not Get All DB!")
	}

	var todos []Todo

	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("static/*.html")

	data := "Hello Go/Gin!"

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"data": data})
	})
	r.Run(":8000")
}