package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.ReadFile("mecenas.json")
	if err != nil {
		log.Fatal(err)
	}

	mecenas := make([]string, 0)
	err = json.Unmarshal(f, &mecenas)
	if err != nil {
		log.Fatal(err)
	}

	players := make([]string, len(mecenas))
	copy(players, mecenas)
	tmp := make([]string, 0)

	for i := 0; ; i++ {
		if len(players) == 1 {
			break
		}

		if len(players)%2 == 0 {
			for idx, p := range players {
				if idx%2 == 0 {
					tmp = append(tmp, p)
				}
			}
		} else {
			for idx, p := range players {
				if idx == 0 {
					continue
				}

				if idx%2 == 0 {
					tmp = append(tmp, p)
				}
			}
		}

		players = tmp
		tmp = tmp[:0]
	}

	idxs := make([]int, 0)
	for idx, m := range mecenas {
		if m == players[0] {
			idxs = append(idxs, idx)
		}
	}

	fmt.Printf("Mecena survivor: %+v, and it indexes: %+v\n", players, idxs)
}
