package tests

import (
	"testing"
)

func TestCountChars(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want2 int
	}{
		{"String1", args{"Распаковка файла Тестовое задание для аналитиков в SkyEng.rar завершена!\n"}, 126, 73},
		{"String1", args{"Статистика - удалено 170 плашек, снесено 36 ненужных файлов"}, 103, 59},
		{"String1", args{"Распаковка файла 7 cайтов для оттачивания навыка написания SQL запросов на 2020 год.pdf.rar завершена!"}, 173, 102},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, got2 := CountChars(tt.args.str); got != tt.want && got2 != tt.want2 {
				t.Errorf("CountChars() = %v, %v, want %v, want2 %v", got, got2, tt.want, tt.want2)
			}
		})
	}
}

// Функция, которую мы хотим протестировать
func CountChars(str string) (int, int) {
	return len(str), len([]rune(str))
}
