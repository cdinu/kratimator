package rules

import "regexp"

// ReplaceRegex keeps a regular expression to be search and a ReplaceWith string that is used to replace the occurrence
type ReplaceRegex struct {
	SearchRegex *regexp.Regexp
	ReplaceWith string
}
