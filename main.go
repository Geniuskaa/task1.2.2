package main

func main() {
	var balance int32 = 0 // 15 миллионов в копейках
	var transfer int32 = 0 // 10 миллионов в копейках
	total := balance + transfer // int32 + int32 будет int32
	println(total)
}
