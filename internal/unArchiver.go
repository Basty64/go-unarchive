package internal

import (
	"fmt"
	"github.com/mholt/archiver"
	"os"
	"path/filepath"
	"strings"
)

func UnArchiver(path string) error {

	extractPath := filepath.Join(path, "unrared")

	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("проблема с корневой папкой, проверьте путь, текст ошибки - %v", err)
	}

	dir, _ := os.ReadDir(extractPath)
	if len(dir) > 1 {
		return fmt.Errorf("папка || %s || уже существует, проверьте её содержимое, она должна быть пустой", extractPath)
	}

	for _, file := range files {

		if file.Name() == ".DS_Store" {
			continue
		}

		fpath := filepath.Join(path, file.Name())

		archiver.NewZip()
		rar := archiver.NewRar()
		rar.OverwriteExisting = true

		// Распаковываем архив
		err := rar.Unarchive(fpath, extractPath)
		//на некоторых итерациях выдаётся ни на что не влияющая ошибка библиотеки, она скрыта где-то глубоко под капотом
		if err != nil {
			if strings.Contains(err.Error(), "is a directory") != true {
				_ = fmt.Errorf("%w", err)
			}
		}

		fmt.Printf("Распаковка файла: %s завершена!\n", file.Name())

	}

	message := fmt.Sprintf("Распаковка файла %s завершена!", files[len(files)-1].Name())
	countMessage := len([]rune(message))

	fmt.Println()
	MessageConstrunctor(countMessage, "Разархивация завершена", "-")

	MessageConstrunctor(countMessage, "Начинаю переименовывание", "-")

	//форматируем все папки
	err = ProcessFolder(extractPath)

	if err != nil {
		fmt.Println("Ошибка при обработке файлов:", err)
	}

	MessageConstrunctor(countMessage, "Работа выполнена, милорд!", "-")
	resultMessage := fmt.Sprintf("Статистика - удалено %d плашек, снесено %d ненужных файлов", Resultprefixes, Resultfiles)
	MessageConstrunctor(countMessage, resultMessage, "-")

	return nil
}
