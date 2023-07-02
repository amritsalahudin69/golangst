package entity

type Author struct {
	name    string
	Article []*Article
}

func NewAuthor(name string) *Author {
	return &Author{
		name: name,
	}
}

func (p *Author) AddListArticle(articles []*Article) *Author {
	for _, article := range articles {
		p.Article = append(p.Article, article)
	}
	return p
}

func (p *Author) GetName() string {
	return p.name
}

//Tabah Salahudin Amri-QUIZ2_30062023
