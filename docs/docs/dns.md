# DNS

The `methodazure dns` family of commands are intended to provide insights and visibility into DNS resources that are being managed within Azure.

## Enumerate

Provides detailed information about Azure's managed DNS resources in the given subscription.

### Usage

```bash
methodazure dns enumerate --subscription-id <id>
```

### Help Text

```bash
$ methodazure dns enumerate --help
Enumerate DNS related resources, retreives DNS Zones and Traiffic Manager details

Usage:
  methodazure dns enumerate [flags]

Flags:
  -h, --help                     help for enumerate
  -s, --subscription-id string   Azure subscription ID

Global Flags:
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -q, --quiet                Suppress output
  -v, --verbose              Verbose output
```
