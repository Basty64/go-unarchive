package internal

import "fmt"

func MessageConstrunctor(countMessage int, endingMessage string, char string) {

	charsCount := (countMessage - len([]rune(endingMessage))) / 2

	//Получаем количество будущих символов с левой стороны
	var resultChars string
	//Такая длина тк сообщение обрамляется с двух сторон, вычитаем из-за знака "|"
	for i := 0; i < charsCount-1; i++ {
		resultChars += char
	}

	var resultMessage string

	if len([]rune(endingMessage))%2 == 0 {
		resultMessage = fmt.Sprint("|" + resultChars + endingMessage + resultChars + "|")
	} else {
		resultMessage = fmt.Sprint("|" + resultChars + endingMessage + resultChars + char + "|")
	}

	if countMessage+charsCount*2+10 <= len([]rune(endingMessage)) {
		fmt.Println(endingMessage)
	} else {
		fmt.Println(resultMessage)
	}
}
