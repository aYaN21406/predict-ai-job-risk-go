package main

// Job represents an occupation with AI risk factors
type Job struct {
	Title            string
	RoutineLevel     float64 // 0-1: how routine/repetitive the job is
	PhysicalTasks    float64 // 0-1: how much involves physical work
	Creativity       float64 // 0-1: creative thinking required
	CognitiveTasks   float64 // 0-1: analytical/cognitive work
	SkillRedundancy  float64 // 0-1: how easily skills can be replicated
}
