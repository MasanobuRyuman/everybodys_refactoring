package utility

import (
	"github.com/google/uuid"
	"fmt"
)

func GetUuid() string{
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Errorf("failed to find user: %w",err)	
  }
	return u.String()
}