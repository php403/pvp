package main

import "fmt"

func main()  {
	ch := make(chan bool,1)
	ch <- true
	close(ch)
	fmt.Println(111)
	for {
		select {
		case a:=  <- ch :
			fmt.Println(a)
			fmt.Println("exit")
			return
		}
	}
}
