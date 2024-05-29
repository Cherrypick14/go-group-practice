package main

func ReverseBits(oct byte) byte {
	var reverseNum byte = 0
	for i := 0; i < 8; i++ {
		reverseNum <<= 1
		if oct&1 == 1 {
			reverseNum ^= 1
		}
		oct >>= 1
	}
	return reverseNum
}
