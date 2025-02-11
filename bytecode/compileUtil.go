package bytecode

func itob(value int32) []byte {
	return []byte{byte(value >> 24), byte(value >> 16), byte(value >> 8), byte(value)}
}

func assignBytes(to []byte, from []byte) {
	if len(from) != len(to) {
		panic("Number of bytes do not match")
	}

	for i := range len(to) {
		to[i] = from[i]
	}
}

func parameterBytesForOperation(data []byte, opPos int, operation Operation) []byte {
	offset := opPos + 1
	return data[offset : offset+operation.numberOfParameterBytes]
}
