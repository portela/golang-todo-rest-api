package iteration

const repeatCount = 5

func Repeat(character string) string {
	var acc string
	for i := 0; i < repeatCount; i++ {
		acc = acc + character
	}
	return acc
}
