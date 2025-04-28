package databasestorageapi

import "github.com/elastic/go-elasticsearch/v8/esapi"

func responseClose(res *esapi.Response) {
	if res == nil || res.Body == nil {
		return
	}

	res.Body.Close()
}
