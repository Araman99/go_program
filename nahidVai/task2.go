package main

import "fmt"

func main() {

	var data interface{} = map[string]interface{}{
		"username":            "nahidhasan98",
		"name":                "Nahid Hasan",
		"honor":               303,
		"clan":                "",
		"leaderboardPosition": 111698,
		"skills":              [3]int{1, 2, 3},
		"ranks": map[string]interface{}{
			"overall": map[string]interface{}{
				"rank":  -5,
				"name":  "5 kyu",
				"color": "yellow",
				"score": 297,
			},
			"languages": map[string]interface{}{
				"go": map[string]interface{}{
					"rank":  -5,
					"name":  "5 kyu",
					"color": "yellow",
					"score": 297,
				},
				"cpp": map[string]interface{}{
					"rank":  -8,
					"name":  "8 kyu",
					"color": "white",
					"score": 2,
				},
			},
		},
		"codeChallenges": map[string]interface{}{
			"totalAuthored":  0,
			"totalCompleted": 32,
		},
	}
	// Exercise:
	// 1. Use this data in your code and find out the value of score of go from languages in ranks.
	// 2. Use type assertion in one line. (Not in multiple line and not part by part)
	// 3. Print the value.

	fmt.Println("--------------------------------")
	goScore := data.(map[string]interface{})["ranks"].(map[string]interface{})["languages"].(map[string]interface{})["go"].(map[string]interface{})["score"]

	fmt.Println("Golang Score: ", goScore)

}
