package main

func main() {
	// Boolean
	// Default is false
	var isTrue bool = true
	println(isTrue)

	// Integer (int8, int16, int32, int64)
	// If not specified, it will be int by default and based on the system
	// On 64-bit system, int is 64-bit and uint is 64-bit
	var i int = 10
	println(i)

	// Unsigned Integer (uint8, uint16, uint32, uint64)
	var ui int = 10000
	println(ui)

	// On Go we can use alias for the data types
	// byte is alias for uint8
	var b byte = 255
	println(b)

	// rune is alias for int32
	var r rune = 1234
	println(r)

	// Floating Point
	// Float requires specific number of bits
	// float32 and float64
	var f float32 = 3.14
	println(f)

	// Complex
	var c complex64 = 5 + 5i
	println(c)

	// String
	var s string = "Hello World"
	println(s)

	// Error
	// Default value is nil <nil>
	var e error
	println(e)

	// Array
	var arr [3]int = [3]int{1, 2, 3}
	println(arr)

	// Slice
	var slice []int = []int{1, 2, 3}
	println(slice)

	// Map
	var m map[string]int = map[string]int{"one": 1, "two": 2}
	println(m)

	// Pointer
	var p *int = &i
	println(p)

	// Struct
	type person struct {
		name string
		age  int
	}
	var p1 person = person{"John", 20}
	println(p1)

	// Channel
	var ch chan int = make(chan int)
	println(ch)

	// Interface
	var inter interface{}
	inter = 1
	println(inter)
}
