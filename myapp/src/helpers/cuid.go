package helpers

import (
	"fmt"
	"myapp/utils"

	"github.com/lucsky/cuid"
)

func GenerateCUID() string {
	result := utils.Add(16, 15)
	fmt.Println("Sum:", result)
	fmt.Println("helper")
	return cuid.New()
}
