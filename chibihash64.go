package chibihashgo

import (
	"encoding/binary"
)

func chibihash64__load64le(p *[8]uint8) uint64 {
	return binary.LittleEndian.Uint64(p[:])
}

func chibihash64(keyIn []uint8, len0 int, seed uint64) uint64 {
	k := keyIn
	kpos := 0
	l := len0

	const P1 = 0x2B7E151628AED2A5
	const P2 = 0x9E3793492EEDC3F7
	const P3 = 0x3243F6A8885A308D

	h := []uint64{P1, P2, P3, seed}

	for ; l >= 32; l -= 32 {
		for i := 0; i < 4; i++ {
			lane := chibihash64__load64le((*[8]uint8)(k[kpos : kpos+8]))
			h[i] ^= lane
			h[i] *= P1
			h[(i+1)&3] ^= ((lane << 40) | (lane >> 24))
			kpos += 8
		}
	}

	h[0] += ((uint64)(len0 << 32)) | ((uint64)(len0 >> 32))
	if (l & 1) != 0 {
		h[0] ^= uint64(k[0])
		l--
		kpos++
	}
	h[0] *= P2
	h[0] ^= h[0] >> 31

	for i := 1; l >= 8; l -= 8 {
		if kpos >= len0 {
			break
		}

		v := chibihash64__load64le((*[8]uint8)(k[kpos : kpos+8]))
		h[i] ^= v
		h[i] *= P2
		h[i] ^= h[i] >> 31
		kpos += 8
		i++
	}

	for i := 0; l > 0; l -= 2 {
		kpos += 2
		i++
		if i >= len(h) {
			break
		}
		h[i] ^= (uint64(k[0]) | ((uint64)(k[1] << 8)))
		h[i] *= P3
		h[i] ^= h[i] >> 31
	}

	x := seed
	x ^= h[0] * ((h[2] >> 32) | 1)
	x ^= h[1] * ((h[3] >> 32) | 1)
	x ^= h[2] * ((h[0] >> 32) | 1)
	x ^= h[3] * ((h[1] >> 32) | 1)

	// moremur: https://mostlymangling.blogspot.com/2019/12/stronger-better-morer-moremur-better.html
	x ^= x >> 27
	x *= (0x3C79AC492BA7B653)
	x ^= x >> 33
	x *= (0x1C69B3F74AC4AE35)
	x ^= x >> 27

	return x
}
