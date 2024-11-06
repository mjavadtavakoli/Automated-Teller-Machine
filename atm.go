package main

import "fmt"

func calcu() {
	var balance float64 = 1000
	for {
		fmt.Printf("wellcome to atm! %v\n", balance)
		println("option 1 withraw")
		println("option 2 send")
		println("optioon 3 show amount")
		println("option 4 exit")

		var userss float64
		fmt.Println("gozine mord nazar ra vard konid :")
		fmt.Scanln(&userss)

		if userss == 1 {
			var bardashtt float64
			fmt.Println("mablagh bardasht ra vard konid \n your balance", balance)
			fmt.Scanln(&bardashtt)
			if bardashtt <= balance {
				balance -= bardashtt
				fmt.Println("mojodi bad az bardasht:", balance)

			} else {
				fmt.Println("invalid aamount!")

			}

		} else if userss == 2 {
			var enteghal float64
			fmt.Println("mablagh enteghal ra vard konid")
			fmt.Scanln(&enteghal)
			if enteghal <= balance {
				balance -= enteghal
				fmt.Println("mojodi bad az enteghal : ", balance)
			}

		} else if userss == 3 {
			fmt.Println("your amount : ", balance)

		} else {
			return
		}

	}

}
