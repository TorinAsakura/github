package exceptions

import "fmt"

type TeamIsEmptyException struct {
	TeamID int64
}

func (e *TeamIsEmptyException) Error() string {
	return fmt.Sprintf("Team with ID %d is empty", e.TeamID)
}
