package byteoperation

func RightMove(repetitions int) ([]byte, int) {
	if repetitions < 1 {
		panic("Repetitions should not be less than 1")
	}

	if repetitions == 1 {
		return []byte{1}, 1
	}

	return []byte{11, byte(repetitions)}, repetitions
}
