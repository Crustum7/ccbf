package bytecode

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
		operation := OperationForPattern(command)
		if operation == nil {
			continue
		}
		data = append(data, operation.opCode)

		switch command {
		case "[":
			opPos := len(data) - 1
			jumpStack = append(jumpStack, opPos)

			data = append(data, 0, 0, 0, 0)
		case "]":
			startOpPos := jumpStack[len(jumpStack)-1]
			jumpStack = jumpStack[:len(jumpStack)-1]

			endOpPos := len(data) - 1

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
