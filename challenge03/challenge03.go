package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Remove all new lines and spaces.
func formatContent(c []byte) []string {
	replacer := strings.NewReplacer("\n", "", " ", "")

	r := replacer.Replace(string(c))
	t := strings.Trim(r, "[]")

	return strings.Split(t, ",")
}

func main() {
	file, err := os.ReadFile("colors.txt")
	if err != nil {
		log.Fatal(err)
	}

	colors := formatContent(file)
	result := make([][]string, 0)
	tmpRes := make([]string, 0)

	for i := 0; i < len(colors); i++ {
		if i == len(colors)-1 {
			if len(tmpRes)%2 == 0 {
				if tmpRes[len(tmpRes)-2] == colors[i] && tmpRes[len(tmpRes)-1] != colors[i] {
					tmpRes = append(tmpRes, colors[i])
					result = append(result, tmpRes)
					continue
				} else {
					result = append(result, tmpRes)
					tmpRes = tmpRes[:0]
					tmpRes = append(tmpRes, colors[i])
					result = append(result, tmpRes)
					continue
				}
			} else {
				if tmpRes[len(tmpRes)-2] == colors[i] && tmpRes[len(tmpRes)-1] != colors[i] {
					tmpRes = append(tmpRes, colors[i])
					result = append(result, tmpRes)
					continue
				} else {
					result = append(result, tmpRes)
					tmpRes = tmpRes[:0]
					tmpRes = append(tmpRes, colors[i])
					result = append(result, tmpRes)
					continue
				}
			}
		}

		// Tmp slice is empty
		if len(tmpRes) == 0 {
			tmpRes = append(tmpRes, colors[i])
			continue
		}

		// Tmp slice have one element
		if len(tmpRes) == 1 {
			if tmpRes[len(tmpRes)-1] == colors[i] {
				result = append(result, tmpRes)
				tmpRes = tmpRes[:0]
				tmpRes = append(tmpRes, colors[i])
				continue
			} else {
				tmpRes = append(tmpRes, colors[i])
				continue
			}
		}

		if len(tmpRes)%2 == 0 {
			if tmpRes[len(tmpRes)-2] == colors[i] && tmpRes[len(tmpRes)-1] != colors[i] {
				tmpRes = append(tmpRes, colors[i])
				continue
			} else {
				result = append(result, tmpRes)
				tmpRes = tmpRes[len(tmpRes)-1:]
				tmpRes = append(tmpRes, colors[i])
				continue
			}
		} else {
			if tmpRes[len(tmpRes)-2] == colors[i] && tmpRes[len(tmpRes)-1] != colors[i] {
				tmpRes = append(tmpRes, colors[i])
				continue
			} else {
				result = append(result, tmpRes)
				tmpRes = tmpRes[len(tmpRes)-1:]
				tmpRes = append(tmpRes, colors[i])
				continue
			}
		}
	}

	sort.Slice(result, func(i, j int) bool {
		l1 := len(result[i])
		l2 := len(result[j])

		return l1 < l2
	})

	fmt.Println(result)
	fmt.Println(len(result[len(result)-1]))
	fmt.Println(result[len(result)-1])
}
