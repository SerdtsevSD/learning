package main

import "fmt"

func main() {
	var pwd string
	fmt.Scan(&pwd)
	runepwd := []rune(pwd)
	if len(runepwd) >= 5 {
		for i := 0; i < len(runepwd); i++ {
			if ((runepwd[i] >= '0' && runepwd[i] <= '9') || (runepwd[i] >= 'A' && runepwd[i] <= 'z')) == false {
				fmt.Println("Wrong password")
				break
			} else if i == len(runepwd)-1 {
				fmt.Println("Ok")
			}
		}
	} else {
		fmt.Println("Wrong password")
	}
}