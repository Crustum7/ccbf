package byteoperation

func Add(repetitions int) ([]byte, int) {
	if repetitions < 1 {
		panic("Repetitions should not be less than 1")
	}

	if repetitions == 1 {
		return []byte{3}, 1
	}

	return []byte{9, byte(repetitions)}, repetitions
}
