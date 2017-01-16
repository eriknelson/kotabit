package kotabit

import (
	"fmt"
	"github.com/pborman/uuid"
)

func Hello() string {
	return fmt.Println("Hello people, this is kotabit %s!", uuid.UUID())
}
