package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day1",
	Long:  `day1`,
	Short: "day1",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf("cmd/year%s/%s/input.txt", parent, command))
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Part2: Score: %d", part2(string(b)))
}

func getNumber(line string) int64 {
	var digits []rune
	for _, ch := range line {
		if unicode.IsDigit(ch) {
			digits = append(digits, ch)
		}
	}

	if len(digits) <= 0 {
		return 0
	}

	n, err := strconv.ParseInt(fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1]), 10, 0)
	if err != nil {
		logrus.Error(err)
	}

	return n
}

func part1(input string) int {
	var total int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		total += int(getNumber(line))
	}

	return total
}

var numTokens = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func parseSpelledNumber(line string) int {
	var digits []rune

	for i := 0; i < len(line); i++ {
		ch := rune(line[i])
		if ch >= '0' && ch <= '9' {
			digits = append(digits, rune(ch))
		} else {
			for digit, word := range numTokens {
				if len(word) > len(line[i:]) {
					break
				}
				if strings.Contains(line[i:len(word)+i], word) {
					digits = append(digits, rune(digit+'0'))
					i += len(word) - 1
				}
			}
		}
	}

	if len(digits) <= 0 {
		return 0
	}

	n, err := strconv.Atoi(fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1]))
	fmt.Printf("%s: [%v]\n", line, n)
	if err != nil {
		logrus.Fatal(err)
	}

	return n
}

func part2(input string) int {
	var total int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		total += parseSpelledNumber(line)
	}

	return total
}
