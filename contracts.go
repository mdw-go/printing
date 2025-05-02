package printing

import "io"

// Printer is the main value-add of this package: types with print methods.
type Printer interface {
	Printf(format string, args ...any) (int, error)
	Println(args ...any) (int, error)
	Print(args ...any) (int, error)
}

// Buffer is the interface implemented by bytes.Buffer, along with Printer
type Buffer interface {
	Printer
	Bytes() []byte
	AvailableBuffer() []byte
	String() string
	Len() int
	Cap() int
	Available() int
	Truncate(n int)
	Reset()
	Grow(n int)
	Write(p []byte) (n int, err error)
	WriteString(s string) (n int, err error)
	ReadFrom(r io.Reader) (n int64, err error)
	WriteTo(w io.Writer) (n int64, err error)
	WriteByte(c byte) error
	WriteRune(r rune) (n int, err error)
	Read(p []byte) (n int, err error)
	Next(n int) []byte
	ReadByte() (byte, error)
	ReadRune() (r rune, size int, err error)
	UnreadRune() error
	UnreadByte() error
	ReadBytes(delim byte) (line []byte, err error)
	ReadString(delim byte) (line string, err error)
}

// Builder is the interface implemented by strings.Builder, along with Printer
type Builder interface {
	Printer
	String() string
	Len() int
	Cap() int
	Reset()
	Grow(n int)
	Write(p []byte) (int, error)
	WriteByte(c byte) error
	WriteRune(r rune) (int, error)
	WriteString(s string) (int, error)
}
