package icon

type Icon int

const (
	RoundStar Icon = iota
	FailedCross
	SuccessTick
)

func (i Icon) String() string {
	switch i {
	case RoundStar:
		return "✺"
	case FailedCross:
		return "✖︎"
	case SuccessTick:
		return "✔︎"
	}
	return "●"
}
