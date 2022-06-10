package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Comment struct {
	Id     uint   `json:"id"`
	PostId uint   `json:"post_id"`
	Text   string `json:"text"`
}

func main() {

	db, err := gorm.Open(mysql.Open("root:toluwase@tcp(127.0.0.1:3306)/comments_ms"), &gorm.Config{})
	if err != nil {
		panic("database error")
	}
	db.AutoMigrate(Comment{})

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/api/posts/:id/comments", func(c *fiber.Ctx) error {
		var comments []Comment

		db.Find(&comments, "post_id = ?", c.Params("id"))
		if err := c.BodyParser(&comments); err != nil {
			log.Println(err)
			return err
		}

		db.Create(&comments)
		return c.JSON(comments)
	})

	app.Post("/api/comments", func(c *fiber.Ctx) error {
		var comment Comment

		if err := c.BodyParser(&comment); err != nil {
			log.Println(err)
			return err
		}
		db.Create(&comment)
		return c.JSON(comment)
	})

	app.Listen(":8001")
}
