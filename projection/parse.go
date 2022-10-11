package projection

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/wellington/go-libsass"
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
	if err != nil {
		return "", fmt.Errorf("failed to open: %w", err)
	}

	buf := bytes.NewBuffer(nil)
	comp, err := libsass.New(buf, scss)
	if err != nil {
		return "", fmt.Errorf("failed to create libsass compiler: %w", err)
	}

	if err := comp.Run(); err != nil {
		return "", fmt.Errorf("failed to compile: %w", err)
	}

	return buf.String(), nil
}
