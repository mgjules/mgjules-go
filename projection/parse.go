package projection

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/bep/godartsass"
	"github.com/mgjules/mgjules-go/logger"
)

func (p *Projection) parseSCSS(file string) (string, error) {
	var (
		scss io.Reader
		err  error
	)
	if p.prod {
		scss, err = p.templates.Open(file)
	} else {
		scss, err = os.Open(file)
	}
	if os.IsNotExist(err) {
		logger.L.Debugf("file '%s' not found", file)
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("failed to open: %w", err)
	}

	source, err := ioutil.ReadAll(scss)
	if err != nil {
		return "", fmt.Errorf("failed to read source fromr reader: %w", err)
	}

	res, err := p.transpiler.Execute(godartsass.Args{
		Source: string(source),
	})
	if err != nil {
		return "", fmt.Errorf("failed to compile: %w", err)
	}

	return res.CSS, nil
}
