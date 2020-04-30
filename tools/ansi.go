package tools

import (
	"fmt"
	"github.com/mitchellh/colorstring"
)

func PrintError(err error) {
	_, _ = colorstring.Println(fmt.Sprintf("[_red_]/!\\ %s /!\\", err))
}
