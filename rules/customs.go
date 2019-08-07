package rules

import "regexp"

var(
	CustomRules = []ReplaceRegex {
		ReplaceRegex{
			SearchRegex: regexp.MustCompile(`(\s)(\d+)-a([\)|\]|\.|,|;|?|!|\s]|$)`),
			ReplaceWith:`${1}${2}‑a${3}`,
		},
		ReplaceRegex{
			SearchRegex: regexp.MustCompile(`(\s)([IVXDCLM]+)-a([\)|\]|\.|,|;|?|!|\s]|$)`),
			ReplaceWith:`${1}${2}‑a${3}`,
		},
		ReplaceRegex{
			SearchRegex: regexp.MustCompile(`prim-minist`),
			ReplaceWith:`prim‑minist`,
		},
		ReplaceRegex{
			SearchRegex: regexp.MustCompile(`prim-plan`),
			ReplaceWith:`prim‑plan`,
		},
		ReplaceRegex{
			SearchRegex: regexp.MustCompile(`(^|[\(|\[|\s])măta([\)|\]|\.|,|;|?|!|\s]|$)`),
			ReplaceWith:`${1}mă‑ta${2}`,  // [are cratimă](https://youtu.be/ALFH3OWCvcc)
		},
	}
)
