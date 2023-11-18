## Blockchain implementation in golang

## Features:

- `Low level badger db (uses leveldb under the hood)`

- `Using proof of work`

- `Merkle tree`

- `Cli tool`

- `Can be run using docker`

#### Const difficulty (10) :(

- You can use some difficulty algorithms to calculate the difficulty

## Project structure:

```go
shinjiru
│
├── blockchain
│   ├── blockchain.go
│   ├── block.go
│   ├── chain_iter.go
│   ├── merkle.go
│   ├── merkle_test.go
│   ├── proof.go
│   ├── transaction.go
│   ├── tx.go
│   └── utxo.go
├── cli
│   └── cli.go
├── cmd
│   └── main.go
├── Dockerfile
├── go.mod
├── go.sum
├── network
│   └── network.go
├── README.md
├── tmp
│   └── blocks
│       └── keep.txt
└── wallet
    ├── utils.go
    ├── wallet.go
    └── wallets.go
```

## Installation

```sh
git clone https://github.com/kenjitheman/shinjiru
```

## Usage

```sh
cd cmd
go run main.go
```

- Or use Docker:

```sh
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

## Contributing

- Pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change.

## License

- [MIT](./LICENSE)
