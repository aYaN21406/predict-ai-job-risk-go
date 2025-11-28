# AI Job Risk Predictor (Go Edition)

A clean, production-style Go implementation that predicts AI automation risk for various job roles.

## Features

- Job database modeled as Go structs
- AI automation exposure model
- Skill redundancy + complexity scoring
- Final `AI Risk %` output for each job
- Very clean and readable code

## Project Structure

```
predict-ai-job-risk-go/
|-- go.mod          # Go module definition
|-- main.go         # Program entry point + formatted report
|-- jobs.go         # Job struct and scoring factors
|-- risk.go         # Weighting formula and computation
|-- data.go         # Preloaded dataset of 25+ jobs
```

## How to Run

```bash
# Clone the repository
git clone https://github.com/aYaN21406/predict-ai-job-risk-go.git
cd predict-ai-job-risk-go

# Run the program
go run .
```

## Sample Output

```
=== AI Job Automation Risk Report ===
Cashier — 87%
Data Entry Clerk — 84%
Telemarketer — 82%
Bookkeeper — 79%
Software Developer — 42%
Nurse — 28%
Physical Therapist — 19%
```

## Risk Calculation Formula

```go
risk := (0.5 * exposure) +
        (0.25 * skillRedundancy) +
        (0.25 * (1 - complexity))
```

Where:
- **exposure**: AI automation exposure score (0-1)
- **skillRedundancy**: How easily AI can replicate required skills
- **complexity**: Job complexity score (higher = harder to automate)

## Why This Project Is Resume-Ready

- Demonstrates Go proficiency
- Includes data modeling and analytics
- Clean modular code
- Real-world AI automation concept
- Strong system-design presentation

Perfect for backend engineering and system-level roles.

## Related Projects

- [predict-ai-job-risk-r](https://github.com/aYaN21406/predict-ai-job-risk-r) - R version with statistical analysis
- [predict-ai-job-risk-c](https://github.com/aYaN21406/predict-ai-job-risk-c) - C version with CLI tool
- [predict-ai-job-risk-sql](https://github.com/aYaN21406/predict-ai-job-risk-sql) - SQL version with database analytics

## License

MIT License
