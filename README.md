# Tax Sandbox — US tax preparation simulator (Go)

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD9?logo=go&logoColor=white)](https://go.dev/dl/)

A production-style US tax-preparation simulator built to:
1. **Train complex returns quickly** (start-to-finish walkthroughs of HNI 1040s, manufacturing 1065s, real-estate 1120S, and expat returns), and
2. **Demonstrate professional tax calculation logic in code.**

Every worksheet and validation rule maps to an actual IRS, California FTB, or state form instruction — no hypothetical math.

Target repository for my public technical portfolio while transitioning into tax-technology roles (Avalara, Vertex, Deloitte USI Tax Technology, Big-4).

**Status:** Early scaffolding. 1065 partnership engine, HNI 1040 intake workflows, and MeF e-file XML stubs are in active development. The `lessons/` directory will ultimately contain fully-worked scenario returns with Go tests, intake questions, and source-document tie-outs.

---

## Why this exists

After 6 years building real tax calculation engines in Go (400+ production files, 1,400+ commits across 27 states and municipal regimes at Corvee/Instead), this project is a public proof-of-concept: a Lacerte/ProConnect-style interface and workflow designed to retrain me on complex returns and to show engineering hiring managers how I structure production tax logic in code.

---

## Quick start

```bash
git clone https://github.com/sireesh0204-stack/tax-sandbox.git
cd tax-sandbox
go run ./cmd/sandbox      # launches the local UI at http://localhost:8080
go test ./...             # run the unit-level tax calculation tests
```

Requirements: [Go 1.22+](https://go.dev/dl/)

---

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
lessons/            # each lesson = one return scenario + intake questions + tests
go.mod              # module github.com/sireeshkumar/tax-sandbox  (requires Go 1.22+)
```

---

## Current modules

| Module | Forms covered | Status |
|---|---|---|
| HNI 1040 + K-1 intake | 1040, 4952, 8959, 1065 K-1 pass-throughs | WIP (`lessons/01_hni_intake_interview.md`) |
| 1065 partnership engine | 1065, Schedule K-1 (individual/entity), 820 | building |
| 1065 e-file XML | MeF transmission + binary attachment stubs | planned |

---

## Lessons

The `lessons/` directory contains fully-worked returns written as **interactive scenarios**:

1. A client fact pattern with hidden contradictions (to train intake questioning),
2. Required source documents,
3. The workpaper tying each amount to a form line,
4. Go code that computes and validates the return against IRS instructions.

Run `go test ./internal/tax/form1065/...` to see the K-1 allocation engine compute distributive shares for a 3-partner real-estate partnership with a 754 election.

---

## License

MIT — use it for practice, or to see how professional tax software structures calculation logic. **No tax advice is rendered**; cross-check any number against the latest form instructions and circulars.

---

*Author: Anasuri Sireesh Kumar — US tax analyst turned tax-technology engineer. Available for US tax-compliance roles through H1B sponsorship or contract-to-hire.*
