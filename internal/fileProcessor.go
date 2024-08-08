package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var Resultprefixes, Resultfiles int

func ProcessFolder(path string) (error error) {

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {

		workdirpath := filepath.Join(path, file.Name())

		//переименовываем папки и файлы
		if strings.HasPrefix(file.Name(), "[sharewood.biz] ") {

			newFilename := strings.TrimPrefix(file.Name(), "[sharewood.biz] ")
			newPath := filepath.Join(path, newFilename)
			err = os.Rename(workdirpath, newPath)
			if err != nil {
				fmt.Println(err)
			}
			workdirpath = newPath
			Resultprefixes++
		}

		if file.IsDir() {

			//Вызываем рекурсию
			err = ProcessFolder(workdirpath)

			if err != nil {
				fmt.Println(err)
			}
		} else {

			//сносим ненужные файлы
			switch file.Name() {
			case "SHAREWOOD_BIZ_Качай_популярные_курсы.url":
				err = os.Remove(workdirpath)
				if err != nil {
					fmt.Println("не могу удалить эту ссылку, к сожалению:", err)
				}
				Resultfiles++
			case "[sharewood.biz] Информация.docx":
				err = os.Remove(workdirpath)
				if err != nil {
					fmt.Println("не могу удалить этот файл с ненужной инфой, к сожалению:", err)
				}
				Resultfiles++
			default:
				continue
			}

		}

	}
	return nil
}
