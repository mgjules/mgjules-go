package projecter

import (
	"errors"
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/mapstruct"
	"github.com/samber/lo"
)

func (p *Projecter) BuildBlogIndex() ([]byte, error) {
	link, found := lo.Find(p.fetcher.Links(), func(link entity.Link) bool {
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
	for _, post := range p.fetcher.Posts() {
		tabs = append(tabs, entity.Tab{
			Name:      post.Title,
			Icon:      "ooui:article-ltr",
			URL:       "/blog/" + post.Slug,
			Extension: "post",
		})
	}

	values := map[string]any{
		"title":       p.fetcher.Meta().FullName + " - " + currentTab.Name + "." + currentTab.Extension,
		"tabs":        mapstruct.FromSlice(tabs),
		"current_tab": mapstruct.FromSingle(currentTab),
		"posts":       mapstruct.FromSlice(p.fetcher.Posts()),
	}

	out, err := p.render(values, "Blog", "templates/blog/index.dhtml")
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}

func (p *Projecter) BuildBlogPost(post entity.Post) ([]byte, error) {
	link, found := lo.Find(p.fetcher.Links(), func(link entity.Link) bool {
		return link.Name == "Blog"
	})
	if !found {
		return nil, errors.New("missing current link")
	}

	currentTab := entity.Tab{
		Name:      post.Title,
		Icon:      "ooui:article-ltr",
		Extension: "post",
	}

	tabs := []entity.Tab{
		{
			Name:      link.Name,
			Icon:      link.Icon,
			Extension: "index",
			URL:       link.URL,
		},
	}
	for _, post := range p.fetcher.Posts() {
		tabs = append(tabs, entity.Tab{
			Name:      post.Title,
			Icon:      "ooui:article-ltr",
			URL:       "/blog/" + post.Slug,
			Extension: "post",
		})
	}

	values := map[string]any{
		"title":       p.fetcher.Meta().FullName + " - " + currentTab.Name + "." + currentTab.Extension,
		"tabs":        mapstruct.FromSlice(tabs),
		"current_tab": mapstruct.FromSingle(currentTab),
		"post":        mapstruct.FromSingle(post),
	}

	out, err := p.render(values, "Blog", "templates/blog/post.dhtml")
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}
