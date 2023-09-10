package handlers

import (
	"database/sql"
	"go-htmx/pkg/models"
)

type PostService struct {
	Db *sql.DB
}

func NewPostService(db *sql.DB) PostService {
	return PostService{Db: db}
}

func (this *PostService) GetAll() (*[]models.Post, error) {
	rows, err := this.Db.Query(`
		select 
			Id, 
			Title, 
			Content, 
			(
				select count(*) from Post_Likes pl where pl.Id_Post = p.Id = pl.Id_Post  
			)	as Likes_Count

		from Posts p;
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []models.Post{}

	for rows.Next() {
		post := models.Post{}

		if err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Likes_Count); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return &posts, nil
}

func (service *PostService) Like(Id int, IdUser int) error {
	var likeId int = -1

	if err := service.Db.QueryRow(`select Id from Post_Likes where Id_Post = ? and Id_User = ?`, Id, IdUser).Scan(&likeId); err != nil {
		if err.Error() != "sql: no rows in result set" {
			return err
		}
	}

	if likeId == -1 {
		_, err := service.Db.Query(`insert into Post_Likes (Id_Post, Id_User) values (?, ?)`, Id, IdUser)
		if err != nil {
			return err
		}
	} else {
		_, err := service.Db.Query(`delete from Post_Likes where Id = ?`, likeId)

		if err != nil {
			return err
		}
	}

	return nil
}

func (service *PostService) Create(post struct {
	Title   string
	Content string
}) (*models.Post, error) {

	stmt, err := service.Db.Prepare(`insert into Posts (Title, Content) values (?, ?)`)

	defer stmt.Close()

	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(post.Title, post.Content)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &models.Post{Id: int(id), Title: post.Title, Content: post.Content, Likes_Count: 0}, nil
}

func (service *PostService) LikesCount(Id int) (*int, error) {
	var count int = 0

	err := service.Db.QueryRow(`select count(*) Post_Likes from Post_Likes where Id_Post = ?`, Id).Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}
