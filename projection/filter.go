package projection

import (
	"github.com/flosch/pongo2/v6"
	"github.com/russross/blackfriday/v2"
)

func init() {
	pongo2.RegisterFilter("markdown", filterMarkdown)
}

func filterMarkdown(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsSafeValue(string(blackfriday.Run([]byte(in.String())))), nil
}
