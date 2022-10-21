package boot

import (
	"DJ-Blog/pkg/search"
	"DJ-Blog/pkg/viperlib"
)

func createIndex() {
	search.ZincCli.CreateIndex(viperlib.GetString("zinc.indexName"), &search.ZincIndexPropertyMap{
		"id": &search.ZincIndexProperty{
			Type:     "numeric",
			Index:    true,
			Store:    true,
			Sortable: true,
		},
		"created_at": &search.ZincIndexProperty{
			Type:     "text",
			Index:    true,
			Sortable: true,
		},
		"updated_at": &search.ZincIndexProperty{
			Type:     "text",
			Index:    true,
			Sortable: true,
		},
		"title": &search.ZincIndexProperty{
			Type:           "text",
			Index:          true,
			Store:          true,
			Aggregatable:   true,
			Highlightable:  true,
			Analyzer:       "gse_search",
			SearchAnalyzer: "gse_standard",
		},
		"tag": &search.ZincIndexProperty{
			Type:  "keyword",
			Index: true,
			Store: true,
		},
		"content": &search.ZincIndexProperty{
			Type:           "text",
			Index:          true,
			Store:          true,
			Aggregatable:   true,
			Highlightable:  true,
			Analyzer:       "gse_search",
			SearchAnalyzer: "gse_standard",
		},
		"likes": &search.ZincIndexProperty{
			Type:  "numeric",
			Index: true,
			Store: true,
		},
		"stared": &search.ZincIndexProperty{
			Type:  "boolean",
			Index: true,
			Store: true,
		},
		"user_id": &search.ZincIndexProperty{
			Type:  "numeric",
			Index: true,
			Store: true,
		},
	})
}
