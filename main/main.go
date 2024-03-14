package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(SpamMasker(inputUsers())) //все в одну строку. Думаю норм, читабельно:)
}

func inputUsers() []byte {
	read := bufio.NewReader(os.Stdin)
	inputUser, _ := read.ReadString('\n') // к сожалению, нельзя сразу записать с консоли в срез, либо я криворукий
	textSlice := []byte(inputUser)
	return textSlice
}

func SpamMasker(textSlice []byte) string {
	check := "http://"
	l1 := len(textSlice)

	var result []byte
	flagReplace := false
	//
	for i := 0; i < l1; i++ {
		if hasPrefix(textSlice[i:], check) {
			// ищем спам ссылки"
			result = append(result, []byte(check)...) // распаковочка :) (раньше кст не знал, что так можно)
			i += len(check) - 1
			flagReplace = true
		} else if flagReplace && textSlice[i] != ' ' {
			/*если сработал флаг и символ не пробел, заменяем символы на "*"
			так же добавил проверку на символы каретки и перевода строки так как из-за этого появлялся баг ( а может фича xD)
			при указании ссылки к примеру в букву после подстроки "http:// появлялось допом еще 2 звездочки, как раз таки из-за
			этих двух ребят:), но если дальше присутствует текст, после ссылки (таргет я ловлю на пробеле) то все ок, так как
			они не отображаются, так как, это по сути вообще не символы:)*/
			result = append(result, 42)
		} else { // иначе записываем все как есть
			result = append(result, textSlice[i])
		}
		if textSlice[i] == ' ' { // если встретился пробел, отключаем флаг замены
			flagReplace = false
		}
	}
	var resultt string
	resultt = string(result)
	return resultt
}
func hasPrefix(textSlice []byte, check string) bool { //Галя, неси костыли.
	// Та же реализация по сути, что и в стандартном пакете strings, только мы принимаем срез битов, а не строку
	if len(textSlice) < len(check) {
		return false
	}

	for j := 0; j < len(check); j++ { // сравнивает символы, кхмн, точнее байты, т.к. у нас строки сравниваются по байтам а не по символам в Golang
		if textSlice[j] != check[j] {
			return false
		}
	}

	return true
}
