package bytecode

import "fmt"

func RunBytecode(bytes []byte) {
	for i := 0; i < len(bytes); i++ {
		command := bytes[i]

		switch command {
		case 0:
			fmt.Print(">")
		case 1:
			fmt.Print("<")
		case 2:
			fmt.Print("+")
		case 3:
			fmt.Print("-")
		case 4:
			fmt.Print(".")
		case 5:
			fmt.Print(",")
		case 6:
			fmt.Print("[")
		case 7:
			fmt.Print("]")
		}
	}
}
