package projecter

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/bep/godartsass"
	"github.com/mgjules/mgjules-go/logger"
)

func (p *Projecter) parseSCSS(file string) (string, error) {
	var (
		scss io.ReadCloser
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
	defer scss.Close()

	buf := &strings.Builder{}
	if _, err := io.Copy(buf, scss); err != nil {
		return "", fmt.Errorf("failed to read source from reader: %w", err)
	}

	res, err := p.transpiler.Execute(godartsass.Args{
		Source: buf.String(),
	})
	if err != nil {
		return "", fmt.Errorf("failed to compile: %w", err)
	}

	return res.CSS, nil
}
