package main

import "github.com/fluent/fluent-bit-go/output"
import (
	"github.com/ugorji/go/codec"
	"fmt"
	"unsafe"
	"C"
)

//export FLBPluginInit
func FLBPluginInit(ctx unsafe.Pointer) int {
	return output.FLBPluginRegister(ctx, "gstdout", "Stdout GO!")
}

//export FLBPluginFlush
func FLBPluginFlush(data unsafe.Pointer, length C.int, tag *C.char) int {
	var count int
	var h codec.Handle = new(codec.MsgpackHandle)
	var b []byte
	var m interface{}
	var err error

	b = C.GoBytes(data, length)
	dec := codec.NewDecoderBytes(b, h)

	count = 0
	for {
		err = dec.Decode(&m)
		if err != nil {
			break
		}
		fmt.Printf("[%d] %s: %v\n", count, C.GoString(tag), m)
		count++
	}

	return 0
}

//export FLBPluginExit
func FLBPluginExit() int {
	return 0
}

func main() {
}
