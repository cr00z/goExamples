package main

import "fmt"

func Int32ToIp(n uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", n>>24, (n&0xff0000)>>16, (n&0xff00)>>8, n&0xff)
}

func main() {
	fmt.Println(Int32ToIp(2149583361))
}
