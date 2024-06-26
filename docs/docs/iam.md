# Azure IAM

The `methodazure iam` family of commands are intended to provide support to security teams looking to dig deeper into the Roles and Role Assignments throughout a subscription.

## Enumerate

Provides detailed information about Entra objects.

### Usage

```bash
methodazure iam enumerate
```

### Help Text

```bash
$ methodazure iam enumerate --help
Enumerate IAM related resources; retreives roles and role assignments in a given subscription

Usage:
  methodazure iam enumerate [flags]

Flags:
  -h, --help                     help for enumerate
  -s, --subscription-id string   Azure subscription ID

Global Flags:
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -q, --quiet                Suppress output
  -v, --verbose              Verbose output
```
