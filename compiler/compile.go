package compiler

import (
	"bytes"
	"encoding/binary"
	"os"
)

func dump(bytes []byte, outFileName string) {
	os.WriteFile(outFileName, bytes, 0777)
}

func itob(value int32) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		return nil, err
	}
	if len(buf.Bytes()) != 4 {
		panic("Wrong length of byte slice produced by itob()")
	}
	return buf.Bytes(), nil
}

func CompileProgram(program string, outFileName string) {
	data := make([]byte, 0)
	jumpStack := make([]int, 0)

	for i := 0; i < len(program); i++ {
		command := string(program[i])

		switch command {
		case ">":
			data = append(data, 0)
		case "<":
			data = append(data, 1)
		case "+":
			data = append(data, 2)
		case "-":
			data = append(data, 3)
		case ".":
			data = append(data, 4)
		case ",":
			data = append(data, 5)
		case "[":
			opPos := len(data)
			jumpStack = append(jumpStack, opPos)

			data = append(data, 6)
			data = append(data, 0, 0, 0, 0)
		case "]":
			startOpPos := jumpStack[len(jumpStack)-1]
			jumpStack = jumpStack[:len(jumpStack)-1]

			endOpPos := len(data)
			data = append(data, 7)

			toAddress, err := itob(int32(startOpPos + 4))
			if err != nil {
				panic("Could not parse jump address to byte slice")
			}
			data = append(data, toAddress...)

			backAddress, err := itob(int32(endOpPos + 4))
			if err != nil {
				panic("Could not parse jump address to byte slice")
			}
			for j := 0; j < 4; j++ {
				data[startOpPos+1+j] = backAddress[j]
			}
		}
	}

	dump(data, outFileName)
}
