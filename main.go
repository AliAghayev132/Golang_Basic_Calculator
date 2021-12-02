package main

import "fmt"

func main() {

	var calculator Calculator

	for true {
		var a, b, result int
		c := "+"
		print("\n1-ci ədədi daxil edin: ")
		fmt.Scan(&a)
		print("2-ci ədədi daxil edin: ")
		fmt.Scan(&b)
		println("Prosesi seçin: \n" +
			"Toplama + \n" +
			"Çıxma - \n" +
			"Vurma *\n" +
			"Bölmə /")
		fmt.Scan(&c)

		switch c {
		case "+":
			result = calculator.toplama(a, b)
			break
		case "-":
			result = calculator.cixma(a, b)
			break
		case "*":
			result = calculator.vurma(a, b)
			break
		case "/":
			result = calculator.bolme(a, b)
			break
		}

		println(a, c, b, "=", result)

	}

}

type Calculator struct{}

func (Calculator) toplama(par1 int, par2 int) int {
	return par1 + par2
}

func (Calculator) cixma(par1 int, par2 int) int {
	return par1 - par2
}

func (Calculator) vurma(par1 int, par2 int) int {
	return par1 * par2
}

func (Calculator) bolme(par1 int, par2 int) int {
	return par1 / par2
}
