package projection

import (
	"regexp"

	"github.com/araddon/dateparse"
	"github.com/extemporalgenome/slug"
	"github.com/flosch/pongo2/v6"
	"github.com/russross/blackfriday/v2"
)

func init() {
	pongo2.RegisterFilter("slugify", filterSlugify)
	pongo2.RegisterFilter("markdown", filterMarkdown)
	pongo2.RegisterFilter("removescheme", filterRemoveScheme)
	pongo2.RegisterFilter("formatdate", filterFormatDate)
}

func filterMarkdown(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsSafeValue(string(blackfriday.Run([]byte(in.String())))), nil
}

func filterSlugify(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsValue(slug.Slug(in.String())), nil
}

var filterRemoveSchemeRe = regexp.MustCompile(`^https://`)

func filterRemoveScheme(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsSafeValue(filterRemoveSchemeRe.ReplaceAllString(in.String(), "")), nil
}

func filterFormatDate(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	val := in.String()
	if val == "" {
		return pongo2.AsSafeValue(""), nil
	}

	t, err := dateparse.ParseAny(val)
	if err != nil {
		return nil, &pongo2.Error{
			Sender:    "filter:formatdate",
			OrigError: err,
		}
	}

	format := param.String()
	if format == "" {
		format = "Jan 2006"
	}

	return pongo2.AsSafeValue(t.Format(format)), nil
}
