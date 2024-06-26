# Tenant

The `methodazure Tenant` family of commands are intended to provide support to security teams looking to dig deeper into their Azure Tenants.

## Enumerate

Provides detailed information about Tenants.

This command does not require a Subscription ID since it is scoped at the Tenant level.

### Usage

```bash
methodazure tenant enumerate
```

### Help Text

```bash
$ methodazure tenant enumerate --help
Enumerate Tenants for the provided credentials

Usage:
  methodazure tenant enumerate [flags]

Flags:
  -h, --help   help for enumerate

Global Flags:
  -g, --graph-service-endpoint string   Microsoft Graph Service Endpoint (default "https://graph.microsoft.com/.default")
  -o, --output string                   Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string              Path to output file. If blank, will output to STDOUT
  -q, --quiet                           Suppress output
  -s, --subscription-id string          Azure subscription ID
  -v, --verbose                         Verbose output
```
