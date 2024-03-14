package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(string(SpamMasker(inputUsers()))) //все в одну строку. Думаю норм, читабельно:)
}

func inputUsers() []byte {
	read := bufio.NewReader(os.Stdin)
	inputUser, _ := read.ReadString('\n') // к сожалению, нельзя сразу записать с консоли в переменную, либо я криворукий
	textSlice := []byte(inputUser)
	return textSlice
}

func SpamMasker(textSlice []byte) []byte {
	check := "http://"
	l1 := len(textSlice)

	var result []byte
	flagReplace := false
	fmt.Println(textSlice)

	for i := 0; i < l1; i++ {
		if strings.HasPrefix(string(textSlice[i:]), check) {
			// ищем спам ссылки (надеюсь на работе не будет таких задач, это больнючие костыли, я попробовал - мне не понравилось"
			// все можно было сделать проще, но по условиям нельзя:(
			result = append(result, []byte(check)...) // распаковочка :) (раньше кст не знал, как это делается)
			i += len(check) - 1
			flagReplace = true
		} else if flagReplace && textSlice[i] != ' ' && textSlice[i] != 13 && textSlice[i] != 10 {
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

	return result
}
