# Network Security Groups (NSG)

The `methodazure nsg` family of commands are intended to provide detailed data regarding Azure's Network Security Groups present within an Azure subscription

## Enumerate

Provides detailed information about all Network Security Groups discovered in the given subscription.

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
  -h, --help                     help for enumerate
  -s, --subscription-id string   Azure subscription ID

Global Flags:
  -c, --cloud-config string  Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina) (default "AzurePublic")
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -q, --quiet                Suppress output
  -v, --verbose              Verbose output

```
