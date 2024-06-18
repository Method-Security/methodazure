# Capabilities

methodazure provides a number of capabilities to cyber security professionals working within Azure, spanning many of Microsoft's most important resource types. Each of the below pages will provide you with a more in depth look at the methodazure capabilities related the specified resource.

- [AKS](./aks.md)
- [Database](./database.md)
- [DNS](./dns.md)
- [NSG](./nsg.md)
- [Resource Groups](./resourcegroup.md)
- [Storage](./storage.md)
- [VMs](./vm.md)
- [VNet](./vnet.md)

## Top Level Flags

methodazure has several top level flags that can be used on any subcommand. These include:

```bash
Flags:
  -h, --help                     help for methodazure
  -o, --output string            Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string       Path to output file. If blank, will output to STDOUT
  -q, --quiet                    Suppress output
  -s, --subscription-id string   Azure subscription ID
  -v, --verbose                  Verbose output
```

## Version Command

Run `methodazure version` to get the exact version information for your binary

## Output Formats

For more information on the various output formats that are supported by methodazure, see the [Output Formats](https://method-security.github.io/docs/output.html) page in our organization wide documentation.
