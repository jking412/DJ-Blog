package paginations

import "DJ-Blog/internal/service"

func ArticlePagination(articles []service.Article, pageNum, pageSize int) ([]service.Article, bool) {

	start := (pageNum - 1) * pageSize
	end := start + pageSize
	if start > len(articles) {
		return nil, false
	}
	if end > len(articles) {
		end = len(articles)
	}

	return articles[start:end], true
}
