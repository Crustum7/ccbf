package byteoperation

func Input(repetitions int) ([]byte, int) {
	if repetitions < 1 {
		panic("Repetitions should not be less than 1")
	}

	return []byte{6}, 1
}
