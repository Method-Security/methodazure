# Resource Groups

The `methodazure resourcegroup` family of commands are intended to provide support to security teams looking to dig deeper into their Azure managed Kubernetes clusters.

## Enumerate

Provides detailed information about all AKS clusters discovered in the given subscription.

This command requires a Subscription ID and is scoped only to a single subscription.

### Usage

```bash
methodazure resourcegroup enumerate --subscription-id <id>
```

### Help Text

```bash
$ methodazure resourcegroup enumerate --help
Enumerate Resource Groups

Usage:
  methodazure resourcegroup enumerate [flags]

Flags:
  -h, --help   help for enumerate

Global Flags:
  -o, --output string            Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string       Path to output file. If blank, will output to STDOUT
  -q, --quiet                    Suppress output
  -s, --subscription-id string   Azure subscription ID
  -v, --verbose                  Verbose output
```
