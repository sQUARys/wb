package task

import (
	"fmt"
	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func GetTime() string {
	time, err := ntp.Time("time.nist.gov")
	if err != nil {
		fmt.Println("Error:", err)
		return "error"
	}
	const layout = "3:04:05 PM (MST) on Monday, January _2, 2006"
	result := "Current Local Time:" + time.Local().Format(layout)
	return result
}
