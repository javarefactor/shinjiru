<h2 align="center">shinjiru - blockchain implementation in golang</h2>

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

## project features:

- `low level badger db (uses leveldb under the hood)`

- `using proof of work`

- `merkle tree`

- `cli tool`

- `can be run using docker`

#### const difficulty (10) 

- you can use some difficulty algorithms to calculate the difficulty

## project structure:

```
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

## installation

- use git clone:

```
git clone https://github.com/kenjitheman/shinjiru
```

## usage

- use go run:

```
cd cmd
go run main.go
```

- or use docker:

```
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
