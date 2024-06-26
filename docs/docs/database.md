# Databases

The `methodazure database` family of commands are intended to provide support to security teams looking to dig deeper into their Azure managed databases.

## Enumerate

Provides detailed information about all managed databases discovered in the given subscription.

### Usage

```bash
methodazure database enumerate --subscription-id <id>
```

### Help Text

```bash
$ methodazure database enumerate --help
Enumerate managed Database instances; retreives managed SQL, Postgres, and Postgres Flexible instance details

Usage:
  methodazure database enumerate [flags]

Flags:
  -h, --help                     help for enumerate
  -s, --subscription-id string   Azure subscription ID

Global Flags:
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -q, --quiet                Suppress output
  -v, --verbose              Verbose output

```
