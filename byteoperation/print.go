package byteoperation

func Print(repetitions int) ([]byte, int) {
	if repetitions < 1 {
		panic("Repetitions should not be less than 1")
	}

	return []byte{5}, 1
}
