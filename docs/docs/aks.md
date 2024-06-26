# Azure Kubernetes Service (AKS)

The `methodazure aks` family of commands are intended to provide support to security teams looking to dig deeper into their Azure managed Kubernetes clusters.

## Enumerate

Provides detailed information about all AKS clusters discovered in the given subscription.

This command requires a Subscription ID and is scoped only to a single subscription.

### Usage

```bash
methodazure aks enumerate --subscription-id <id>
```

### Help Text

```bash
$ methodazure aks enumerate --help
Enumerate AKS clusters

Usage:
  methodazure aks enumerate [flags]

Flags:
  -h, --help   help for enumerate

Global Flags:
  -o, --output string            Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string       Path to output file. If blank, will output to STDOUT
  -q, --quiet                    Suppress output
  -s, --subscription-id string   Azure subscription ID
  -v, --verbose                  Verbose output
```
