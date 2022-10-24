package directus

import "github.com/mgjules/mgjules-go/entity"

type Technology struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (t Technology) ToEntity() entity.Technology {
	return entity.Technology{
		ID:   t.ID,
		Name: t.Name,
	}
}
