package chibihashgo

import (
	"fmt"
)

func main() {
	plain := []uint8{0x5d, 0xf7, 0xa0, 0x73,
		0xef, 0xb8, 0x2b, 0xb8,
		0x0c, 0x9d, 0x68, 0x6e,
		0x4e, 0x0f, 0x7f, 0xd8}
	res := Chibihash64(plain, 16, 0)
	fmt.Printf("Hash: %X\n", res)
}
