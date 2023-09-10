package controllers

import (
	"database/sql"
	dto "go-htmx/pkg/dto/posts"
	"go-htmx/pkg/handlers"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var pool *sql.DB // Database connection pool.

type PostController struct {
	PostService *handlers.PostService
}

func (this PostController) Route(app *fiber.App) {
	app.Post("posts/", this.Create)
	app.Get("posts/", this.GetAll)
	app.Put("posts/:id/likes", this.Like)
}

func NewPostController(postService *handlers.PostService) *PostController {
	return &PostController{PostService: postService}
}

func (this PostController) GetAll(c *fiber.Ctx) error {
	posts, _ := this.PostService.GetAll()

	return c.Render("posts", fiber.Map{
		"Posts": posts,
	})
}

func (this PostController) Create(c *fiber.Ctx) error {
	p := new(dto.Post)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	post, err := this.PostService.Create(struct {
		Title   string
		Content string
	}{Title: p.Title, Content: p.Content})

	if err != nil {
		return err
	}

	return c.Render("post-in-feed", post)
}

func (this PostController) Like(c *fiber.Ctx) error {
	postId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Fatal(err)
	}

	err = this.PostService.Like(postId, 1)
	if err != nil {
		log.Fatal(err)
	}

	count, err := this.PostService.LikesCount(postId)
	if err != nil {
		log.Fatal(err)
	}

	return c.Render("like-button", fiber.Map{
		"Id":          postId,
		"Likes_Count": count,
	})
}
