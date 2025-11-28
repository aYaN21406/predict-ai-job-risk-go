package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== AI Job Automation Risk Report ===")
	fmt.Println()

	for _, job := range Jobs {
		risk := PredictAIRisk(job)
		fmt.Printf("%s - %d%%\n", job.Title, risk)
	}
}
