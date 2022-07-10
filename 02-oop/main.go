package main

import "fmt"

type CheckingAccount struct {
	accountHolder string
	branchNumber  int
	accountNumber int
	balance       float64
}

func main() {
	homeroAccount := CheckingAccount{accountHolder: "Homero",
		branchNumber: 589, accountNumber: 123456, balance: 125.5}

	beatrizAccount := CheckingAccount{"Beatriz", 222, 111222, 200}

	fmt.Println(homeroAccount)
	fmt.Println(beatrizAccount)
}
