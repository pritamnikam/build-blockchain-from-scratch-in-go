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

## Step-by-step: create the GitHub repo and push this project

### 1) Create the repository on GitHub
1. Go to https://github.com/new
2. **Repository name:** `build-blockchain-from-scratch-in-go`
3. **Description:** (copy/paste the description above)
4. Choose **Public** or **Private**
5. Leave **Initialize this repository with** unchecked (no README / .gitignore / license) since this folder already has them
6. Click **Create repository**

### 2) Initialize git locally and make the first commit (PowerShell)
From this folder:

```powershell
git init
git add .
git commit -m "Initial commit"
git branch -M main
```

### 3) Add the remote and push
On your new GitHub repo page, copy the HTTPS URL, then run:

```powershell
git remote add origin https://github.com/pritamnikam/build-blockchain-from-scratch-in-go.git
git push -u origin main
```

### 4) Next (project creation steps)
- Add `go.mod` with `go mod init <module-path>` (usually `github.com/pritamnikam/build-blockchain-from-scratch-in-go`)
- Implement the first `Block` struct and hashing
- Add a simple `Blockchain` type and persistence strategy (file or in-memory)
- Add Proof-of-Work
- Add transactions (basic UTXO or account-based)

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
A blockchain is a distributed ledger that consists of a chain of blocks linked together. They provide benefits such as decentralization, transparency, and enhanced security through cryptographic algorithms.

A block is a fundamental blockchain unit that contains a single or a set of transactions or digital records. Each block contains a unique hash generated from its data and the previous block’s hash, forming an immutable chain.

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