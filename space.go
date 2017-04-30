package main

type Space int

const (
	FREE Space = 0 + iota
	NOUGHT
	CROSS
)

func (s Space) ToString() string {
	switch s {
	case FREE:
		return "-"
	case NOUGHT:
		return "O"
	case CROSS:
		return "X"
	}

	return "Error"
}
