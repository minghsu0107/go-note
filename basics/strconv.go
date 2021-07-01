package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "0.12345678901234567890"
	f, err := strconv.ParseFloat(s, 32)
	fmt.Println(f, err)          // 0.12345679104328156 <nil>
	fmt.Println(float32(f), err) // 0.12345679 <nil>
	f, err = strconv.ParseFloat(s, 64)
	fmt.Println(f, err) // 0.12345678901234568 <nil>

	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	// bitSize: 0:int, 8:int8, 16:int16, 32:int32, 64:int64
	fmt.Println(strconv.ParseInt("123", 10, 8))
	fmt.Println(strconv.Atoi("123")) // equivalent to ParseInt(s, 10, 0)

	fmt.Println(strconv.ParseUint("FF", 16, 8)) // for uint

	i := int64(-2048)
	fmt.Println(strconv.FormatInt(i, 2))  // -100000000000
	fmt.Println(strconv.FormatInt(i, 8))  // -4000
	fmt.Println(strconv.FormatInt(i, 10)) // -2048
	fmt.Println(strconv.FormatInt(i, 16)) // -800
	fmt.Println(strconv.FormatInt(i, 36)) // -1kw

	fmt.Println(strconv.Itoa(-2048)) // -2048 (string)
}
