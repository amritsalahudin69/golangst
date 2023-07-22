package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	FullName       string `json:"fullName"`
	ProfilePicture string `json:"profilePicture"`
	UserName       string `json:"userName"`
}

type ImageArticle struct {
	Image          string `json:"image"`
	ImageDesktop   string `json:"imageDesktop"`
	ImageHighlight string `json:"imageHighlight"`
}

type Translation struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Excerpt string `json:"excerpt"`
	Content string `json:"content"`
}

type Article struct {
	ArticleID    int          `json:"articleId"`
	ImageArticle ImageArticle `json:"imageArticle"`
	Author       Author       `json:"author"`
	Translation  Translation  `json:"translation"`
}

// Getter untuk mendapatkan nilai ArticleID
func (a *Article) GetArticleID() int {
	return a.ArticleID
}

// Getter untuk mendapatkan nilai FullName dari Author
func (a *Article) GetAuthorFullName() string {
	return a.Author.FullName
}

// Setter untuk mengatur nilai FullName dari Author
func (a *Article) SetAuthorFullName(fullName string) {
	a.Author.FullName = fullName
}

// Getter untuk mendapatkan URL gambar dari ImageArticle
func (a *Article) GetImageURL() string {
	return a.ImageArticle.Image
}

// Setter untuk mengatur URL gambar pada ImageArticle
func (a *Article) SetImageURL(url string) {
	a.ImageArticle.Image = url
}

func main() {
	article := Article{
		ArticleID: 23928,
		ImageArticle: ImageArticle{
			Image:          "https://api.duniagames.co.id/api/content/upload/file/9874800061688113932.jpg",
			ImageDesktop:   "https://api.duniagames.co.id/api/content/upload/file/9074251731688113932.jpg",
			ImageHighlight: "https://api.duniagames.co.id/api/content/upload/file/9842838701688113932.jpg",
		},
		Author: Author{
			FullName:       "haris13",
			ProfilePicture: "https://api.duniagames.co.id/api/user/upload/image/223816331649748676.jpg",
			UserName:       "Haris Firmansyah",
		},
		Translation: Translation{
			Title:   "Benedetta Mobile Legends (ML) Hero: Skill, Build dan Kelebihan",
			Slug:    "hero-benedetta-mobile-legends",
			Excerpt: "Penasaran ingin tahu tentang hero Benedetta Mobile Legends (ML)? Lihat artikel Dunia Games dan pelajari tentang skills, build, dan kelebihannya.",
			Content: "<p><span style=\"font-weight: 400;\">Benedetta adalah seorang hero<span><em><span style=\"font-weight: 400;\"> Assassin<span><em>...",
		},
	}

	// Menggunakan metode getter
	fmt.Println("Article ID:", article.GetArticleID())
	fmt.Println("Author FullName:", article.GetAuthorFullName())
	fmt.Println("Image URL:", article.GetImageURL())

	// Menggunakan metode setter
	article.SetAuthorFullName("John Doe")
	article.SetImageURL("https://example.com/image.jpg")

	// Marshal struktur Article menjadi JSON
	jsonData, err := json.MarshalIndent(article, "", "\t")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}
