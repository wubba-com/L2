package main

import (
	"fmt"
	"strconv"
)

func RepeatS(s string) string {
	var repeatCh string
	beyond := func(c rune) bool {
		return c >= 48 && c <= 57
	}

	for i, v := range s {
		// Если строка является численной строкой
		if beyond(v) {
			// И берем предыдущий элемент на который вскоре будем умножать
			c := s[i-1]
			// проверяем, что он не является численной строкой
			if beyond(rune(c)) {
				continue
			}
			var n string
			// В цикле начинаем проверять является ли следующий символ числовой строки, да бы проверить ее в разряде десятков, тысяч итд.
			for _, v2 := range s[i:] {
				if beyond(v2) {
					// создаем будущее число из строки
					n+=string(v2)
				} else {
					// Если это буква тогда закрываем цикл
					break
				}

			}
			// переводим строку в число
			x, err := strconv.Atoi(n)
			if err != nil {
				return ""
			}
			// повторяем букву на x
			for (x - 1) > 0 {
				repeatCh += string(c)
				x--
			}
		} else {
			// если это не число, то прибавляем к будущей новой строке
			repeatCh += string(v)
		}
	}
	return repeatCh
}

func main()  {
	s := "s3"
	s = RepeatS(s)
	fmt.Println(s)
}
