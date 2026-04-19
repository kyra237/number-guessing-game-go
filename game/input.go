package game

import (
	"bufio"
	"strconv"
	"strings"
)

func ReadInt(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return num, nil
}
