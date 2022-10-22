package search

import (
	"DJ-Blog/pkg/viperlib"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

/*
Index demo
{
    "name": "article",
    "storage_type": "disk",
    "shard_num": 1,
    "mappings": {
        "properties": {
            "title": {
                "type": "text",
                "index": true,
                "store": true,
                "highlightable": true
            },
            "content": {
                "type": "text",
                "index": true,
                "store": true,
                "highlightable": true
            },
            "status": {
                "type": "keyword",
                "index": true,
                "sortable": true,
                "aggregatable": true
            },
            "publish_date": {
                "type": "date",
                "format": "2006-01-02T15:04:05Z07:00",
                "index": true,
                "sortable": true,
                "aggregatable": true
            }
        }
    }
}

*/

/*
Response demo
{
    "took": 0,
    "timed_out": false,
    "max_score": 7.6978611753656345,
    "hits": {
        "total": {
            "value": 3
        },
        "hits": [
            {
                "_index": "games3",
                "_type": "games3",
                "_id": "bd3e67f0-679b-4aa4-b0f5-81b9dc86a26a",
                "_score": 7.6978611753656345,
                "@timestamp": "2021-10-20T04:56:39.000871Z",
                "_source": {
                    "Athlete": "DEMTSCHENKO, Albert",
                    "City": "Turin",
                    "Country": "RUS",
                    "Discipline": "Luge",
                    "Event": "Singles",
                    "Gender": "Men",
                    "Medal": "Silver",
                    "Season": "winter",
                    "Sport": "Luge",
                    "Year": 2006
                }
            },
            {
                "_index": "games3",
                "_type": "games3",
                "_id": "230349d9-72b3-4225-bac7-a8ab31af046d",
                "_score": 7.6978611753656345,
                "@timestamp": "2021-10-20T04:56:39.215124Z",
                "_source": {
                    "Athlete": "DEMTSCHENKO, Albert",
                    "City": "Sochi",
                    "Country": "RUS",
                    "Discipline": "Luge",
                    "Event": "Singles",
                    "Gender": "Men",
                    "Medal": "Silver",
                    "Season": "winter",
                    "Sport": "Luge",
                    "Year": 2014
                }
            },
            {
                "_index": "games3",
                "_type": "games3",
                "_id": "338fea31-81f2-4b56-a096-b8294fb6cc92",
                "_score": 7.671309826309841,
                "@timestamp": "2021-10-20T04:56:39.215067Z",
                "_source": {
                    "Athlete": "DEMTSCHENKO, Albert",
                    "City": "Sochi",
                    "Country": "RUS",
                    "Discipline": "Luge",
                    "Event": "Mixed Relay",
                    "Gender": "Men",
                    "Medal": "Silver",
                    "Season": "winter",
                    "Sport": "Luge",
                    "Year": 2014
                }
            }
        ]
    },
    "buckets": null,
    "error": ""
}
*/

type ZincClient struct {
	ZincEndpoint string
	ZincUsername string
	ZincPassword string
	IndexName    string
}

type ZincIndex struct {
	Name        string        `json:"name"`
	StorageType string        `json:"storage_type"`
	Mappings    *ZincIndexMap `json:"mappings"`
}

type ZincIndexMap struct {
	Properties *ZincIndexPropertyMap `json:"properties"`
}

type ZincIndexProperty struct {
	Type           string `json:"type"`
	Index          bool   `json:"index"`
	Store          bool   `json:"store"`
	Sortable       bool   `json:"sortable"`
	Aggregatable   bool   `json:"aggregatable"`
	Highlightable  bool   `json:"highlightable"`
	Analyzer       string `json:"analyzer"`
	SearchAnalyzer string `json:"search_analyzer"`
	Format         string `json:"format"`
}

type QueryResult struct {
	Took     int         `json:"took"`
	TimedOut bool        `json:"timed_out"`
	Hits     *HitsResult `json:"hits"`
}

type HitsResult struct {
	Total    *HitsResultTotal `json:"total"`
	MaxScore float64          `json:"max_score"`
	HitItems []*HitItem       `json:"hits"`
}

type HitsResultTotal struct {
	Value int64 `json:"value"`
}

type HitItem struct {
	Index     string      `json:"_index"`
	Type      string      `json:"_type"`
	ID        string      `json:"_id"`
	Score     float64     `json:"_score"`
	Timestamp time.Time   `json:"@timestamp"`
	Source    interface{} `json:"_source"`
}

type ZincIndexPropertyMap map[string]*ZincIndexProperty

var ZincCli *ZincClient

func InitZincClient() {
	ZincCli = &ZincClient{
		ZincEndpoint: viperlib.GetString("zinc.endpoint"),
		ZincUsername: viperlib.GetString("zinc.username"),
		ZincPassword: viperlib.GetString("zinc.password"),
		IndexName:    viperlib.GetString("zinc.indexName"),
	}
	logrus.Info("ZincClient init success")
}

func (zc *ZincClient) CreateIndex(indexName string, p *ZincIndexPropertyMap) bool {
	data := &ZincIndex{
		Name:        indexName,
		StorageType: "disk",
		Mappings: &ZincIndexMap{
			Properties: p,
		},
	}

	resp, err := zc.request().SetBody(data).Post("/api/index")

	if err != nil || resp.StatusCode() != http.StatusOK {
		if zc.IsExistIndex(zc.IndexName) {
			return true
		}
		logrus.Warn("create index failed", err)
		return false
	}
	return true
}

func (zc *ZincClient) IsExistIndex(indexName string) bool {
	resp, err := zc.request().Head(fmt.Sprintf("/api/index/%s", indexName))

	if err != nil || resp.StatusCode() != http.StatusOK {
		logrus.Warn("get index failed", err)
		return false
	}

	return true
}

func (zc *ZincClient) Query(indexName string, keyword string) (*QueryResult, error) {
	resp, err := zc.request().SetBody(map[string]interface{}{
		"search_type": "wildcard",
		"query": map[string]interface{}{
			"term": "*" + keyword + "*",
		},
		"from":        0,
		"max_results": 10,
	}).Post(fmt.Sprintf("/api/%s/_search", indexName))
	if err != nil || resp.StatusCode() != http.StatusOK {
		logrus.Warn("query failed", err)
		return nil, err
	}

	data := &QueryResult{}
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		logrus.Warn("unmarshal failed", err)
		return nil, err
	}

	return data, nil

}

func (zc *ZincClient) PutDoc(indexName string, id uint64, data interface{}) bool {
	data, _ = json.Marshal(data)
	resp, err := zc.request().SetBody(data).Put(fmt.Sprintf("/api/%s/_doc/%d", indexName, id))
	if err != nil || resp.StatusCode() != http.StatusOK {
		logrus.Warn("put doc failed", err)
		return false
	}
	return true
}

func (zc *ZincClient) DeleteDoc(indexName string, id uint64) bool {
	resp, err := zc.request().Delete(fmt.Sprintf("/api/%s/_doc/%d", indexName, id))
	if err != nil || resp.StatusCode() != http.StatusOK {
		logrus.Warn("del doc failed", err)
		return false
	}
	return true
}

func (zc *ZincClient) request() *resty.Request {
	client := resty.New()
	client.DisableWarn = true
	client.SetBaseURL(zc.ZincEndpoint)
	client.SetBasicAuth(zc.ZincUsername, zc.ZincPassword)

	return client.R()
}
