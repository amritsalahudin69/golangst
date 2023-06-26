package main

import (
	"belajar_golang_new/domain/entity"
	"encoding/json"
	"fmt"
)

func main() {
	jsonArticles := "{\"article\":[{\"title\":\"Mobile Legends\",\"description\":\"Desc Mobile Legends\"},{\"title\":\"Game Terbaru\",\"description\":\"Desc Game Terbaru\"}]}"
	articles := new(entity.ListArticle)
	errJson := json.Unmarshal([]byte(jsonArticles), articles)
	if errJson != nil {
		fmt.Println(errJson)
	}

	author := entity.NewAuthor("Aziz")

	author.AddListArticle(articles.Article)

	newArticle := []*entity.Article{
		{
			Title:       "Episode Terbaru Naruto",
			Description: "Desc",
		},
	}

	author.AddListArticle(newArticle)

	fmt.Println(author)

	for _, article := range author.Article {
		fmt.Println("Nama Author :", author.GetName(), ", Judul Article :", article.GetArticleTitle())
	}
}
