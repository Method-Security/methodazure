# Subscription

The `methodazure subscription` family of commands are intended to provide support to security teams looking to dig deeper into their Azure Subscriptions.

## Enumerate

Provides detailed information about Subscriptions.

This command does not require a Subscription ID since it is scoped at the Tenant level.

### Usage

```bash
methodazure subscription enumerate
```

### Help Text

```bash
$ methodazure subscription enumerate --help
Enumerate Subscriptions for the provided credentials

Usage:
  methodazure subscription enumerate [flags]

Flags:
  -h, --help   help for enumerate

Global Flags:
  -o, --output string                   Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string              Path to output file. If blank, will output to STDOUT
  -q, --quiet                           Suppress output
  -v, --verbose                         Verbose output
```
