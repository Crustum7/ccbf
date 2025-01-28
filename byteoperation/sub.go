package byteoperation

func Sub(repetitions int) ([]byte, int) {
	if repetitions < 1 {
		panic("Repetitions should not be less than 1")
	}

	if repetitions == 1 {
		return []byte{4}, 1
	}

	return []byte{10, byte(repetitions)}, repetitions
}
