package tools

import (
	"fmt"
	"github.com/mitchellh/colorstring"
)

func PrintError(err error) {
	_, _ = colorstring.Println(fmt.Sprintf("[red]/!\\ %s. /!\\", err))
}
