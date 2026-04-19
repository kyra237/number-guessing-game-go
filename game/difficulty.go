package game

const (
	Easy   = 1
	Medium = 2
	Hard   = 3
)

var DifficultyToString = map[int]string{
	Easy:   "Easy",
	Medium: "Medium",
	Hard:   "Hard",
}

var DifficultyToChances = map[int]int{
	Easy:   10,
	Medium: 5,
	Hard:   3,
}
