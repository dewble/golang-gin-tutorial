package models

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "Title1", Content: "Content 1"},
	article{ID: 2, Title: "Title2", Content: "Content 2"},
}

// return a list of all the articles
func getAllArticles() []article {
	return articleList
}
