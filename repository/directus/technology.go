package directus

import "github.com/mgjules/mgjules-go/entity"

type Technology struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

func (t Technology) ToEntity() entity.Technology {
	return entity.Technology{
		ID:   t.ID,
		Name: t.Name,
		Link: t.Link,
	}
}
