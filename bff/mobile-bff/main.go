package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/resty.v1"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

type PostsResponse struct {
	UserID   int       `json:"userId"`
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Comments []Comment `json:"comments"`
}

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	app.Get("/posts", func(c *fiber.Ctx) error {
		response := []PostsResponse{}
		client := resty.New()
		posts := []Post{}
		_, errP := client.R().
			SetResult(&posts).
			Get(os.Getenv("POST_URL") + "/posts")

		if errP != nil {
			c.Status(500).Send([]byte(errP.Error()))
		}

		comments := []Comment{}
		_, errC := client.R().
			SetResult(&comments).
			Get(os.Getenv("COMMENT_URL") + "/comments")

		if errC != nil {
			c.Status(500).Send([]byte(errC.Error()))
		}

		for i := 0; i < len(posts); i++ {
			response = append(response, PostsResponse{
				ID:       posts[i].ID,
				UserID:   posts[i].UserID,
				Title:    posts[i].Title,
				Body:     posts[i].Body,
				Comments: filter(posts[i], comments),
			})
		}

		return c.Status(200).JSON(response)
	})

	app.Listen(":8000")
}

func filter(post Post, comments []Comment) (ret []Comment) {
	for i := 0; i < len(comments); i++ {
		if post.ID == comments[i].PostID {
			ret = append(ret, comments[i])
		}
	}
	return
}
