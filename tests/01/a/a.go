package a

import "github.com/eaglexiang/costwhere-go/tests/01/z"

func A(sig chan struct{}) {
	z.Z(sig)
}
