package temp

import (
	"fmt"
)

func PrintMsg(msg string) error {
	if msg == "" {
		fmt.Printf("msg %v", "msg is empty")
	} else {
		fmt.Printf("msg %v", msg)
	}
	return nil
}