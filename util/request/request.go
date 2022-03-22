package request

import (
	"encoding/json"

	"fmt"
	"io"

	"github.com/uccu/autom/http_error"
)

type ShouldBind interface {
	ShouldBind(i interface{}) error
}

func Bind(c ShouldBind, i interface{}) {

	err := c.ShouldBind(i)
	if err != nil && err != io.EOF {

		terr := http_error.MissingParametersError
		if e, ok := err.(*json.UnmarshalTypeError); ok {
			terr.ErrorMsg += ", 参数类型错误: " + e.Field
		}

		fmt.Println(err.Error())
		panic(terr)
	}

}
