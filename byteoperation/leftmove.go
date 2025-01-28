package byteoperation

func LeftMove(repetitions int) ([]byte, int) {
	if repetitions < 1 {
		panic("Repetitions should not be less than 1")
	}

	if repetitions == 1 {
		return []byte{2}, 1
	}

	return []byte{12, byte(repetitions)}, repetitions
}
