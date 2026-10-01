# README — Tax Sandbox (Portfolio)

**A production-grade US tax-preparation simulator built to (1) train complex returns fast and (2) serve as verifiable proof I can implement tax calculation logic in code.**

Target audience: hiring managers at tax-software companies (Avalara, Vertex) and Big 4 tax-technology teams. Every worksheet and validation rule in this repo maps to an IRS or state form instruction — no hypothetical math.

## Why this exists

After 6 years building real tax calculation engines in Go (400+ files, 1,400+ commits across 27 states and municipal regimes at Corvee/Instead), I needed a public way to:
- prove the domain logic without exposing proprietary code, and
- have a Lacerte/ProConnect-style workflow I can train on for interviews.

## Quick start

```bash
git clone https://github.com/sireeshkumar/tax-sandbox.git
cd tax-sandbox
go run ./cmd/sandbox   # launches the local UI at http://localhost:8080
```

## Project map

```
cmd/sandbox/        # Lacerte-style desktop shell (form navigator + data-entry panes)
internal/tax/
├── form1040/       # HNI / multi-K-1 individual returns
├── form1065/       # partnership returns + K-1 allocation engine
├── form1120s/      # S-corporation returns
├── meef/           # IRS Modernized e-File XML generation/validation
└── state/          # multi-state (CA, NY, NC, ...) worksheets
pkg/workpaper/      # audit trail (source-doc -> amount -> form line)
lessons/            # each lesson = one return scenario + intake questions
go.mod              # module github.com/sireeshkumar/tax-sandbox  (requires Go 1.22+)
```

> **Note:** the engine module is written in Go. If Go is not installed locally, you can still browse the calculation logic in `internal/` and the worked lessons in `lessons/` via any text editor — those are the artifacts recruiters cite. Install Go 1.22+ at https://go.dev/dl/ to run `go test ./...`

## Current modules

| Module | Forms covered | Status |
|---|---|---|
| HNI 1040 + K-1 intake | 1040, 4952, 8959, 1065 K-1 pass-throughs | WIP |
| 1065 partnership engine | 1065, Schedule K-1 (individual/entity), 820 | building |
| 1065 e-file XML | MeF transmission + binary attachment stubs | planned |

## Lessons

`lessons/` contains fully-worked returns written as scenarios:
- client fact pattern
- required source documents
- the questions a real preparer should ask
- the workpaper tying each amount to a form line
- Go code that computes and validates it

Run `go test ./internal/tax/form1065/...` to see the K-1 allocation engine compute distributive shares for a 3-partner real-estate partnership with a 754 election.

## License

MIT — use it for practice or to see how professional tax software structures calculation logic. No tax advice is rendered; cross-check any number against the latest form instructions.
