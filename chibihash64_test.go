package chibihashgo

import (
	"testing"
)

func TestChibihash641(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		len0  int
		seed  uint64
		want  uint64
	}{
		// the input-output table
		// https://github.com/N-R-K/ChibiHash/issues/1
		{"blank1", "", 0, 0, 0x9EA80F3B18E26CFB},
		{"blank2", "", 0, 55555, 0x2EED9399FC4AC7E5},
		{"hi", "hi", 2, 0, 0xAF98F3924F5C80D6},
		{"123", "123", 3, 0, 0x893A5CCA05B0A883},
		{"abcdefgh", "abcdefgh", 8, 0, 0x8F922660063E3E75},
		{"Helloworld", "Hello, world!", 13, 0, 0x5AF920D8C0EBFE9F},
		{"qwerty1", "qwertyuiopasdfghjklzxcvbnm123456", 32, 0, 0x2EF296DB634F6551},
		{"qwerty2", "qwertyuiopasdfghjklzxcvbnm123456789", 35, 0, 0x0F56CF3735FFA943},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u8arr := []uint8(tt.input)
			ans := Chibihash64(u8arr, tt.len0, tt.seed)
			if ans != tt.want {
				t.Errorf("got %x, want %x", ans, tt.want)
			}
		})
	}
}
