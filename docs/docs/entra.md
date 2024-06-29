# Entra ID

The `methodazure entra` family of commands are intended to provide support to security teams looking to dig deeper into their Entra ID instance.

## Enumerate

Provides detailed information about Entra objects.

This command does not require a Subscription ID since it is scoped at the Tenant level.

### Usage

```bash
methodazure entra enumerate
```

### Help Text

```bash
$ methodazure entra enumerate --help
Enumerate Entra ID users, groups, and service principals in a given Tenant

Usage:
  methodazure entra enumerate [flags]

Flags:
  -g, --graph-service-endpoint string   Scope of Microsoft Graph Service Endpoint (e.g. https://graph.microsoft.com/.default), this is automatically defaulted based on --cloud-config value
  -h, --help                            help for enumerate

Global Flags:
  -c, --cloud-config string   Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina) (default "AzurePublic")
  -o, --output string         Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string    Path to output file. If blank, will output to STDOUT
  -q, --quiet                 Suppress output
  -v, --verbose               Verbose output
```
