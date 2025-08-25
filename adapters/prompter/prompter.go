package prompter

import (
	"github.com/manifoldco/promptui"
)

type Prompter struct{}

func NewPrompter() *Prompter {
	return new(Prompter)
}

func (p *Prompter) Confirm(message string) (bool, error) {
	prompt := promptui.Select{
		Label: message,
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		return false, err
	}
	return result == "Yes", nil
}
