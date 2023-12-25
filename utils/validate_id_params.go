package utils

import "strconv"

func ValidateRequestIdParams(id string) bool {
	_, err := strconv.Atoi(id)

	return err == nil
}
