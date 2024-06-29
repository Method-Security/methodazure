var __index = {"config":{"lang":["en"],"separator":"[\\s\\-]+","pipeline":["stopWordFilter"]},"docs":[{"location":"index.html","title":"methodazure Documentation","text":"<p>Hello and welcome to the methodazure documentation. While we always want to provide the most comprehensive documentation possible, we thought you may find the below sections a helpful place to get started.</p> <ul> <li>The Getting Started section provides onboarding material</li> <li>The Development header is the best place to get started on developing on top of and with methodazure</li> <li>See the Docs section for a comprehensive rundown of methodazure capabilities</li> </ul>"},{"location":"index.html#about-methodazure","title":"About methodazure","text":"<p>methodazure provides security operators with a number of data-rich AWS enumeration capabilities to help them gain visibility into their AWS environments. Designed with data-modeling and data-integration needs in mind, methodazure can be used on its own as an interactive CLI, orchestrated as part of a broader data pipeline, or leveraged from within the Method Platform.</p> <p>The number of security-relevant AWS resources that methodazure can enumerate are constantly growing. For the most up to date listing, please see the documentation here</p> <p>To learn more about methodazure, please see the Documentation site for the most detailed information.</p>"},{"location":"index.html#quick-start","title":"Quick Start","text":""},{"location":"index.html#get-methodazure","title":"Get methodazure","text":"<p>For the full list of available installation options, please see the Installation page. For convenience, here are some of the most commonly used options:</p> <ul> <li><code>docker run methodsecurity/methodazure</code></li> <li><code>docker run ghcr.io/method-security/methodazure</code></li> <li>Download the latest binary from the Github Releases page</li> <li>Installation documentation</li> </ul>"},{"location":"index.html#authentication","title":"Authentication","text":"<p>methodazure is built using the AWS Go SDK and leverages the same AWS Credentials that are used by the AWS CLI. Specifically, it looks for the proper environment variables to be exported with credential information. For more information, please see the AWS documentation on how to export AWS credentials as environment variables.</p>"},{"location":"index.html#general-usage","title":"General Usage","text":"<pre><code>methodazure &lt;resource&gt; enumerate --subscription-id &lt;id&gt;\n</code></pre>"},{"location":"index.html#examples","title":"Examples","text":"<pre><code>methodazure storage enumerate --subscription-id &lt;id&gt;\n</code></pre> <pre><code>methodazure vm enumerate --subscription-id &lt;id&gt;\n</code></pre>"},{"location":"index.html#contributing","title":"Contributing","text":"<p>Interested in contributing to methodazure? Please see our organization wide Contribution page.</p>"},{"location":"index.html#want-more","title":"Want More?","text":"<p>If you're looking for an easy way to tie methodazure into your broader cybersecurity workflows, or want to leverage some autonomy to improve your overall security posture, you'll love the broader Method Platform.</p> <p>For more information, visit us here</p>"},{"location":"index.html#community","title":"Community","text":"<p>methodazure is a Method Security open source project.</p> <p>Learn more about Method's open source source work by checking out our other projects here or our organization wide documentation here.</p> <p>Have an idea for a Tool to contribute? Open a Discussion here.</p>"},{"location":"community/community.html","title":"Contributing","text":"<p>For more information on how to get involved in the Method community, please see our organization wide documentation:</p> <ul> <li>Discussions</li> <li>Issues</li> <li>Pull Requests</li> </ul>"},{"location":"development/adding.html","title":"Adding a new capability","text":"<p>By design, methodazure breaks every unique AWS resource into its own top level command. If you are looking to add a brand new capability to the tool, you can take the following steps.</p> <ol> <li>Add a file to <code>cmd/</code> that corresponds to the sub-command name you'd like to add to the <code>methodazure</code> CLI</li> <li>You can use <code>cmd/ec2.go</code> as a template</li> <li>Your file needs to be a member function of the <code>methodazure</code> struct and should be of the form <code>Init&lt;cmd&gt;Command</code></li> <li>Add a new member to the <code>methodazure</code> struct in <code>cmd/root.go</code> that corresponsds to your command name. Remember, the first letter must be capitalized.</li> <li>Call your <code>Init</code> function from <code>main.go</code></li> <li>Add logic to your commands runtime and put it in its own package within <code>internal</code> (e.g., <code>internal/vm</code>)</li> </ol>"},{"location":"development/azure-dev.html","title":"Azure Development","text":"<p>The Azure and Microsoft SDKs and APIs can be difficult to deal with. Compared to AWS for example, they are not clearly documented, data is less organized, and there are different patterns used for different resources.</p> <p>Below are some tips for dealing with the different SDKs.</p>"},{"location":"development/azure-dev.html#using-the-azure-sdk","title":"Using the Azure SDK","text":"<p>The API Reference and Go documentation pages example, are both great resources for understanding query patterns and data structures.</p>"},{"location":"development/azure-dev.html#using-the-microsoft-graph-sdk","title":"Using the Microsoft Graph SDK","text":"<p>The Microsoft Graph SDK is the preferred and up to date method when dealing with Entra ID. For more information on the transition from Azure AD SDKs, view this article.</p> <p>The Microsoft Graph SDK is particularly difficult to deal with. Similar to the Azure SDK, the API Reference is very useful because you need to specify fields to query.</p> <p>Additionally, the Graph Explorer is also useful for testing requests over HTTPS before brining them to the SDK.</p>"},{"location":"development/principles.html","title":"Project Principles","text":""},{"location":"development/principles.html#pre-run-run-post-run","title":"Pre-run -&gt; Run -&gt; Post-run","text":"<p>In the root command, we set a <code>PersistentPreRunE</code> and <code>PersistentPostRunE</code> function that is responsible for initializing the output format and Signal data (in the pre-run) and then write that data in the proper format (in the post-run).</p> <p>Within the Run command that every command must implement, the output of the collected data needs to be written back to the struct's <code>OutputSignal.Content</code> value in order to be properly written out to the caller.</p>"},{"location":"development/principles.html#cmd-vs-internal","title":"Cmd vs Internal","text":"<p>By design, the functionality within each command should focus around parsing the variety of flags and options that the command may need to control capability, passing off all real logic into internal modules.</p>"},{"location":"development/setup.html","title":"Development Setup","text":""},{"location":"development/setup.html#adding-a-new-capability","title":"Adding a new capability","text":"<p>To add a new AWS capability to methodazure, providing new enumeration capabilities to security operators everywhere, please see the adding a new capability page.</p>"},{"location":"development/setup.html#setting-up-your-development-environment","title":"Setting up your development environment","text":"<p>If you've just cloned methodazure for the first time, welcome to the community! We use Palantir's godel to streamline local development and goreleaser to handle the heavy lifting on the release process.</p> <p>To get started with godel, you can run</p> <pre><code>./godelw verify\n</code></pre> <p>This will run a number of checks for us, including linters, tests, and license checks. We run this command as part of our CI pipeline to ensure the codebase is consistently passing tests.</p>"},{"location":"development/setup.html#building-the-cli","title":"Building the CLI","text":"<p>We can use godel to build our CLI locally by running</p> <pre><code>./godelw build\n</code></pre> <p>You should see output in <code>out/build/methoaws/&lt;version&gt;/&lt;os&gt;-&lt;arch&gt;/methodazure</code>.</p> <p>If you'd like to clean this output up, you can run</p> <pre><code>./godelw clean\n</code></pre>"},{"location":"development/setup.html#testing-releases-locally","title":"Testing releases locally","text":"<p>We can use goreleaser locally as well to test our builds. As methodazure uses cosign to sign our artifacts and Docker containers during our CI pipeline, we'll want to skip this step when running locally.</p> <pre><code>goreleaser release --snapshot --clean --skip sign\n</code></pre> <p>This should output binaries, distributable tarballs/zips, as well as docker images to your local machine's Docker registry.</p>"},{"location":"docs/index.html","title":"Capabilities","text":"<p>methodazure provides a number of capabilities to cyber security professionals working within Azure, spanning many of Microsoft's most important resource types. Each of the below pages will provide you with a more in depth look at the methodazure capabilities related the specified resource.</p> <ul> <li>AKS</li> <li>Database</li> <li>DNS</li> <li>NSG</li> <li>Resource Groups</li> <li>Storage</li> <li>VMs</li> <li>VNet</li> </ul>"},{"location":"docs/index.html#top-level-flags","title":"Top Level Flags","text":"<p>methodazure has several top level flags that can be used on any subcommand. These include:</p> <pre><code>Flags:\n  -h, --help                     help for methodazure\n  -o, --output string            Output format (signal, json, yaml). Default value is signal (default \"signal\")\n  -f, --output-file string       Path to output file. If blank, will output to STDOUT\n  -q, --quiet                    Suppress output\n  -s, --subscription-id string   Azure subscription ID\n  -v, --verbose                  Verbose output\n</code></pre>"},{"location":"docs/index.html#version-command","title":"Version Command","text":"<p>Run <code>methodazure version</code> to get the exact version information for your binary</p>"},{"location":"docs/index.html#output-formats","title":"Output Formats","text":"<p>For more information on the various output formats that are supported by methodazure, see the Output Formats page in our organization wide documentation.</p>"},{"location":"docs/aks.html","title":"Azure Kubernetes Service (AKS)","text":"<p>The <code>methodazure aks</code> family of commands are intended to provide support to security teams looking to dig deeper into their Azure managed Kubernetes clusters.</p>"},{"location":"docs/aks.html#enumerate","title":"Enumerate","text":"<p>Provides detailed information about all AKS clusters discovered in the given subscription.</p>"},{"location":"docs/aks.html#usage","title":"Usage","text":"<pre><code>methodazure aks enumerate --subscription-id &lt;id&gt;\n</code></pre>"},{"location":"docs/aks.html#help-text","title":"Help Text","text":"<pre><code>$ methodazure aks enumerate --help\nEnumerate AKS clusters\n\nUsage:\n  methodazure aks enumerate [flags]\n\nFlags:\n  -h, --help                     help for enumerate\n  -s, --subscription-id string   Azure subscription ID\n\nGlobal Flags:\n  -c, --cloud-config string  Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina) (default \"AzurePublic\")\n  -o, --output string        Output format (signal, json, yaml). Default value is signal (default \"signal\")\n  -f, --output-file string   Path to output file. If blank, will output to STDOUT\n  -q, --quiet                Suppress output\n  -v, --verbose              Verbose output\n</code></pre>"},{"location":"docs/database.html","title":"Databases","text":"<p>The <code>methodazure database</code> family of commands are intended to provide support to security teams looking to dig deeper into their Azure managed databases.</p>"},{"location":"docs/database.html#enumerate","title":"Enumerate","text":"<p>Provides detailed information about all managed databases discovered in the given subscription.</p>"},{"location":"docs/database.html#usage","title":"Usage","text":"<pre><code>methodazure database enumerate --subscription-id &lt;id&gt;\n</code></pre>"},{"location":"docs/database.html#help-text","title":"Help Text","text":"<pre><code>$ methodazure database enumerate --help\nEnumerate managed Database instances; retreives managed SQL, Postgres, and Postgres Flexible instance details\n\nUsage:\n  methodazure database enumerate [flags]\n\nFlags:\n  -h, --help                     help for enumerate\n  -s, --subscription-id string   Azure subscription ID\n\nGlobal Flags:\n  -c, --cloud-config string  Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina) (default \"AzurePublic\")\n  -o, --output string        Output format (signal, json, yaml). Default value is signal (default \"signal\")\n  -f, --output-file string   Path to output file. If blank, will output to STDOUT\n  -q, --quiet                Suppress output\n  -v, --verbose              Verbose output\n</code></pre>"},{"location":"docs/dns.html","title":"DNS","text":"<p>The <code>methodazure dns</code> family of commands are intended to provide insights and visibility into DNS resources that are being managed within Azure.</p>"},{"location":"docs/dns.html#enumerate","title":"Enumerate","text":"<p>Provides detailed information about Azure's managed DNS resources in the given subscription.</p>"},{"location":"docs/dns.html#usage","title":"Usage","text":"<pre><code>methodazure dns enumerate --subscription-id &lt;id&gt;\n</code></pre>"},{"location":"docs/dns.html#help-text","title":"Help Text","text":"<pre><code>$ methodazure dns enumerate --help\nEnumerate DNS related resources, retreives DNS Zones and Traiffic Manager details\n\nUsage:\n  methodazure dns enumerate [flags]\n\nFlags:\n  -h, --help                     help for enumerate\n  -s, --subscription-id string   Azure subscription ID\n\nGlobal Flags:\n  -c, --cloud-config string  Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina) (default \"AzurePublic\")\n  -o, --output string        Output format (signal, json, yaml). Default value is signal (default \"signal\")\n  -f, --output-file string   Path to output file. If blank, will output to STDOUT\n  -q, --quiet                Suppress output\n  -v, --verbose              Verbose output\n</code></pre>"},{"location":"docs/entra.html","title":"Entra ID","text":"<p>The <code>methodazure entra</code> family of commands are intended to provide support to security teams looking to dig deeper into their Entra ID instance.</p>"},{"location":"docs/entra.html#enumerate","title":"Enumerate","text":"<p>Provides detailed information about Entra objects.</p> <p>This command does not require a Subscription ID since it is scoped at the Tenant level.</p>"},{"location":"docs/entra.html#usage","title":"Usage","text":"<pre><code>methodazure entra enumerate\n</code></pre>"},{"location":"docs/entra.html#help-text","title":"Help Text","text":"<pre><code>$ methodazure entra enumerate --help\nEnumerate Entra ID users, groups, and service principals in a given Tenant\n\nUsage:\n  methodazure entra enumerate [flags]\n\nFlags:\n  -g, --graph-service-endpoint string   Scope of Microsoft Graph Service Endpoint (e.g. https://graph.microsoft.com/.default), this is automatically defaulted based on --cloud-config value\n  -h, --help                            help for enumerate\n\nGlobal Flags:\n  -c, --cloud-config string   Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina) (default \"AzurePublic\")\n  -o, --output string         Output format (signal, json, yaml). Default value is signal (default \"signal\")\n  -f, --output-file string    Path to output file. If blank, will output to STDOUT\n  -q, --quiet                 Suppress output\n  -v, --verbose               Verbose output\n</code></pre>"},{"location":"docs/iam.html","title":"Azure IAM","text":"<p>The <code>methodazure iam</code> family of commands are intended to provide support to security teams looking to dig deeper into the Roles and Role Assignments throughout a subscription.</p>"},{"location":"docs/iam.html#enumerate","title":"Enumerate","text":"<p>Provides detailed information about Entra objects.</p>"},{"location":"docs/iam.html#usage","title":"Usage","text":"<pre><code>methodazure iam enumerate\n</code></pre>"},{"location":"docs/iam.html#help-text","title":"Help Text","text":"<pre><code>$ methodazure iam enumerate --help\nEnumerate IAM related resources; retreives roles and role assignments in a given subscription\n\nUsage:\n  methodazure iam enumerate [flags]\n\nFlags:\n  -h, --help                     help for enumerate\n  -s, --subscription-id string   Azure subscription ID\n\nGlobal Flags:\n  -c, --cloud-config string  Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina) (default \"AzurePublic\")\n  -o, --output string        Output format (signal, json, yaml). Default value is signal (default \"signal\")\n  -f, --output-file string   Path to output file. If blank, will output to STDOUT\n  -q, --quiet                Suppress output\n  -v, --verbose              Verbose output\n</code></pre>"},{"location":"docs/nsg.html","title":"Network Security Groups (NSG)","text":"<p>The <code>methodazure nsg</code> family of commands are intended to provide detailed data regarding Azure's Network Security Groups present within an Azure subscription</p>"},{"location":"docs/nsg.html#enumerate","title":"Enumerate","text":"<p>Provides detailed information about all Network Security Groups discovered in the given subscription.</p>"},{"location":"docs/nsg.html#usage","title":"Usage","text":"<pre><code>methodazure nsg enumerate --subscription-id &lt;id&gt;\n</code></pre>"},{"location":"docs/nsg.html#help-text","title":"Help Text","text":"<pre><code>$ methodazure nsg enumerate --help\nEnumerate Network Security Groups\n\nUsage:\n  methodazure nsg enumerate [flags]\n\nFlags:\n  -h, --help                     help for enumerate\n  -s, --subscription-id string   Azure subscription ID\n\nGlobal Flags:\n  -c, --cloud-config string  Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina) (default \"AzurePublic\")\n  -o, --output string        Output format (signal, json, yaml). Default value is signal (default \"signal\")\n  -f, --output-file string   Path to output file. If blank, will output to STDOUT\n  -q, --quiet                Suppress output\n  -v, --verbose              Verbose output\n</code></pre>"},{"location":"docs/resourcegroup.html","title":"Resource Groups","text":"<p>The <code>methodazure resourcegroup</code> family of commands are intended to provide support to security teams looking to dig deeper into their Azure managed Kubernetes clusters.</p>"},{"location":"docs/resourcegroup.html#enumerate","title":"Enumerate","text":"<p>Provides detailed information about all AKS clusters discovered in the given subscription.</p>"},{"location":"docs/resourcegroup.html#usage","title":"Usage","text":"<pre><code>methodazure resourcegroup enumerate --subscription-id &lt;id&gt;\n</code></pre>"},{"location":"docs/resourcegroup.html#help-text","title":"Help Text","text":"<pre><code>$ methodazure resourcegroup enumerate --help\nEnumerate Resource Groups\n\nUsage:\n  methodazure resourcegroup enumerate [flags]\n\nFlags:\n  -h, --help                     help for enumerate\n  -s, --subscription-id string   Azure subscription ID\n\nGlobal Flags:\n  -c, --cloud-config string  Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina) (default \"AzurePublic\")\n  -o, --output string        Output format (signal, json, yaml). Default value is signal (default \"signal\")\n  -f, --output-file string   Path to output file. If blank, will output to STDOUT\n  -q, --quiet                Suppress output\n  -v, --verbose              Verbose output\n</code></pre>"},{"location":"docs/storage.html","title":"Storage","text":"<p>The <code>methodazure storage</code> family of commands provide visibility into Azure Storage Accounts</p>"},{"location":"docs/storage.html#enumerate","title":"Enumerate","text":"<p>Provides detailed information about all storage accounts discovered in the given subscription.</p>"},{"location":"docs/storage.html#usage","title":"Usage","text":"<pre><code>methodazure storage enumerate --subscription-id &lt;id&gt;\n</code></pre>"},{"location":"docs/storage.html#help-text","title":"Help Text","text":"<pre><code>$ methodazure storage enumerate --help\nEnumerate Storage Accounts\n\nUsage:\n  methodazure storage enumerate [flags]\n\nFlags:\n  -h, --help                     help for enumerate\n  -s, --subscription-id string   Azure subscription ID\n\nGlobal Flags:\n  -c, --cloud-config string  Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina) (default \"AzurePublic\")\n  -o, --output string        Output format (signal, json, yaml). Default value is signal (default \"signal\")\n  -f, --output-file string   Path to output file. If blank, will output to STDOUT\n  -q, --quiet                Suppress output\n  -v, --verbose              Verbose output\n</code></pre>"},{"location":"docs/subscription.html","title":"Subscription","text":"<p>The <code>methodazure subscription</code> family of commands are intended to provide support to security teams looking to dig deeper into their Azure Subscriptions.</p>"},{"location":"docs/subscription.html#enumerate","title":"Enumerate","text":"<p>Provides detailed information about Subscriptions.</p> <p>This command does not require a Subscription ID since it is scoped at the Tenant level.</p>"},{"location":"docs/subscription.html#usage","title":"Usage","text":"<pre><code>methodazure subscription enumerate\n</code></pre>"},{"location":"docs/subscription.html#help-text","title":"Help Text","text":"<pre><code>$ methodazure subscription enumerate --help\nEnumerate Subscriptions for the provided credentials\n\nUsage:\n  methodazure subscription enumerate [flags]\n\nFlags:\n  -h, --help             help for enumerate\n  -t, --try-all-clouds   Attempt to list subscriptions in AzurePublic, Azure Government, and AzureChina; if true overrides cloud-config flag\n\nGlobal Flags:\n  -c, --cloud-config string   Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina) (default \"AzurePublic\")\n  -o, --output string         Output format (signal, json, yaml). Default value is signal (default \"signal\")\n  -f, --output-file string    Path to output file. If blank, will output to STDOUT\n  -q, --quiet                 Suppress output\n  -v, --verbose               Verbose output\n</code></pre>"},{"location":"docs/tenant.html","title":"Tenant","text":"<p>The <code>methodazure Tenant</code> family of commands are intended to provide support to security teams looking to dig deeper into their Azure Tenants.</p>"},{"location":"docs/tenant.html#enumerate","title":"Enumerate","text":"<p>Provides detailed information about Tenants.</p> <p>This command does not require a Subscription ID since it is scoped at the Tenant level.</p>"},{"location":"docs/tenant.html#usage","title":"Usage","text":"<pre><code>methodazure tenant enumerate\n</code></pre>"},{"location":"docs/tenant.html#help-text","title":"Help Text","text":"<pre><code>$ methodazure tenant enumerate --help\nEnumerate Tenants for the provided credentials\n\nUsage:\n  methodazure tenant enumerate [flags]\n\nFlags:\n  -h, --help   help for enumerate\n\nGlobal Flags:\n  -c, --cloud-config string  Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina) (default \"AzurePublic\")\n  -o, --output string        Output format (signal, json, yaml). Default value is signal (default \"signal\")\n  -f, --output-file string   Path to output file. If blank, will output to STDOUT\n  -q, --quiet                Suppress output\n  -v, --verbose              Verbose output\n</code></pre>"},{"location":"docs/vm.html","title":"Virtual Machines (VMs)","text":"<p>The <code>methodazure vm</code> family of commands are intended to provide support to security teams looking to dig deeper into the variety of Azure Virtual Machine resources within their subscriptions.</p>"},{"location":"docs/vm.html#enumerate","title":"Enumerate","text":"<p>Provides detailed information about all VMs discovered in the given subscription.</p>"},{"location":"docs/vm.html#usage","title":"Usage","text":"<pre><code>methodazure vm enumerate --subscription-id &lt;id&gt;\n</code></pre>"},{"location":"docs/vm.html#help-text","title":"Help Text","text":"<pre><code>$ methodazure vm enumerate --help\nEnumerate Virtual Machines\n\nUsage:\n  methodazure vm enumerate [flags]\n\nFlags:\n  -h, --help                     help for enumerate\n  -s, --subscription-id string   Azure subscription ID\n\nGlobal Flags:\n  -c, --cloud-config string  Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina) (default \"AzurePublic\")\n  -o, --output string        Output format (signal, json, yaml). Default value is signal (default \"signal\")\n  -f, --output-file string   Path to output file. If blank, will output to STDOUT\n  -q, --quiet                Suppress output\n  -v, --verbose              Verbose output\n</code></pre>"},{"location":"docs/vnet.html","title":"Virtual Network (VNet)","text":"<p>The <code>methodazure vnet</code> family of commands are intended to provide visibility and insights into the configuration of Azure Virtual Networks.</p>"},{"location":"docs/vnet.html#enumerate","title":"Enumerate","text":"<p>Provides detailed information about all VNets discovered in the given subscription.</p>"},{"location":"docs/vnet.html#usage","title":"Usage","text":"<pre><code>methodazure vnet enumerate --subscription-id &lt;id&gt;\n</code></pre>"},{"location":"docs/vnet.html#help-text","title":"Help Text","text":"<pre><code>$ methodazure vnet enumerate --help\nEnumerate VNets\n\nUsage:\n  methodazure vnet enumerate [flags]\n\nFlags:\n  -h, --help                     help for enumerate\n  -s, --subscription-id string   Azure subscription ID\n\nGlobal Flags:\n  -c, --cloud-config string  Azure Cloud to use (AzurePublic, AzureGovernment, AzureChina) (default \"AzurePublic\")\n  -o, --output string        Output format (signal, json, yaml). Default value is signal (default \"signal\")\n  -f, --output-file string   Path to output file. If blank, will output to STDOUT\n  -q, --quiet                Suppress output\n  -v, --verbose              Verbose output\n</code></pre>"},{"location":"getting-started/basic-usage.html","title":"Basic Usage","text":"<p>Before you get started, you will need to export AWS credentials that you want methodazure to utilize as environment variables. For more documentation on how to do this, please see the Amazon documentation here.</p>"},{"location":"getting-started/basic-usage.html#authentication","title":"Authentication","text":"<p>methodazure leverages the Microsoft Azure SDK's environment variable mechanism for accessing the appropriate credentials to communicate with Azure resources. Microsoft's developer blog goes into more detail here but essentially if the <code>AZURE_CLIENT_ID</code>, <code>AZURE_CLIENT_SECRET</code>, and <code>AZURE_TENANT_ID</code> environment variables are set, we can leverage the Azure SDK to authenticate.</p>"},{"location":"getting-started/basic-usage.html#binaries","title":"Binaries","text":"<p>Running as a binary means you don't need to do anything additional for methodazure to leverage the environment variables you have already exported. You can test that things are working properly by running:</p> <pre><code>methodazure vnet enumerate --subscription-id &lt;id&gt; --output json\n</code></pre>"},{"location":"getting-started/basic-usage.html#docker","title":"Docker","text":"<p>Running methodazure within a Docker container requires that you pass the Azure credential environment variables into the container. This can be done with the following command:</p> <pre><code>docker run \\\n  -e AZURE_TENANT_ID=$AZURE_TENANT_ID \\\n  -e AZURE_CLIENT_ID=$AZURE_CLIENT_ID \\\n  -e AZURE_CLIENT_SECRET=$AZURE_CLIENT_SECRET \\\n  ghcr.io/method-security/methodazure:0.0.1 vnet enumerate --subscription-id &lt;id&gt; --output json\n</code></pre>"},{"location":"getting-started/installation.html","title":"Getting Started","text":"<p>If you are just getting started with methodazure, welcome! This guide will walk you through the process of going zero to one with the tool.</p>"},{"location":"getting-started/installation.html#installation","title":"Installation","text":"<p>methodazure is provided in several convenient form factors, including statically compiled binary images on a variety of architectures as well as a Docker image for both x86 and ARM machines.</p> <p>If you do not see an architecture that you require, please open a Discussion to propose adding it.</p>"},{"location":"getting-started/installation.html#binaries","title":"Binaries","text":"<p>methodazure currently supports statically compiled binaries across the following operating systems and architectures:</p> OS Architecture Linux 386 Linux arm (goarm 7) Linux amd64 Linux arm64 MacOS amd64 MacOS arm64 Windows amd64 <p>The latest binaries can be downloaded directly from Github.</p>"},{"location":"getting-started/installation.html#docker","title":"Docker","text":"<p>Docker images for methodazure are hosted in both Github Container Registry as well as on Docker Hub and can be pulled via:</p> <pre><code>docker pull ghcr.io/method-security/methodazure\n</code></pre> <pre><code>docker pull methodsecurity/methodazure\n</code></pre>"}]}