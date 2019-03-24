package main

import (
	"strconv"
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

	dbInit()

	data := "Hello Go/Gin!"

	// Index
	r.GET("/", func(c *gin.Context) {
		todos := dbGetAll()
		c.HTML(200, "index.html", gin.H{"todos": todos})
	})

	// Create
	r.POST("/new", func(c *gin.Context) {
		text := c.PostForm("text")
		status := c.PostForm("status")
		dbInsert(text, status)
		c.Redirect(302, "/")
	})

    //Detail
    r.GET("/detail/:id", func(c *gin.Context) {
        n := c.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic(err)
        }
        todo := dbGetOne(id)
        c.HTML(200, "detail.html", gin.H{"todo": todo})
    })

    //Update
    r.POST("/update/:id", func(c *gin.Context) {
        n := c.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        text := c.PostForm("text")
        status := c.PostForm("status")
        dbUpdate(id, text, status)
        c.Redirect(302, "/")
	})
	
    // Delete Chack
    r.GET("/delete_check/:id", func(c *gin.Context) {
        n := c.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        todo := dbGetOne(id)
        c.HTML(200, "delete.html", gin.H{"todo": todo})
    })

    //Delete
    r.POST("/delete/:id", func(c *gin.Context) {
        n := c.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        dbDelete(id)
        c.Redirect(302, "/")

    })	

	r.Run(":8000")
}