package helper

import "fmt"

func IncorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
