package projecter

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"time"

	"github.com/flosch/pongo2/v6"
	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/mapstruct"
	"github.com/peterbourgon/mergemap"
	"github.com/samber/lo"
)

func (p *Projecter) render(meta *entity.Meta, links []entity.Link, routeName, tplFilename string, values map[string]any) ([]byte, error) {
	tpl, err := p.templateSet.FromFile(tplFilename)
	if err != nil {
		return nil, fmt.Errorf("failed to load templates from file: %w", err)
	}

	ll := lo.Map(links, func(link entity.Link, _ int) entity.Link {
		if link.Name == routeName {
			link.IsCurrent = true
		}

		return link
	})

	editorCSS, err := p.parseSCSS("templates/layouts/editor.scss")
	if err != nil {
		return nil, fmt.Errorf("failed to parse editor.scss: %w", err)
	}

	values = mergemap.Merge(values, map[string]any{
		"meta":         mapstruct.FromSingle(meta),
		"links":        mapstruct.FromSlice(ll),
		"fetched_at":   p.fetcher.FetchedAt(),
		"projected_at": p.projectedAt,
		"editor_css":   editorCSS,
		"current_year": time.Now().Year(),
	})

	out, err := tpl.ExecuteBytes(pongo2.Context(values))
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	b := &bytes.Buffer{}
	w, err := gzip.NewWriterLevel(b, gzip.BestCompression)
	if err != nil {
		return nil, fmt.Errorf("failed to create new gzip writer: %w", err)
	}

	w.Write(out)
	w.Close()

	return b.Bytes(), nil
}
