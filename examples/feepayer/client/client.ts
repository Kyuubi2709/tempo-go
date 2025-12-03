import { createClient, http, walletActions, publicActions } from 'viem';
import { privateKeyToAccount } from 'viem/accounts';
import { tempo } from 'tempo.ts/chains';
import { withFeePayer } from 'tempo.ts/viem';
import { config } from './config.js';

async function main() {
  console.log('Tempo Fee Payer Client Example');

  const account = privateKeyToAccount(config.clientPrivateKey);
  console.log('Client address:', account.address);

  const client = createClient({
    account,
    chain: tempo({ feeToken: config.alphaUsdAddress }),
    transport: withFeePayer(
      http(config.getCleanRpcUrl(), {
        fetchOptions: {
          headers: {
            Authorization: config.getAuthHeader(),
          },
        },
      }),
      http(config.feePayerServerUrl)
    ),
  })
    .extend(publicActions)
    .extend(walletActions);

  console.log('Sending transaction via fee payer relay...');

  const hash = await client.sendTransaction({
    to: '0x0000000000000000000000000000000000000000',
    data: '0xdeadbeef',
    value: 0n,
    feeToken: config.alphaUsdAddress,
    feePayer: true,
  } as any);

  console.log('Transaction hash:', hash);
  console.log('Waiting for confirmation...');

  await client.waitForTransactionReceipt({ hash });

  const transaction = await client.getTransaction({ hash });
  console.log('Transaction confirmed!');
  if (transaction.feePayer) {
    console.log('Fee payer address:', transaction.feePayer);
  }
}

main().catch(console.error);
