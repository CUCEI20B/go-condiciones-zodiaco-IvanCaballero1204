package main

import "fmt"

func main(){
	var day int
	var month int

	fmt.Scan(&day)
	fmt.Scan(&month)

	switch {
		case (day >= 21 && month == 3) || (day <= 20 && month == 4):
			fmt.Println("aries")
		case (day >= 21 && month == 4) || (day <= 20 && month == 5):
			fmt.Println("tauro")
		case (day >= 21 && month == 5) || (day <= 20 && month == 6):
			fmt.Println("geminis")
		case (day >= 21 && month == 6) || (day <= 20 && month == 7):
			fmt.Println("cancer")
		case (day >= 21 && month == 7) || (day <= 20 && month == 8):
			fmt.Println("leo")
		case (day >= 21 && month == 8) || (day <= 20 && month == 9):
			fmt.Println("virgo")
		case (day >= 21 && month == 9) || (day <= 20 && month == 10):
			fmt.Println("libra")
		case (day >= 21 && month == 10) || (day <= 20 && month == 11):
			fmt.Println("escorpio")
		case (day >= 21 && month == 11) || (day <= 20 && month == 12):
			fmt.Println("sagitario")
		case (day >= 21 && month == 12) || (day <= 20 && month == 1):
			fmt.Println("capricornio")
		case (day >= 21 && month == 1) || (day <= 20 && month == 2):
			fmt.Println("acuario")
		case (day >= 21 && month == 2) || (day <= 20 && month == 3):
			fmt.Println("picis")
	}
}