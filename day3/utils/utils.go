package utils

type Void struct{}
type Set map[string]Void

var member Void

func StringToSet(value string) Set {
	myset := make(Set)

	chars := []rune(value)
	for i := 0; i < len(chars); i++ {
		myset[string(chars[i])] = member
	}

	return myset
}

func RuneToValue(char rune) int {
	if (int('A') <= int(char)) && (int(char) <= int('Z')) {
		return int(char) - int('A') + 27
	}
	if (int('a') <= int(char)) && (int(char) <= int('z')) {
		return int(char) - int('a') + 1
	}
	return 0
}
