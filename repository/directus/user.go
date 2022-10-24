package directus

import (
	"github.com/avelino/slugify"
	"github.com/mgjules/mgjules-go/entity"
)

type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	Github      string `json:"github"`
	Username    string `json:"username"`
	Gender      string `json:"gender"`
}

func (u User) ToEntity(directusURL string) entity.User {
	return entity.User{
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Description: u.Description,
		Avatar:      directusURL + "/assets/" + u.Avatar + "/" + slugify.Slugify(u.FirstName+" "+u.LastName) + ".webp",
		Github:      u.Github,
		Username:    u.Username,
		Gender:      u.Gender,
	}
}
