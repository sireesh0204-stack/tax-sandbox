# Job Search Plan — US Tax Roles, Hyderabad-first

## Target employers

### Tier 1 — Big 4 US-tax delivery centers (Hyderabad unless noted)
- **Deloitte USI** (Hyd) — Tax Senior Analyst / Assistant Manager, US Federal & State
- **EY GDS** (Hyd) — Tax Analyst / Senior, US Core Tax Operations
- **PwC AC India** (Hyd) — Tax Associate/Senior Associate, US Compliance
- **KPMG KGS** (Bengaluru + Hyd) — Tax Analyst, US Federal Tax

### Tier 2 — Strong alternatives
- Grant Thornton Bharat (US tax desk), RSM India, BDO, **Vialto Partners** (global mobility/expat — pairs with our expat module), Ryan LLC

### Tier 3 — Volume
- US CPA firms with India offshore centers; tax-tech firms (Corvee-type) needing tax SMEs; H&R Block back (Blockworks experience is a direct door back in)

## Search strings (LinkedIn/Naukri/Indeed)

- `"US tax" (1065 OR 1120S OR 1040) Hyderabad`
- `"tax analyst" "US tax" (EY OR Deloitte OR PwC OR KPMG) Hyderabad`
- `expat tax OR "global mobility" analyst Bengaluru`
- `international assignment tax India` (Vialto/Ryan niche)

## Apify MCP setup (to let me scrape job boards)

I (ZCode) don't ship Apify connected by default. To wire it up:

1. Create a free account at apify.com → Settings → API & Integrations → copy your **token** (free tier includes some credit; LinkedIn/job scraper actors usually consume paid credits — expect ~$5–20/mo for real volume).
2. Add the MCP server to ZCode (Settings → MCP, or edit MCP config) with:
   - command: `npx -y @apify/actors-mcp-server`
   - env: `APIFY_TOKEN=<your token>`
   - optionally `APIFY_ACTORS`: limit to job-scraper actors like `apify~linkedin-job-scraper`, `apify~indeed-scraper` (exact actor names verified on Apify Store once connected)
3. Restart ZCode → I can then call Apify actors directly to pull listings into `job_log.md`.

Until Apify is connected, fallback = my built-in browser automation for targeted pulls (10–30 listings per run) + WebSearch for fresh postings.

## Application log

| Date | Company | Role | Source | Status | Notes |
|---|---|---|---|---|---|
| | | | | | |

## Resume action items (week of Oct 5)

1. Rewrite headline to: US Tax Analyst — 1065/1120S/1040 multi-state + e-file certification (drop the "Go developer" headline; keep Go as a differentiator bullet, not the brand)
2. Add **honest** software line (see below), WayToTax prep experience to the top third
3. Register ProConnect Tax Online trial / Drake demo for real-tool familiarity before interviews

### Honest software-line rule
Sandbox practice ≠ "Lacerte proficient". Defensible line after this program:
> "Prepared complex returns (1065 manufacturing/rentals, HNI & expat 1040s) in a simulated professional tax workflow; familiar with Lacerte/ProConnect workflows via vendor trials."

Big 4 trains hires on internal tools (CCH/Corptax-class) — they hire for tax judgment, not a specific package. Blockworks already covers "has used professional tax software."
