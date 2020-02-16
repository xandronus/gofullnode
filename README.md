# goxds
CLI for XDS

Release binaries here: https://github.com/xandronus/goxds/releases

Start XDS fullnode and then run goxds to interact with the fullnode.

```
NAME:
   GoXDS - Fullnode CLI

USAGE:
   goxds.exe [global options] command [command options] [walletname]

VERSION:
   1.0.0

AUTHOR:
   xandronus <xandronus@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command
   Mining:
     staking-start, ss  starts staking
     staking-quit, sq   stops staking
     staking-info, si   get staking info
   Wallet:
     wallet-create, wc   creates a new wallet
     wallet-receive, wr  gets a wallet address to receive funds

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

Examples:
    goxds wallet-create --password mywalletpswd
    goxds staking-start --password mywalletpswd
    goxds staking-quit
    goxds staking-info
    goxds wallet-receive
```
