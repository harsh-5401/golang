// package main

// import "fmt"

// type payment struct{
// 	gateway stripe
// }

// // open close principle 

// func (p payment) makePayment(amount float32) {
// 	// razorpayPaymentGw:=razorpay{}
// 	// razorpayPaymentGw.pay(amount)
// 	// stripePaymentGw:=stripe{}
// 	// stripePaymentGw.pay(amount)
// 	p.gateway.pay((amount))
// }

// type razorpay struct{}

// func (r razorpay) pay(amount float32) {
// 	//logic to make payment
// 	fmt.Println("make payment using razorpay" , amount)
// }

// type stripe struct{}

// func (s stripe) pay(amount float32){
// 	fmt.Println("make payment using stripe" , amount)
// }

// func main() {
// 	// newPayment:=payment{}
// 	stripePaymentGW:=stripe{}
// 	newPayment:=payment{
// 		gateway:stripePaymentGW,
// 	}
// 	newPayment.makePayment(100)

// 	// problem : still we cant add more payment gateways
// }






package main

import "fmt"

// interfcae - bsically they are contractors 

type paymenter interface{
	pay(amount float32)  // method
	refund(amount float32 ,account string)
}

type payment struct{
	gateway paymenter
}

// open close principle 

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw:=razorpay{}
	// razorpayPaymentGw.pay(amount)
	// stripePaymentGw:=stripe{}
	// stripePaymentGw.pay(amount)
	p.gateway.pay((amount))
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("make payment using razorpay" , amount)
}

func (r razorpay) refund(amount float32 , account string) {
	//logic to make payment
	fmt.Println("make refund using razorpay" , amount , account)
}

type stripe struct{}

func (s stripe) pay(amount float32){
	fmt.Println("make payment using stripe" , amount)
}

type paypal struct{}

func (p paypal) pay(amount float32){
	fmt.Println("making payment using paypal gateway" , amount)
}

func main() {
	// newPayment:=payment{}
	// stripePaymentGW:=stripe{}
	razorpayPaymentGW:=razorpay{}
	newPayment:=payment{
		gateway:razorpayPaymentGW,
	}
	newPayment.makePayment(100)

	// problem : still we cant add more payment gateways
}

