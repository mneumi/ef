package prompt

import (
	"github.com/mneumi/ef/core/color"

	"github.com/manifoldco/promptui"
)

type PromptItem string
type PromptInputValidator func(string) error

type Prompt struct {
	Label string
	Items []PromptItem
}

func PromptItemDefault(format string, args ...interface{}) PromptItem {
	return PromptItem(color.Cyan(format, args...))
}

func PromptItemPrimary(format string, args ...interface{}) PromptItem {
	return PromptItem(color.Blue(format, args...))
}

func PromptItemSuccess(format string, args ...interface{}) PromptItem {
	return PromptItem(color.Green(format, args...))
}

func PromptItemWarning(format string, args ...interface{}) PromptItem {
	return PromptItem(color.Yellow(format, args...))
}

func PromptItemError(format string, args ...interface{}) PromptItem {
	return PromptItem(color.Red(format, args...))
}

func Select(label string, items []PromptItem) (int, string, error) {
	return (&promptui.Select{
		Label: label,
		Items: items,
	}).Run()
}

func Text(label string, validator PromptInputValidator) (string, error) {
	return (&promptui.Prompt{
		Label:    label,
		Validate: promptui.ValidateFunc(validator),
	}).Run()
}
