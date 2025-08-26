package prompter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Prompter struct{}

func NewPrompter() *Prompter {
	return new(Prompter)
}

func (p *Prompter) Confirm(message string) (bool, error) {
	fmt.Printf("%s ([y]/n): ", message)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	input = strings.TrimSpace(strings.ToLower(input))
	return input == "y" || input == "yes" || input == "", nil
}
