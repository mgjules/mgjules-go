package projection

import (
	"errors"
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/mapstruct"
	"github.com/samber/lo"
)

func (p *Projection) BuildBlogIndex() ([]byte, error) {
	link, found := lo.Find(p.links, func(link entity.Link) bool {
		return link.Name == "Blog"
	})
	if !found {
		return nil, errors.New("missing current link")
	}

	currentTab := entity.Tab{
		Name:      link.Name,
		Icon:      link.Icon,
		Extension: "index",
	}

	tabs := []entity.Tab{
		currentTab,
	}
	for _, post := range p.posts {
		tabs = append(tabs, entity.Tab{
			Name:      post.Title,
			Icon:      "ooui:article-ltr",
			URL:       "/blog/" + post.Slug,
			Extension: "post",
		})
	}

	pCtx := map[string]any{
		"title":       p.meta.FullName + " - " + currentTab.Name + "." + currentTab.Extension,
		"tabs":        mapstruct.FromSlice(tabs),
		"current_tab": mapstruct.FromSingle(currentTab),
		"posts":       mapstruct.FromSlice(p.posts),
	}

	out, err := p.render(pCtx, "Blog", "templates/blog/index.dhtml")
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}
