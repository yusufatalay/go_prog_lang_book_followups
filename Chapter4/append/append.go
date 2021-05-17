package main

func appendInt(x []int, y int) []int {

	var z []int
	zlen := len(x) + y

	if zlen <= cap(x) {
		// capacity of x will not be exceeded after the append so slice can be extended
		z = x[:zlen]
	} else {
		// insufficent space, grow the array by doubling its size
		zcap := zlen

		if zcap < 2*len(x) {

			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	z[len(x)] = y
	return z

}

func main() {
}
