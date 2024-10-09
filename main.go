package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("___ Калькулятор индекса массы тела ___")
	for {
		userHeight, userKg, err := getUserInput()
		if err != nil {
			fmt.Println("Введено неправильные данные!")
			continue
		}

		IMT := calculateIMT(userHeight, userKg)
		outputResult(IMT)
		isRepeatCalculation := checkRepearCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f\n", imt)
	fmt.Print(result)

	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateIMT(userHeight float64, userKg float64) float64 {
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64, error) {
	var userHeight float64
	var userKg float64

	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)

	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&userKg)

	if userHeight <= 0 || userKg <= 0 {
		return 0, 0, errors.New("NO_PARAMS_ERROR")
	}

	return userHeight, userKg, nil
}

func checkRepearCalculation() bool {
	var userChoise string
	fmt.Print("Хотите повторить? y/n: ")
	fmt.Scan(&userChoise)

	if userChoise == "Y" || userChoise == "y" {
		return true
	}

	return false
}
