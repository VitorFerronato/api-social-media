package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"

	"encoding/json"
	"io"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(bodyRequest, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userID

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewPostRepository(db)

	post.ID, err = repository.CreatePost(post)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}
func GetPosts(w http.ResponseWriter, r *http.Request)    {}
func GetPostById(w http.ResponseWriter, r *http.Request) {}
func UpdatePost(w http.ResponseWriter, r *http.Request)  {}
func DeletePost(w http.ResponseWriter, r *http.Request)  {}
