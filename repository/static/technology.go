package static

import (
	"strings"

	"github.com/mgjules/mgjules-go/entity"
	"github.com/samber/lo"
)

var technologies = []entity.Technology{
	{
		Name: "Go",
		Link: "https://go.dev",
	},
}

func getTechnologies(tt ...string) []entity.Technology {
	tt = lo.Uniq(tt)
	techs := make([]entity.Technology, 0, len(tt))
	for _, t := range tt {
		tech, found := lo.Find(technologies, func(tech entity.Technology) bool {
			return strings.EqualFold(tech.Name, t)
		})
		if !found {
			continue
		}

		techs = append(techs, tech)
	}
	return lo.Compact(techs)
}
