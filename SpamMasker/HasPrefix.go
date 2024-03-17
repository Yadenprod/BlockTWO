package SpamMasker

func HasPrefix(textSlice []byte, check string) bool { //Галя, неси костыли.
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
