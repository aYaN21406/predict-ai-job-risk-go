package main

// ComputeAutomationExposure calculates how exposed a job is to AI automation
func ComputeAutomationExposure(j Job) float64 {
	return (0.5 * j.RoutineLevel) +
		(0.2 * j.PhysicalTasks) +
		(0.3 * (1 - j.Creativity))
}

// ComputeComplexity score = average of cognitive + creativity
func ComputeComplexity(j Job) float64 {
	return (j.CognitiveTasks + j.Creativity) / 2
}

// PredictAIRisk calculates final AI takeover risk percentage
func PredictAIRisk(j Job) int {
	exposure := ComputeAutomationExposure(j)
	complexity := ComputeComplexity(j)

	risk := (0.5 * exposure) +
		(0.25 * j.SkillRedundancy) +
		(0.25 * (1 - complexity))

	return int(risk * 100)
}
