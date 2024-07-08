package main

import (
	"flag"
	"os"

	"github.com/Method-Security/methodazure/cmd"
)

var version = "none"

func main() {
	flag.Parse()

	methodazure := cmd.NewMethodAzure(version)
	methodazure.InitRootCommand()
	methodazure.InitVMCommand()
	methodazure.InitStorageAccountCommand()
	methodazure.InitAKSCommand()
	methodazure.InitDatabaseCommand()
	methodazure.InitDNSCommand()
	methodazure.InitVNetCommand()
	methodazure.InitResourceGroupCommand()
	methodazure.InitNSGCommand()
	methodazure.InitEntraCommand()
	methodazure.InitIAMCommand()
	methodazure.InitTenantCommand()
	methodazure.InitSubscriptionCommand()
	methodazure.InitLoadBalancerCommand()

	if err := methodazure.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
