package DSA

// 232. Implement Queue using Stacks
type MyQueue struct {
	stack1 []int
	stack2 []int
}

// 2390. Removing Stars From a String
func RemoveStars(s string) string {
	stack := []rune{}
	for _, char := range s {
		if char == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return string(stack)
}

// 649. Dota2 Senate
func PredictPartyVictory(senate string) string {
	radiant := []int{}
	dire := []int{}
	for i, char := range senate {
		if char == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}
	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			radiant = radiant[1:]
			dire = dire[1:]
			radiant = append(radiant, len(senate))
		} else {
			dire = dire[1:]
			radiant = radiant[1:]
			dire = append(dire, len(senate))
		}
	}
	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
