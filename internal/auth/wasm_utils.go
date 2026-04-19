package auth

var lookupTable = [40]int32{
	5, 8, 4, 7, 9,
	4, 6, 9, 5, 5,
	6, 5, 3, 5, 4,
	4, 9, 6, 6, 8,
	8, 6, 8, 6, 5,
	8, 4, 9, 5, 9,
	8, 5, 3, 4, 7,
	7, 4, 7, 3, 9,
}

// getHTU extracts the hundreds, tens and units digits of x
// and returns their sum
func getHTU(x int32) (h, t, u, sum int32) {
	h = (x / 100) % 10
	t = (x / 10) % 10
	u = x % 10
	sum = h + t + u
	return
}

func cdx(_ int32, x int32, _ int32, _ int32, _ int32) int32 {
	_, _, _, s := getHTU(x)
	val := lookupTable[s]
	return val + 22
}

func rdx(_ int32, x int32, _ int32, _ int32, _ int32) int32 {
	h, t, _, s := getHTU(x)
	val := lookupTable[s]
	return h + t + val + 32
}

func bdx(_ int32, x int32, _ int32, _ int32, _ int32) int32 {
	h, t, _, s := getHTU(x)
	val := lookupTable[s]
	return h + t + val + 60
}

func ndx(_ int32, x int32, _ int32, _ int32, _ int32) int32 {
	_, t, _, s := getHTU(x)
	val := lookupTable[s]
	return t + val + 88
}

func mdx(_ int32, x int32, _ int32, _ int32, _ int32) int32 {
	h, _, _, s := getHTU(x)
	val := lookupTable[s]
	return h + val + 110
}
