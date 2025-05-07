package printing_test

import (
	"bytes"
	"fmt"

	"github.com/mdw-go/printing"
)

func ExampleNewPrinter() {
	var buffer bytes.Buffer
	printer := printing.NewPrinter(&buffer)
	printer.Print("hello,")
	printer.Printf(" %s", "world")
	printer.Println("!")
	fmt.Println(buffer.String())
	// Output: hello, world!
}

func ExampleNewBuffer() {
	buffer := printing.NewBuffer([]byte("(initial) "))
	buffer.Print("hello,")
	buffer.Printf(" %s", "world")
	buffer.Println("!")
	fmt.Println(buffer.Inner().String())
	// Output: (initial) hello, world!
}

func ExampleNewBufferString() {
	buffer := printing.NewBufferString("(initial) ")
	buffer.Print("hello,")
	buffer.Printf(" %s", "world")
	buffer.Println("!")
	fmt.Println(buffer.Inner().String())
	// Output: (initial) hello, world!
}

func ExampleNewBuilder() {
	builder := printing.NewBuilder()
	builder.Print("hello,")
	builder.Printf(" %s", "world")
	builder.Println("!")
	fmt.Println(builder.Inner().String())
	// Output: hello, world!
}
