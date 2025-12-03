# Fee Payer

## Overview

Example fee payer relay server that enables gasless transactions on Tempo.

The server accepts user-signed Type 0x76 transactions, adds its own fee payer signature, and broadcasts the dual-signed transaction to the network. 

## Running

1. Copy the environment file and configure your settings:

```bash
cd examples/feepayer
cp env.example .env
```

2. Edit `.env` with your values:

```env
FEE_PAYER_PORT=3000
TEMPO_RPC_URL=https://rpc.testnet.tempo.xyz
TEMPO_USERNAME=your-username
TEMPO_PASSWORD=your-password
TEMPO_FEE_PAYER_PRIVATE_KEY=0x...
ALPHAUSD_ADDRESS=0x20c0000000000000000000000000000000000001
TEMPO_CHAIN_ID=42424
```

3. Run the server:

```bash
go run cmd/main.go
```

The server exposes `eth_sendRawTransaction` and `eth_sendRawTransactionSync` JSON-RPC methods on the configured port.
