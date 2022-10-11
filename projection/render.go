package projection

import (
	"fmt"
	"time"

	"github.com/flosch/pongo2/v6"
	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/mapstruct"
	"github.com/peterbourgon/mergemap"
	"github.com/samber/lo"
)

func (p *Projection) render(ctx map[string]any, routeName, tplFilename string) ([]byte, error) {
	tpl, err := p.templateSet.FromFile(tplFilename)
	if err != nil {
		return nil, fmt.Errorf("failed to load templates from file: %w", err)
	}

	links := lo.Map(p.links, func(link entity.Link, _ int) entity.Link {
		if link.Name == routeName {
			link.IsCurrent = true
		}

		return link
	})

	editorCSS, err := p.parseSCSS("templates/layouts/editor.scss")
	if err != nil {
		return nil, fmt.Errorf("failed to parse editor.scss: %w", err)
	}

	pCtx := mergemap.Merge(ctx, map[string]any{
		"meta":         mapstruct.FromSingle(p.meta),
		"links":        mapstruct.FromSlice(links),
		"editor_css":   editorCSS,
		"current_year": time.Now().Year(),
	})

	out, err := tpl.ExecuteBytes(pongo2.Context(pCtx))
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}
