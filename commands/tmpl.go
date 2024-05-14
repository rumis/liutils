package commands

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

type tmplData struct {
	*Command
	*CommandEnvHolder
}

func makeTmplData(cmd *Command) tmplData {
	// Minimum width of the command column
	width := 12
	for _, c := range cmd.Commands {
		l := len(c.Name())
		if width < l {
			width = l
		}
	}
	CommandEnv.CommandsWidth = width
	return tmplData{
		Command:          cmd,
		CommandEnvHolder: &CommandEnv,
	}
}

// An errWriter wraps a writer, recording whether a write error occurred.
type errWriter struct {
	w   io.Writer
	err error
}

func (w *errWriter) Write(b []byte) (int, error) {
	n, err := w.w.Write(b)
	if err != nil {
		w.err = err
	}
	return n, err
}

// buildCommandText build command text as template
func buildCommandText(cmd *Command) {
	data := makeTmplData(cmd)
	cmd.UsageLine = buildText(cmd.UsageLine, data)
	cmd.Long = buildText(cmd.Long, data)
	cmd.Short = buildText(cmd.Short, data)
	// build sub command usage
	for _, c := range cmd.Commands {
		buildCommandText(c)
	}
}

func buildText(text string, data interface{}) string {
	buf := bytes.NewBuffer([]byte{})
	tmpl(buf, text, data)
	return buf.String()
}

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace, "capitalize": capitalize, "width": width})
	template.Must(t.Parse(text))
	ew := &errWriter{w: w}
	err := t.Execute(ew, data)
	if ew.err != nil {
		// I/O error writing. Ignore write on closed pipe.
		if strings.Contains(ew.err.Error(), "pipe") {
			SetExitStatus(1)
			Exit()
		}
		Fatalf("writing output: %v", ew.err)
	}
	if err != nil {
		panic(err)
	}
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToTitle(r)) + s[n:]
}

func width(width int, value string) string {
	format := fmt.Sprintf("%%-%ds", width)
	return fmt.Sprintf(format, value)
}
