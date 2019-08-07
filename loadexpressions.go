package main

import (
	"bufio"
	"kratimator/rules"
	"log"
	"regexp"
	"strings"
	"unicode"
)

type loadExpressionParams struct {
	beforeRegex string
	afterRegex string
	expressions string
}

func loadExpressions(params loadExpressionParams) []rules.ReplaceRegex {
	beforeRegex := params.beforeRegex
	afterRegex := params.afterRegex
	regexBattery := make([]rules.ReplaceRegex, 0)

	scanner := bufio.NewScanner(strings.NewReader(params.expressions))
	for scanner.Scan() {
		expression := scanner.Text()
		if len(expression) == 0 {
			continue
		}

		expressionSafeDash := strings.ReplaceAll(expression, "-", "\\-")
		re := regexp.MustCompile(beforeRegex + expressionSafeDash + afterRegex)

		replaceWith := "${1}" + strings.ReplaceAll(expression, "-", "‑") + "${3}"
		regexBattery = append(regexBattery, rules.ReplaceRegex{SearchRegex: re, ReplaceWith: replaceWith})

		uppercaseExpression := string(unicode.ToUpper([]rune(expression)[0])) + string([]rune(expression)[1:])
		uppercaseExpressionSafe := strings.ReplaceAll(uppercaseExpression, "-", "\\-")
		replaceWith = "${1}" + strings.ReplaceAll(uppercaseExpression, "-", "‑") + "${3}"

		re = regexp.MustCompile(beforeRegex + uppercaseExpressionSafe + afterRegex)

		regexBattery = append(regexBattery, rules.ReplaceRegex{SearchRegex: re, ReplaceWith: replaceWith})
	}

	return regexBattery
}

func loadStandaloneExpressions() []rules.ReplaceRegex {
	const beforeRegex = `(^|[\(|\[|„|\s])(`
	const afterRegex = `)([\)|\]|\.|,|;|“|?|!|\s]|$)`
	regexBattery := loadExpressions(loadExpressionParams{
		beforeRegex: beforeRegex,
		afterRegex:  afterRegex,
		expressions: rules.FullWords,
	})
	log.Printf("Compiled %d regular expressions for standalone words", len(regexBattery))
	return regexBattery
}

func loadTerminatorExpressions() []rules.ReplaceRegex {
	const beforeRegex = `([a-zăîșț])(`
	const afterRegex = `)([\)|\]|\.|,|;|?|!|\s]|$)`
	regexBattery := loadExpressions(loadExpressionParams{
		beforeRegex: beforeRegex,
		afterRegex:  afterRegex,
		expressions: rules.Terminations,
	})
	log.Printf("Compiled %d regular expressions for terminator words", len(regexBattery))
	return regexBattery
}

func loadCustomExpressions() []rules.ReplaceRegex {
	return rules.CustomRules
}