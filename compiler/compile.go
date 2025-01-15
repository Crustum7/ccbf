package compiler

import "os"

func dump(bytes []byte, outFileName string) {
	os.WriteFile(outFileName, bytes, 0777)
}

func CompileProgram(program string, outFileName string) {
	bytes := make([]byte, 0)

	for i := 0; i < len(program); i++ {
		command := string(program[i])

		switch command {
		case ">":
			bytes = append(bytes, 0)
		case "<":
			bytes = append(bytes, 1)
		case "+":
			bytes = append(bytes, 2)
		case "-":
			bytes = append(bytes, 3)
		case ".":
			bytes = append(bytes, 4)
		case ",":
			bytes = append(bytes, 5)
		case "[":
			bytes = append(bytes, 6)
		case "]":
			bytes = append(bytes, 7)
		}
	}

	dump(bytes, outFileName)
}
