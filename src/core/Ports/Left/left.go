package left

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

type API interface {
	ForRoot(port string) error
}

var _api API = nil

func Serve(port string) error {
	if _api == nil {
		return errors.New(color.RedString("Error: API Provider Null"))
	}
	return _api.ForRoot(port)
}

func InjectApi(api API) {
	if _api == nil {
		_api = api
		return
	}
	fmt.Println(color.YellowString("WARN: API has been already provided!"))
}
