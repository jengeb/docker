package tags

import (
	"fmt"

	"github.com/docker/distribution/registry/api/v2"
)

// DefaultTag is the default tag for the case where no explicit tag is
// specified.
const DefaultTag = "latest"

// ErrTagInvalidFormat is an error type used when the tag name has invalid
// characters or is longer than allowed.
type ErrTagInvalidFormat struct {
	name string
}

func (e ErrTagInvalidFormat) Error() string {
	return fmt.Sprintf("Illegal tag name (%s): only [A-Za-z0-9_.-] are allowed ('.' and '-' are NOT allowed in the initial), minimum 1, maximum 128 in length", e.name)
}

// ValidateTagName validates the name of a tag
func ValidateTagName(name string) error {
	if name == "" {
		return fmt.Errorf("tag name can't be empty")
	}

	if !v2.TagNameAnchoredRegexp.MatchString(name) {
		return ErrTagInvalidFormat{name}
	}
	return nil
}
