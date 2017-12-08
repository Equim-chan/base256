package base256_test

import (
	"fmt"

	"ekyu.moe/base256"
)

func Example() {
	fmt.Println(base256.EncodeToString([]byte("Hello, 世界")))
	fmt.Println(string(base256.DecodeString("👾🍧🙆🍬🙇🌱😌🚟💦🏥🐴🏤👈")))
	// Output:
	// 👾🍧🙆🍬🙇🌱😌🚟💦🏥🐴🏤👈
	// Hello, 世界
}
