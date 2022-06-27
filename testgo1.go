package testgo1

import (
	"fmt"

	"github.com/ningzero/testgo2"
)

func CallTestgo1() {
	fmt.Println("CallTestgo1")
	testgo2.CallTestgo2()
}
