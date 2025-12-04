package main

import "fmt"

type orderStatus string

// const (
// 	received orderStatus = iota
// 	confirmed
// 	prepared
// 	delivered
// )

const (
	received  orderStatus = "received"
	confirmed             = "confirmed"
	prepared              = "prepared"
	delivered             = "delivered"
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("changing order status to ", status)
}

func main() {
	changeOrderStatus(confirmed)
	// orderStatus()
}
