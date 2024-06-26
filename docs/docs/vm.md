# Virtual Machines (VMs)

The `methodazure vm` family of commands are intended to provide support to security teams looking to dig deeper into the variety of Azure Virtual Machine resources within their subscriptions.

## Enumerate

Provides detailed information about all VMs discovered in the given subscription.

### Usage

```bash
methodazure vm enumerate --subscription-id <id>
```

### Help Text

```bash
$ methodazure vm enumerate --help
Enumerate Virtual Machines

Usage:
  methodazure vm enumerate [flags]

Flags:
  -h, --help                     help for enumerate
  -s, --subscription-id string   Azure subscription ID

Global Flags:
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -q, --quiet                Suppress output
  -v, --verbose              Verbose output
```
