package printing_test

import (
	"bytes"
	"fmt"

	"github.com/mdw-go/printing"
)

func ExampleNewPrinter() {
	var buffer bytes.Buffer
	printer := printing.NewPrinter(&buffer)
	_, _ = printer.Print("hello,")
	_, _ = printer.Printf(" %s", "world")
	_, _ = printer.Println("!")
	fmt.Println(buffer.String())
	// Output: hello, world!
}

func ExampleNewBuffer() {
	buffer := printing.NewBuffer([]byte("(initial) "))
	_, _ = buffer.Print("hello,")
	_, _ = buffer.Printf(" %s", "world")
	_, _ = buffer.Println("!")
	fmt.Println(buffer.String())
	// Output: (initial) hello, world!
}

func ExampleNewBufferString() {
	buffer := printing.NewBufferString("(initial) ")
	_, _ = buffer.Print("hello,")
	_, _ = buffer.Printf(" %s", "world")
	_, _ = buffer.Println("!")
	fmt.Println(buffer.String())
	// Output: (initial) hello, world!
}

func ExampleNewBuilder() {
	builder := printing.NewBuilder()
	_, _ = builder.Print("hello,")
	_, _ = builder.Printf(" %s", "world")
	_, _ = builder.Println("!")
	fmt.Println(builder.String())
	// Output: hello, world!
}
