package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	encMsg := "11610497110107115 102111114 11210897121105110103 9911110010110998101114 11210810197115101 11510497114101"
	message := strings.Builder{}

	for i := 0; i < len(encMsg); i++ {
		if string(encMsg[i]) == " " {
			message.WriteRune(32)
			continue
		}

		if i+3 > len(encMsg) || i+2 > len(encMsg) {
			v, _ := strconv.Atoi(encMsg[i-1:])
			message.WriteRune(rune(v))
			break
		}

		if string(encMsg[i]) == "1" {
			v, _ := strconv.Atoi(encMsg[i : i+3])
			message.WriteRune(rune(v))
			i = i + 2
			continue
		}

		v, _ := strconv.Atoi(encMsg[i : i+2])
		message.WriteRune(rune(v))
		i = i + 1
	}

	fmt.Println(message.String())
}
