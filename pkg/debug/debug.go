package debug

import (
	"encoding/json"
	"fmt"
)

func Debug(i any) {
	b, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println()
	green := "\033[32m"
	reset := "\033[0m"
	fmt.Println(green + string(b) + reset)
}
