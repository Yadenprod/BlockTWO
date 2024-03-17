package SpamMasker

import (
	"unicode/utf8"
)

func SpamMasker(inputUser string) string {
	textSlice := []byte(inputUser)
	check := "http://"
	l1 := len(textSlice)

	var result []byte
	flagReplace := false

	for i := 0; i < l1; i++ {
		if HasPrefix(textSlice[i:], check) {
			// ищем спам ссылки"
			result = append(result, []byte(check)...) // распаковочка :) (раньше кст не знал, что так можно)
			i += len(check) - 1
			flagReplace = true
		} else if flagReplace && textSlice[i] != ' ' && textSlice[i] != 13 && textSlice[i] != 10 {
			/*если сработал флаг и символ не пробел, заменяем символы на "*"
			так же добавил проверку на символы каретки и перевода строки так как из-за этого появлялся баг ( а может фича xD)
			при указании ссылки к примеру в букву после подстроки "http:// появлялось допом еще 2 звездочки, как раз таки из-за
			этих двух ребят:), но если дальше присутствует текст, после ссылки (таргет я ловлю на пробеле) то все ок, так как
			они не отображаются, так как, это по сути вообще не символы:)*/
			if utf8.RuneStart(textSlice[i]) && utf8.RuneLen(rune(textSlice[i])) > 1 {
				result = append(result, 42)
				i += utf8.RuneLen(rune(textSlice[i])) - 1
				//Если символ киррилицы, удаляем один байт, чтобы была одна"*" на один символ киррилицы в котором 2 байта
			} else {
				result = append(result, 42)
			}
		} else { // иначе записываем все как есть
			result = append(result, textSlice[i])
		}
		if textSlice[i] == ' ' { // если встретился пробел, отключаем флаг замены
			flagReplace = false
		}
	}
	return string(result)
}
