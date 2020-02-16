# goxds
CLI for XDS

Release binaries here: https://github.com/xandronus/goxds/releases

Start XDS fullnode and then run goxds to interact with the fullnode.

```
NAME:
   GoXDS - Fullnode CLI

USAGE:
   goxds [global options] command [command options] [walletname]

VERSION:
   0.9.2

AUTHOR:
   xandronus <xandronus@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command
   Mining:
     staking-start, ss  starts staking
     staking-quit, sq   stops staking
     staking-info, si   get staking info
   Node:
     node-add, na  adds a peer
   Wallet:
     wallet-create, wc   creates a new wallet
     wallet-receive, wr  gets a wallet address to receive funds
     wallet-send, ws     sends coins from wallet to address

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

Examples:
    goxds wallet-create --password mywalletpswd
    goxds staking-start --password mywalletpswd
    goxds staking-quit
    goxds staking-info
    goxds wallet-receive
    goxds wallet-send --password mywalletpswd 5 xds1qhyfwk5m773r44hq7hrav3zksc5zk6e4l30wc0e
    goxds node-add 192.168.0.1
```
