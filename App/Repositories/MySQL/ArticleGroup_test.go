package MySQL

import (
	"fmt"
	"testing"
)

func TestArticleGroup_TableName(t *testing.T) {
	db := GetDB()
	groups := []ArticleGroup{}

	db.Find(&groups)

	for _, v := range groups {
		fmt.Println(v.Id, v.Name, v.Created_at, v.Updated_at)
	}
}
