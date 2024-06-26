# Network Security Groups (NSG)

The `methodazure nsg` family of commands are intended to provide detailed data regarding Azure's Network Security Groups present within an Azure subscription

## Enumerate

Provides detailed information about all Network Security Groups discovered in the given subscription.

This command requires a Subscription ID and is scoped only to a single subscription.

### Usage

```bash
methodazure nsg enumerate --subscription-id <id>
```

### Help Text

```bash
$ methodazure nsg enumerate --help
Enumerate Network Security Groups

Usage:
  methodazure nsg enumerate [flags]

Flags:
  -h, --help   help for enumerate

Global Flags:
  -o, --output string            Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string       Path to output file. If blank, will output to STDOUT
  -q, --quiet                    Suppress output
  -s, --subscription-id string   Azure subscription ID
  -v, --verbose                  Verbose output

```
