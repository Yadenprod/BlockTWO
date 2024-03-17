package SpamMasker

import (
	"bufio"
	"os"
)

func InputUsers() string {
	read := bufio.NewReader(os.Stdin)
	inputUser, _ := read.ReadString('\n') // к сожалению, нельзя сразу записать с консоли в срез, либо я криворукий
	return inputUser
}
