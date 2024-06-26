# Storage

The `methodazure storage` family of commands provide visibility into Azure Storage Accounts

## Enumerate

Provides detailed information about all storage accounts discovered in the given subscription.

### Usage

```bash
methodazure storage enumerate --subscription-id <id>
```

### Help Text

```bash
$ methodazure storage enumerate --help
Enumerate Storage Accounts

Usage:
  methodazure storage enumerate [flags]

Flags:
  -h, --help                     help for enumerate
  -s, --subscription-id string   Azure subscription ID

Global Flags:
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -q, --quiet                Suppress output
  -v, --verbose              Verbose output
```
