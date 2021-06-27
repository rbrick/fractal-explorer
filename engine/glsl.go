package engine

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	pattern = regexp.MustCompile(`<([\w\/]+.glsl)>`)
)

//Preprocessor is a simple low-level directive pre-processor for GLSL
//primarily used to implement an include directive.
//probably not very secure, but functional
type Preprocessor struct {
	Directives map[string]ProcessFunc
}

func (p *Preprocessor) Process(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	var s string

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "#") {
			splitArray := strings.Split(text, " ")
			directive := splitArray[0][1:] // remove the '#'

			if f, ok := p.Directives[directive]; ok {
				if len(splitArray) > 1 {
					text = f(p, directive, splitArray[1:])
				} else {
					text = f(p, directive, []string{})
				}
			}
		}

		s += text + "\n"
	}

	return s
}

type ProcessFunc func(p *Preprocessor, directive string, tokens []string) string

func Include(p *Preprocessor, directive string, tokens []string) string {
	if len(tokens) >= 1 {
		include := tokens[0] // the include path
		if pattern.MatchString(include) {
			groups := pattern.FindAllStringSubmatch(include, -1)
			includePath := groups[0][1]

			path := filepath.Join("shaders", includePath)

			_, err := os.Stat(path)

			if os.IsNotExist(err) {
				return ""
			}

			f, err := os.Open(path)

			if err != nil {
				return ""
			}

			return p.Process(f)
		}

	}
	return ""
}
