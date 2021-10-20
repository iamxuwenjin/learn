package golang

import (
	"fmt"
	"log"
)

type MyErr struct {
	Msg string
}

func (m MyErr) Error() string {
	panic("implement me")
}

func GetErr() *MyErr {
	return nil
}

func main1() {
	var e *MyErr
	e = GetErr()
	log.Println(e == nil)
}

func main2() {
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil)
	fmt.Println(in, in == nil)

	in = data
	fmt.Println(in, in == nil)
}
