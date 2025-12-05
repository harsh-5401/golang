// package main

// import (
// 	"fmt"
// 	"time"
// )

// func task(index int) {
// 	fmt.Println("task is completed" , index)
// }

// func main() {
// 	for i:=1 ; i<=10 ;i++ {
// 		go task(i)  // light weight threads
// 	}

// 	time.Sleep(time.Microsecond*10)  // use waithroups instead
// }

package main

import (
	"fmt"
	"sync"
)

func task(index int , w*sync.WaitGroup) {
	fmt.Println("task is completed" , index)
	defer w.Done()
}

func main() {

	var wg sync.WaitGroup

	for i:=1 ; i<=10 ;i++ {
		wg.Add(1)
		go task(i , &wg)  // light weight threads
	}

	wg.Wait()

	// time.Sleep(time.Microsecond*10)  // use waithroups instead 
}