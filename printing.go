package printing

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type printer struct {
	writer io.Writer
}

// NewPrinter wraps an io.Writer, exposing Print, Printf, and Println methods.
func NewPrinter(writer io.Writer) *printer {
	return &printer{writer: writer}
}

func (this *printer) Printf(format string, args ...any) (int, error) {
	return fmt.Fprintf(this.writer, format, args...)
}
func (this *printer) Println(args ...any) (int, error) {
	return fmt.Fprintln(this.writer, args...)
}
func (this *printer) Print(args ...any) (int, error) {
	return fmt.Fprint(this.writer, args...)
}

type buffer struct {
	*bytes.Buffer
	*printer
}

// NewBuffer is a near drop-in replacement for bytes.NewBuffer(), and exposes the result as a printer.
func NewBuffer(content []byte) *buffer {
	b := bytes.NewBuffer(content)
	return &buffer{Buffer: b, printer: NewPrinter(b)}
}

// NewBufferString is a near drop-in replacement for bytes.NewBufferString(), and exposes the result as a printer.
func NewBufferString(content string) *buffer {
	return NewBuffer([]byte(content))
}

type builder struct {
	*strings.Builder
	*printer
}

// NewBuilder is a near drop-in replacement for new(strings.Builder) and exposes the result as a printer.
func NewBuilder() *builder {
	b := new(strings.Builder)
	return &builder{Builder: b, printer: NewPrinter(b)}
}
