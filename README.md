# Build a Blockchain from Scratch in Go

**Suggested GitHub repository**
- **Name:** `build-blockchain-from-scratch-in-go`
- **Description:** Learn blockchain fundamentals by implementing blocks, hashing, proof-of-work, and simple transactions in Go.

In this project, we will learn how to build a blockchain from scratch in Go by implementing block structures and hash functions. We’ll also explore the proof of work algorithm and transaction management.

## We will learn to:
- Implement the fundamental concepts and principles of blockchain technology.
- Implement the core blockchain components.
- Implement the proof of work (PoW) consensus algorithm.
- Develop proficiency in Go by implementing a blockchain from scratch.

## Skills
- Blockchain Development
- Blockchain Mining

## Prerequisites
- Intermediate understanding of Go
- Basic understanding of blockchain

## Project Description
In this project, we will get hands-on experience building a functional blockchain system from scratch using Go. We will explore the core concepts of blockchain, including blocks, hash functions, the proof of work algorithm, and transaction management.

We’ll begin by implementing the block structure and blockchain structures, and learn how hashes are computed and utilized within the blockchain by implementing hash functions. Next, we’ll create blocks, including the Genesis block and blockchain. We’ll implement the proof of work algorithm responsible for mining blocks in the blockchain and also delve into transactions and wallets within the blockchain system, creating a transaction structure that enables the secure and reliable transfer of assets within the blockchain network. Also, we’ll design and implement wallets with associated functions, allowing users to manage their digital assets effectively.

## Task 0: Get Started
A blockchain is a `distributed ledger` that consists of a `chain of blocks` linked together. They provide benefits such as `decentralization`, `transparency`, and `enhanced security` through `cryptographic algorithms`.

A `block` is a fundamental `blockchain` unit that contains a single or a set of `transactions` or digital records. Each `block` contains a `unique hash` generated from its data and the `previous block’s hash`, forming an immutable chain.

## Task 1: Create Blocks and Blockchain Structures
In this task, we’ll create structures representing a block and blockchain.

## Task 2: Compute the Hash of a Block
After creating the structures of Block and BlockChain, implement a function to compute the hash of a block. Hash functions are vital in a blockchain because they provide data integrity and security. Calculate the hash value based on the block’s data and previous hash using the MD5 hashing algorithm.

## Task 3: Create a New Block with the Genesis Block
The very first block in a blockchain is known as a `Genesis block`. It serves as the foundation of the entire blockchain network and is often created during the initialization of a new blockchain system. Since there are no preceding blocks in the chain, the Genesis block is unique because it does not reference previous blocks. You’ll also create a function to generate new blocks in the blockchain, each connected to the previous block.

## Task 4: Initialize and Add a Block to the Blockchain
In Task 3, you developed methods to create a Genesis block and add new blocks to the blockchain. In this task, we will initialize a new blockchain and add a block to it. This involves creating an instance of the blockchain, generating the Genesis block, and then adding subsequent blocks to the chain.

## Task 5: Test the Blockchain
Now that you’ve got the initial structure of the blockchain in place, set up a blockchain, add blocks to it, and examine its structure.

## Task 6: Create and Instantiate the Proof of Work Structure
Blockchain networks employ the `proof of work (PoW) consensus algorithm` to validate transactions and add new blocks. The PoW algorithm requires miners to solve a complex mathematical puzzle, which requires a lot of computational power to solve, ensuring a secure blockchain network.

In the PoW algorithm, you have a `target` that represents the `difficulty level` for mining new blocks. It determines the criteria the block’s hash must meet, ensuring that mining requires significant computational work. When a new block is mined, the PoW algorithm repeatedly calculates the hash of the block until it meets the difficulty criteria specified by the target. `Miners` have to find a hash that, compared to the target, has the required number of leading zeros.

*Note:* The number of leading zeros required in the block’s hash determines the difficulty level. The more leading zeros required, the higher the difficulty.

You will start by creating a `proof of work` structure and a function that generates an instance of it for a given block after computing the target value.

*Note:* The target for mining a new block will remain constant in this example; however, in a real-world blockchain implementation, it is typically `adjusted periodically` to adapt to changes in the network’s computational power and maintain a consistent `block creation rate`.

## Task 7: Mine a Block
In this task, implement the mining process for a block in a blockchain using the PoW algorithm. Mining involves finding a valid hash for a given block that meets the difficulty criteria specified by the target. One important component in the mining process is the use of a `nonce`. You’ll start by initializing the data to compute the block’s hash and then perform the mining process. Simulate the mining process by repeatedly computing the hash of the block’s data with different nonces until a hash less than the target is found.

## Task 8: Validate the Proof of Work Algorithm
The PoW algorithm involves `finding a valid block hash` that meets a specific `target`. After running the PoW algorithm and obtaining the nonce, you will use it to derive the hash that meets the target. Implement a function that validates this derived hash and essentially validates the work done by the PoW algorithm.

*Note:* This task assumes that the Block struct includes a Nonce field (set during mining in Task 7), which will be used here for validation.

## Task 9: Test the Proof of Work Algorithm
Once you have computed the `nonce`, you must store it inside that block. Currently, the Block structure has no attribute to store the nonce. Remember that you are now computing the block’s hash in the PoW algorithm; therefore, there is no need to do so separately for the block. Update the Block struct and the CreateBlock function to incorporate the PoW algorithm.

## Task 10: Create the Transaction Structure
In a blockchain, a `transaction` represents data transfer from one participant to another within the network. Transactions generally include data such as the `sender’s and receiver’s addresses` and the `amount` or quantity being transferred. When a transaction is initiated, it undergoes validation to ensure its authenticity. Once validated, the transaction is included in a block and added to the blockchain through mining.

A `coinbase transaction` is a special type of transaction that is primarily used to `reward miners` who successfully `mine a new block`. It is the `first transaction` in each block and does not have a specific sender like regular transactions. Instead, it is `created by the miner` who successfully solves the cryptographic puzzle associated with mining a block. It `rewards` the miner by specifying the `recipient’s address` and an `amount as a reward` for their mining efforts.

You will create a new structure for transactions and update the functionality for adding a block such that a new coinbase transaction is created when adding a block to the blockchain.

## Task 11: Add Transactions to the Block
The transactions are stored inside the block, but so far there is no attribute to store the transactions in a block. You have to update the Block structure and the CreateBlock function to include a new attribute for storing transactions and receiving a list of transactions as input. You’ll also update the Genesis function to create a coinbase transaction that will reward the miner who successfully mines the block.

## Task 12: Test the Transactions
Now that the transactions can be added to a block, add blocks to the blockchain with transactions to see if everything works as expected.

## Task 13: Add Wallet Functionality
In blockchain, a `wallet` is a digital container that `stores the cryptographic keys` used to access and manage a user’s assets. It consists of a `private key` and a `public key`. The `public key` is shared within the blockchain network and is mainly used to `receive funds` and `verify the authenticity of digital signatures`. With a public key, anyone can encrypt data that can only be decrypted by the corresponding private key. The `private key` is a secret cryptographic key that should be kept secure and known only to the owner. It `creates digital signatures`, `signs` transactions, and `decrypts data` encrypted with the corresponding public key.

You will use RSA for key generation, sign transactions, and verify transaction signatures. You’ll implement a wallet structure, generate public and private keys, and create a new wallet based on the generated keys.

## Task 14: Sign and Verify the Transactions
To ensure security and integrity in a blockchain system, signing and verifying transactions is essential. A `transaction is signed using a private key` that provides a way to prove the authenticity of the transaction. The sender associates a transaction with their private key by `appending a digital signature` to a transaction, allowing other participants in the blockchain to `verify that the sender authorized the transaction`.

A transaction is verified using the `sender’s public key`, ensuring it has not been tampered with since it was signed. The `public key` is available to all participants and can be used by any participant to `verify the signature`. If the transaction data has been modified in any way, the verification process will fail, indicating that the transaction has been tampered with. You will implement functions to `sign and verify` transactions using a particular sender’s `public` and `private` keys.


## Task 15: Test the Wallets
Now that wallet functionality has been added, test it in the blockchain implementation. You will create wallets for two users, simulate a transaction between them, sign the transaction, verify its authenticity, and display the transaction details.

## Executive Summary
This project walks you through building a working (educational) blockchain in Go from first principles. You start by defining the core data structures (a `Block` and a `Blockchain`) and linking blocks together using cryptographic hashes so that each block commits to the previous one, forming an append-only chain beginning with a `Genesis` block.

From there, you implement and test the key mechanics that make blockchains resilient to tampering:
- **Hashing & integrity (Tasks 1–5):** create blocks, compute a block hash from its contents + previous hash, initialize a chain, append new blocks, and inspect the resulting structure.
- **Proof of Work (Tasks 6–9):** define a PoW target (difficulty), “mine” blocks by searching for a `nonce` that produces a valid hash, validate the mined result, and evolve the block creation flow so PoW is integrated into block creation.

Finally, you extend the chain from “data blocks” into a simple value-transfer system:
- **Transactions (Tasks 10–12):** model transactions (including a miner-reward `coinbase` transaction), store transactions inside blocks, create a coinbase transaction in the genesis block, and test adding blocks containing transactions.
- **Wallets & signatures (Tasks 13–15):** generate wallet keypairs (RSA in this project), sign transactions with a private key, verify signatures with the public key, and test end-to-end transfers between wallets.

By the end, you’ll have a small Go codebase that demonstrates how blocks link together, how PoW controls block creation, and how signed transactions and wallets enable authenticated value transfer—while keeping the implementation intentionally simple and approachable for learning.

