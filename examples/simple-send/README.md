# Simple Send

## Overview

A minimal example demonstrating how to create, sign, and broadcast a Type 0x76 transaction to the Tempo network. 

This example shows the basic flow for sending a transaction with a single call.

## Running

1. Copy the environment file and configure your settings:

```bash
cp env.example .env
```

2. Edit `.env` with your values:

```bash
TEMPO_PRIVATE_KEY=0x...        # Your private key
TEMPO_RECIPIENT_ADDRESS=0x...  # Recipient address
```

3. Run the example:

```bash
export $(cat .env | xargs) && go run main.go
```
