## Subzy 

Subdomain takeover tool which works based on matching response fingerprints from [can-i-take-over-xyz]

![Subzy subdomain takeover](https://i.imgur.com/ggB8zKx.png "Subzy subdomain takeover")

### Installation

go install

or 

go install -v https://github.com/ReXiOP/Subzy
```bash
$ subzy --help
Subdomain takeover tool Mod by Sajid

Usage:
  subzy [command]

Available Commands:
  help        Help about any command
  run         Run subzy
  update      Update local fingerprints.json file
  version     Print subzy version

Flags:
  -h, --help   help for subzy

Use "subzy [command] --help" for more information about a command.
``` 

### Options

Only required flag for `run` subcommand(`r` short version) is either `--target` or `--targets`  

`--target` (string) - Set single or multiple (comma separated) target subdomain/s  
`--targets` (string) - File name/path to list of subdomains    
`--concurrency` (integer) - Number of concurrent checks (default 10)    
`--hide_fails` (boolean) - Hide failed checks and invulnerable subdomains (default false)    
`--https` (boolean) - Use HTTPS by default if protocol not defined on targeted subdomain (default false)  
`--timeout` (integer) - HTTP request timeout in seconds (default 10)  
`--verify_ssl` (boolean) - If set to true, it won't check site with invalid SSL

### Usage

Target subdomain can have protocol defined, if not `http://` will be used by default if `--https` not specifically set to true.

-  List of subdomains
   - ````./subzy run --targets list.txt````

- Single or multiple targets 
  - ```./subzy run --target test.google.com```
  - ```./subzy run --target test.google.com,https://test.yahoo.com```

### Command aliases

Each `subzy` subcommand has its own short version. Running `subzy version` or `subzy v` is the same.

* run - r
* update - u
* version - v
