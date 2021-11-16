package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "email") {
		return errors.New("email already taken")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("incorrect password")
	}
	return errors.New("incorrect details")
}
