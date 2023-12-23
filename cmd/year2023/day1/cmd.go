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
	b, err := os.ReadFile(fmt.Sprintf("cmd/year%s/%s/test2.txt", parent, command))
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Score: %d", part1(string(b)))
	logrus.Infof("Score: %d", part2(string(b)))
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

var numTokens = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func parseSpelledNumber(line string) int {
	var digits []rune

	for _, ch := range line {
		if ch >= '0' && ch <= '9' {
			digits = append(digits, ch)
			fmt.Println("Test: ", digits[len(digits)-1])
		} else {
			for k, v := range numTokens {
				if strings.HasPrefix(string(ch), k) {
					fmt.Printf("ch: %v, testingFor: %v value: %v\n", ch, k, v)
					if strings.Contains(line, k) {
						fmt.Printf("will apend!: %v", v)
						digits = append(digits, v)
					}
				}
			}
		}
	}

	if len(digits) <= 0 {
		return 0
	}

	n, err := strconv.Atoi(fmt.Sprintf("%v%v", digits[0], digits[len(digits)-1]))
	fmt.Println(n)
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
