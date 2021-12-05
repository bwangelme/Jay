package validator

import (
	"fmt"
	"regexp"
)

const APP_NAME_MAX_LEN = 16

func ValidateAppName(name string) error {
	if len(name) >= APP_NAME_MAX_LEN {
		return fmt.Errorf("name length gte 16")
	}

	re := regexp.MustCompile(`^[a-zA-Z]{1}[a-zA-Z0-9\\-]{0,15}$`)
	if !re.MatchString(name) {
		return fmt.Errorf("name only allow `-`, digit and letters, and first char cannot be digit `-`")
	}

	return nil
}

func ValidateAppRepo(repo string) error {
	re := regexp.MustCompile(`https?://(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
	if !re.MatchString(repo) {
		return fmt.Errorf("url re match failed")
	}

	return nil
}
