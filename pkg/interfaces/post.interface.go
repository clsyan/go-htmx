package interfaces

import "go-htmx/pkg/models"

type PostService interface {
	GetAll() ([]models.Post, error)
	Like(Id int, IdUser int) error
	Create(post struct {
		Title   string
		Content string
	}) (*models.Post, error)
	LikesCount(Id int) (int, error)
}
