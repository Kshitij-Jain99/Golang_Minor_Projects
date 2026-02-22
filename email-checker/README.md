# Email Checker

A simple CLI tool written in Go that checks DNS records for email domains.

The program reads domain names from standard input and outputs a CSV line for each domain showing whether it has mail (MX) records, SPF, and DMARC DNS records.

## Features

- Checks if a domain has MX (Mail Exchange) records
- Checks for SPF (Sender Policy Framework) DNS records
- Checks for DMARC (Domain-based Message Authentication, Reporting & Conformance) DNS records

## Usage

1. **Clone the repository:**
```bash
   git clone https://github.com/Kshitij-Jain99/Golang_Minor_Projects.git
   cd Golang_Minor_Projects/email-checker
```

2. **Build the tool:**
```bash
   go build -o email-checker
```

3. **Run the tool and provide domains to check (one per line):**
```bash
   echo -e "example.com\ngoogle.com" | ./email-checker
```

## Output Format

The tool prints a header followed by a line per domain:
```
domain,hasMX,hasSPF,sprRecord,hasDMARC,dmarcRecord
```

**Example output:**
```
example.com, true, true, v=spf1 include:_spf.google.com ~all, true, v=DMARC1; p=none
```

Each line reports:

| Field | Description |
|---|---|
| `domain` | The input domain |
| `hasMX` | Whether MX records exist |
| `hasSPF` | Whether an SPF TXT record was found |
| `spfRecord` | The SPF record string (if present) |
| `hasDMARC` | Whether a DMARC TXT record was found |
| `dmarcRecord` | The DMARC record string (if present) |

## Dependencies

This project uses only the Go standard library — no external modules required.

| Package | Usage |
|---|---|
| `net` | DNS lookups (MX, TXT records) |
| `bufio` | Buffered input reading |
| `fmt` | Formatted output |
| `log` | Error logging |
| `os` | Standard input/output |
| `strings` | String parsing |

## How It Works

For each domain, the program:

1. Looks up MX records using `net.LookupMX`
2. Retrieves TXT records with `net.LookupTXT` and extracts any SPF entries
3. Checks TXT records under `_dmarc.<domain>` for DMARC policies
