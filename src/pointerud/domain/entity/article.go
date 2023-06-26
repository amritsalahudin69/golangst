package entity

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ListArticle struct {
	Article []*Article `json:"article"`
}

func (a *Article) GetArticleTitle() string {
	return a.Title
}
