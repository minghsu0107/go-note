// this shows how to jump out the for-switch block
package main

import (
	"fmt"
)

func main() {
loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			//break    // forever printing breaking out...
			break loop
		}
	}
	fmt.Println("out...")
}
