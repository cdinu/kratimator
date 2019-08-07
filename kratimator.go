package main

var (
	standaloneReplacer = loadStandaloneExpressions()
	terminatorReplacer = loadTerminatorExpressions()
	customsReplacer = loadCustomExpressions()
)

func kratimate(src string) string {
	result := src

	for _, pp := range standaloneReplacer {
		result = pp.SearchRegex.ReplaceAllString(result, pp.ReplaceWith)
	}

	for _, pp := range terminatorReplacer {
		result = pp.SearchRegex.ReplaceAllString(result, pp.ReplaceWith)
	}

	for _, pp := range customsReplacer {
		result = pp.SearchRegex.ReplaceAllString(result, pp.ReplaceWith)
	}

	return result;
}
