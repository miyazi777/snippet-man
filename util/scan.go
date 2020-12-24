package util

import (
	"errors"
	"io"
	"strings"

	"github.com/chzyer/readline"
)

func Scan(prompt string, required bool, checkFunc func(string) bool) (string, error) {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          prompt,
		InterruptPrompt: "^C",
	})
	if err != nil {
		return "", err
	}
	defer l.Close()

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			if required == true {
				continue
			} else {
				return "", nil
			}
		} else {
			if checkFunc != nil {
				if checkFunc(line) == false {
					continue
				}
			}
		}
		return line, nil
	}
	return "", errors.New("canceled")
}
