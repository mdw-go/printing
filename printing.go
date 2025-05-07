package printing

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Printer is the main value-add of this package: types with print methods.
// This printer's methods, unlike the fmt methods with corresponding names,
// do not return the number of bytes written nor an error value. Indeed,
// any error values generated are discarded. You have been warned.
type Printer interface {
	Printf(format string, args ...any)
	Println(args ...any)
	Print(args ...any)
}

type printer struct{ io.Writer }

// NewPrinter wraps an io.Writer, exposing Print, Printf, and Println methods.
func NewPrinter(w io.Writer) Printer { return &printer{w} }

func (this *printer) Printf(f string, a ...any) { _, _ = fmt.Fprintf(this.Writer, f, a...) }
func (this *printer) Println(a ...any)          { _, _ = fmt.Fprintln(this.Writer, a...) }
func (this *printer) Print(a ...any)            { _, _ = fmt.Fprint(this.Writer, a...) }

type Buffer struct {
	inner *bytes.Buffer
	Printer
}

// NewBuffer is a near drop-in replacement for bytes.NewBuffer(), and exposes the result as a printer.
func NewBuffer(c []byte) *Buffer { b := bytes.NewBuffer(c); return &Buffer{b, NewPrinter(b)} }

// NewBufferString is a near drop-in replacement for bytes.NewBufferString(), and exposes the result as a printer.
func NewBufferString(c string) *Buffer { return NewBuffer([]byte(c)) }

// Inner provides read-only access to the underlying *bytes.Buffer.
func (this *Buffer) Inner() *bytes.Buffer { return this.inner }

type Builder struct {
	inner *strings.Builder
	Printer
}

// NewBuilder is a near drop-in replacement for new(strings.Builder) and exposes the result as a printer.
func NewBuilder() *Builder { b := new(strings.Builder); return &Builder{b, NewPrinter(b)} }

// Inner provides read-only access to the underlying *strings.Builder.
func (this *Builder) Inner() *strings.Builder { return this.inner }
