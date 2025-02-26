# Virtual Network (VNet)

The `methodazure vnet` family of commands are intended to provide visibility and insights into the configuration of Azure Virtual Networks.

## Enumerate

Provides detailed information about all VNets discovered in the given subscription.

### Usage

```bash
methodazure vnet enumerate --subscription-id <id>
```

### Help Text

```bash
$ methodazure vnet enumerate --help
Enumerate VNets

Usage:
  methodazure vnet enumerate [flags]

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
