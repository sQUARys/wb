package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ==

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	doneCh := make(chan interface{}) // создаем канал который передает интерфейсы

	switch len(channels) { // в зависимости от количества каналов делаем:
	case 0: // если каналов нет
		return nil
	case 1: // если единственный канал
		return channels[0]
	default: // если больше чем 1 канал
		go func(ch chan interface{}) { // запускаем горутину
			defer close(doneCh) // в конце закроем канал
			select {
			case <-channels[0]: //считываем из первого канала
			case <-or(append(channels[1:])...): // рекурсивно запускаем все кроме первого каналы
			}
		}(doneCh) // передаем done канал
	}
	return doneCh
}

func main() {

	sig := func(after time.Duration) <-chan interface{} { // запускаем функцию по засеканию времени
		c := make(chan interface{})
		go func() { //запускаем горутину для небольшой задержки на несколько секунд
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now() // засекаем начальное время
	<-or(
		sig(2*time.Hour), // передаем каналы с временем в часах, минутах, секундах
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("Done after %v", time.Since(start))

}
