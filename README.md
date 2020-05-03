# gofullnode
CLI for Blockcore FullNodes

Release binaries here: https://github.com/xandronus/gofullnode/releases

Start a blockcore fullnode and then run gofullnode to interact with the fullnode.

```
NAME:
   Gofullnode - Fullnode CLI

USAGE:
   gofullnode [global options] command [command options] [walletname]

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
    gofullnode wallet-create --password mywalletpswd
    gofullnode staking-start --password mywalletpswd
    gofullnode staking-quit
    gofullnode staking-info
    gofullnode wallet-receive
    gofullnode wallet-send --password mywalletpswd 5 xds1qhyfwk5m773r44hq7hrav3zksc5zk6e4l30wc0e
    gofullnode node-add 192.168.0.1
```
