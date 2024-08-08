package main

import (
	"bufio"
	"fmt"
	"go-unarchive/internal"
	"os"
	"strings"
)

func main() {

	//              TO DO:
	//- проверка на наличие нужных расширений
	//- добавить разархивацию zip
	//- добавить выбор для переименования и отдельно удаления

	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите путь к папке, для выхода введите 'exit':")

		path, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		path = strings.TrimSpace(path)
		if path == "exit" {
			break
		}

		err = internal.UnArchiver(path)
		if err != nil {
			fmt.Printf("\nПроблема с запуском приложения -%v\n\n", err)
			continue
		} else {
			break
		}
	}

}
