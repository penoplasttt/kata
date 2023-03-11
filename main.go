package main

import (
	"bufio"
	"strconv"
	"strings"
	"fmt"
	"os"
	"unicode"
	"errors"
)

func romanToInt(s string) int {    //римские цифры в арабские
    
    rom := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
    }

    answer := rom[s[len(s)-1]]

    for i:= len(s)-2; i>=0; i-- {
        if rom[s[i]] < rom[s[i+1]]{
           answer -= rom[s[i]]
        } else {
        answer += rom[s[i]]
    }
}

    return answer
}

func intToRoman (n int) string {    //арабские цифры в римские

	if (n >= 100) {
		if n >= 400  {
		return "CD" + intToRoman(n - 400)
	} else { 
		return "C" + intToRoman(n - 100)
	}
	}

	if (n >= 50) {
		if n >= 90 {
		return "XC" + intToRoman(n - 90)
	} else { 
		return "L" + intToRoman(n - 50)
	}
	}

	if (n >= 10) {
		if n >= 40  {
		return "XL" + intToRoman(n - 40)
	} else { 
		return "X" + intToRoman(n - 10)
	}
	}

	if (n >= 5) {
		if n == 9  {
		return "IX" 
	} else { 
		return "V" + intToRoman(n - 5)
	}
	}


	if (n > 0) {
		if n == 4 {
		return "IV" 
	} else { 
		return "I" + intToRoman(n - 1)
	}
	}

	return ""
}

func operationDefinition(ex string) string {    //определение арифметической операции
	if strings.Contains(ex, "+"){
		return "+"
	} else if strings.Contains(ex, "-"){
		return "-"
	} else if strings.Contains(ex, "*"){
		return "*"
	} else if strings.Contains(ex, "/"){
		return "/"
	} else {
		return ""
	}
}

func calculation(a,b int, oper string)int{    //вычисление
	var res int
	switch oper{
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}
	return res
}

func numberSystem(arr []string) (int, int, string){   //определение системы счисления
	var flag1, flag2 string
	var a, b int

	for _,v := range arr[0]{
		
			if unicode.IsLetter(v){
				a = romanToInt(arr[0])
				flag1 = "roman"
			} else if unicode.IsNumber(v){
				a,_ = strconv.Atoi(arr[0])
				flag1 = "arabic"
			} 

	}

	for _,v := range arr[1]{
		
		if unicode.IsLetter(v){
			b = romanToInt(arr[1])
			flag2 = "roman"
		} else if unicode.IsNumber(v){
			b,_ = strconv.Atoi(arr[1])
			flag2 = "arabic"
		} 

	}
	if flag1 != flag2{
		flag1 = ""
	}
	return a, b, flag1
}

func main(){
	ex, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//ex = strings.ReplaceAll(example, " ", "")

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	arr := strings.FieldsFunc(ex, f)
	
	if len(arr) != 2{
		err := errors.New("the calculator only works with two operands")
		fmt.Print(err)
		return
	}

	oper := operationDefinition(ex)
	if oper == ""{
		err := errors.New("arithmetic operation not found")
		fmt.Print(err)
		return
	}

	a, b, flag := numberSystem(arr)
	if flag == ""{
		err := errors.New("different number systems")
		fmt.Print(err)
		return
	}

	if a > 10 || a < 1 || b > 10 || b < 1 {
		err := errors.New("the calculator accepts numbers from 1 to 10")
		fmt.Print(err)
		return
	}

	answer := calculation(a, b, oper)

	if flag == "roman"{
		if intToRoman(answer) == ""{
			err := errors.New("the result of working with Roman numerals is less than 1")
			fmt.Print(err)
		}
		fmt.Print(intToRoman(answer))
		return
	}

	fmt.Print(answer)
}