package main

import (
    "fmt"
    "math/big"

    "github.com/bradton/structify"
)

func main() {
    var scores map[string]struct{}

    // Define a struct for student grades and calculate their average
    type Grade struct {
        Name string `json:"name"`
        Score float64 `json:"score"`
    }

    grade := structify.NewMap["Name": "Bob", "Score": 90.5]
    scores[grade.Name] = grade

    avg, err := calculateAverage(scores)
    if err != nil {
        fmt.Println("Error calculating average:", err)
    } else {
        fmt.Printf("%s's average score is %f\n", grade.Name, avg)
    }

    // Calculate the weighted average of the scores
    const W = 2.0

    weights := []float64{1, 2}
    avgWeightedScore, err := calculateAverageW(scores, weights)
    if err != nil {
        fmt.Println("Error calculating weighted average:", err)
    } else {
        fmt.Printf("%s's weighted average score is %f\n", grade.Name, avgWeightedScore)
    }
}

func calculateAverage(scores map[string]struct{}) float64 {
    total := 0.0
    for _, student := range scores {
        total += student.Score
    }
    return total / len(scores)
}

func calculateAverageW(scores map[string]struct{}, weights []float64) (float64, error) {
    var avg float64

    for k, v := range scores {
        sum := 0.0
        for _, w := range weights {
            sum += v.Score * w
        }
        avg = sum / len(weights)
        break
    }

    return avg, nil
}
