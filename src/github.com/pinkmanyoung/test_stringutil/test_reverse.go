package main

import (
	"fmt"
	"github.com/pinkmanyoung/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG,olleH"))
}

/**
go mod init github.com/pinkmanyoung/test_stringutil
go mod edit -replace github.com/pinkmanyoung/stringutil = ../stringutil
go mod tidy

*/


