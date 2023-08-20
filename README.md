# Block-Chain
Block chain is a distributed database that maintains the records of transactions. It can be used to store

> # <U> Block-Chains</u>
>> ## <U> <B> block.go</B></U>
>>> In this file , we have created one block class that holdes multiple elements
>>>> #### Element of Block class
>>>>> 1. <u>timestamp</u> : that help me to store the time record of each block
>>>>> 2. <u>nonce</u> : nonce is helpful to know that that block is correct or not
>>>>> 3. <u>previousHash</u>: previous hash represent the previous block in form of hash
>>>>> 4. <u>transactions</u>: transactions is useful for all the transactions.

>>> **Block**. There is a few methods/functionality using for block.
>>>><u><b>FUNCTIONALITY OF BLOCK</B></u>
>>>>   1. **NewBlock** create a new Block. That helps to create a new block in our block chain.
>>>>   2. **Print**  that helps to print the block in our blockchain
>>>>   3. **Hash** Hash a block using SHA256. this functionality helps to generate a hash for store in new block.
>>>>   4. **MashalJSON** it's a overridden method that helps you for marshal json. also help for print block into json format. It's helps to stringify the block through JSON format then we can able to convert into HASH 
>>>>   5. **Nonce**  : this functionality is used for retrive the value of nonce from particular block.
>>>>   6. **PreviousHash** : this functionality is used for retrive the value of previous hash from particular block.
>>>>   7. **Transactions** : this functionality is used for retrive the values of transactions from particular block.
>>>>   6. **UnmarshalJSON** it's a overridden method that helps you for Unmarshal json. also help for print block into json format. It's helps to stringify the block through JSON format then we can able to convert into HASH 

>> ## <U> <B> chain.go</B></U>
>>> It's a file in which we are just create a classes  
>>>> **Block**. There is a few methods/functionality using for block.
>>>><u><b>FUNCTIONALITY OF BLOCK</B></u>
>>>>   1. **NewBlock** create a new Block. That helps to create a new block in our block chain.
>>>>   2. **Print**  that helps to print the block in our blockchain
>>>>   3. **Hash** Hash a block using SHA256. this functionality helps to generate a hash for store in new block.
>>>>   4. *MashalJSON* it's a overridden method that helps you for unmarshal json. also help for print block into json format. It's helps to stringify the block through JSON format then we can able to convert into HASH 

>>>> **Transactions**. There is a few methods/functionality using for transactions.
>>>><u><b>FUNCTIONALITY OF transactions</B></u>
>>>>   1. **NewTransaction** create a new Transaction for a blockchain. 
>>>>   2. **Print** print a transactions. That helps to print about the transaction
>>>>   3. *MashalJSON* it's a overridden method that helps you for Un-marshal json. It's help to convert into JSON format.It will help us to convert our object into JSON format.
<!-- >>>>   3. **Hash** Hash a block using SHA256. That help to convert the Stringfy JSON into HASH -->

>>>> **Chain**. 
>>>>>It's our Chain(Block-chain) class. that holds 2 things 
>>>>> 1. list of blocks 
>>>>> 2. list of transactions
>>>>> 3. *blockchainAddress* : that is a address one person. it's a block-chain address.
>>>><u><b>FUNCTIONALITY OF  BLOCK-CHAIN</B></u>
>>>>   1. **NewChain** create a new Blockchain. That helps to create a new blockchain.This function helps to add a new block . Also passing the blockchain address as argument. 
>>>>  2.  **LastBlock** that helps to know the last block of our chain. 
>>>>   3. **Print**  that helps to print the entire blockchain with blocks with Block's transactions
>>>>   4. **AddTransaction** that helps to adding the transactions into the blockchain.
>>>> 5. **VerifyTransactionSignature** this is very important functionality for block chain transactions because we need to verify the transaction is done by authentic person
>>>> 6. **CopyTransactionPool** that help to copy the entire transaction of Blockchain
>>>> 7. **ValidPoof** there is concept nonce we need to find the nonce. there is difficulty index that help if use consensus algo and getting zero inital to till difficulty that means that the nonce for next block
>>>> 8. **ProofOfWork** in this function we need to check nonce value is correct or not.
>>>> 9. **Mining** is the method for miner who really find the nonce and transfer money from one to another person
>>>> **CalculateTotalAmount** this function help to know which person have how much amount in the wallet

<!-- >>>>   4. *MashalJSON* it's a overridden method that helps you for unmarshal json. also help for print block into json format. It's helps to stringify the block through JSON format then we can able to convert into HASH  -->


> ## <U> Wallets</u>
>> <U> <B> wallet.go</B></U>
>>> ### Wallet
>>>> **Wallet** this class helps with private key and public key along with blockChainAddress.This is wallet indvidual. One wallet for one person <br/>
>>>><u><b>FUNCTIONALITY OF Wallet</B></u>
>>>>   1. **NewWallet** this is the functionality to create a new Wallet into block-chain
>>>>   1. **NewWallet** this is the functionality to create a new Wallet into block-chain. There is algo to create address of blockchain.
>>>> 2. **PublicKeyStr**  that helps you get publickey of particular wallet into string format
>>>> 3. **Publickey** that helps to get Public key with encrypted data of particular wallet
>>>> 4. **PrivateKey** that helps to get Private key with encrypted data of particular wallet
>>>> 5. **PrivateKeyStr** that helps to get privateKey string of particular wallet
>>>> 6. **BlockchainAddress** : this helps to get blockchain address of particular wallet

>>> ### Transactions
>>>> **Transaction** this use for wallet's transaction in which we required <br/>
>>>>> 1. senderPrivateKey : that helps to miner for verify the sender's authentications <br/>
>>>>> 2.  senderPublicKey : its helps for encryption of transactions into block-chain<br/>
>>>>> 3. senderBlockchainAddress : this is sender block chain address that helps to take money from this block-chain address <br/>
>>>>> 4. recipientBlockchainAddress : this is use for recevier address. in which money will be transfered <br/>
>>>>> 5. value : how much amount transfer from one account to another

>>>> <u><b>FUNCTIONALITY OF Transaction</B></u>
>>>>>   1. **NewTransaction** it's give transaction object with Privatekey, publicKey , sender and recipient with value.
>>>>> 2. **GenerateSignature**  that helps to added transaction and helps to put a signatures of transaction. Encryption using SHA256 and then sign using ECDSA 
>>>>> 3. **MarshalJSON** It's a overidden method.It's helps to convert into json from golang object  and converting into bytes. 
<!-- 
>>> ### Wallet
>>>> **Wallet** this class helps with private key and public key along with blockChainAddress.This is wallet indvidual. One wallet for one person
>>>><u><b>FUNCTIONALITY OF Wallet</B></u>
>>>>   1. **NewWallet** this is the functionality to create a new Wallet into block-chain
>>>>   1. **NewWallet** this is the functionality to create a new Wallet into block-chain. There is algo to create address of blockchain.
>>>> 2. **PublicKeyStr**  that helps you get publickey of particular wallet into string format
>>>> 3. **Publickey** that helps to get Public key with encrypted data of particular wallet
>>>> 4. **PrivateKey** that helps to get Private key with encrypted data of particular wallet
>>>> 5. **PrivateKeyStr** that helps to get privateKey string of particular wallet
>>>> 6. **BlockchainAddress** : this helps to get blockchain address of particular wallet

>>> ### Wallet
>>>> **Wallet** this class helps with private key and public key along with blockChainAddress.This is wallet indvidual. One wallet for one person

>>>><u><b>FUNCTIONALITY OF Wallet</B></u>
>>>>   1. **NewWallet** this is the functionality to create a new Wallet into block-chain
>>>>   1. **NewWallet** this is the functionality to create a new Wallet into block-chain. There is algo to create address of blockchain.
>>>> 2. **PublicKeyStr**  that helps you get publickey of particular wallet into string format
>>>> 3. **Publickey** that helps to get Public key with encrypted data of particular wallet
>>>> 4. **PrivateKey** that helps to get Private key with encrypted data of particular wallet
>>>> 5. **PrivateKeyStr** that helps to get privateKey string of particular wallet
>>>> 6. **BlockchainAddress** : this helps to get blockchain address of particular wallet

>>>>   1. **NewWallet** this is the functionality to create a new Wallet into block-chain. There is algo to create address of blockchain.
>>>> 2. **PublicKeyStr**  that helps you get publickey of particular wallet into string format
>>>> 3. **Publickey** that helps to get Public key with encrypted data of particular wallet
>>>> 4. **PrivateKey** that helps to get Private key with encrypted data of particular wallet
>>>> 5. **PrivateKeyStr** that helps to get privateKey string of particular wallet
>>>> 6. **BlockchainAddress** : this helps to get blockchain address of particular wallet
 -->



> ## <U> Signature</u>
>> <U> <B> signature.go</B></U>
>>>> **Signature** This file use for singnature our wallet. This help to verfiy our signature.  
>>>><u><b>FUNCTIONALITY OF Signature</B></u>
>>>>   1. **String** that function help to return the signature R and S.

<!-- 
>> <u><b>**chain.go** </b></u><br/>
    It's a main structure for chain.that helps to linked one block to another.
>>> <U><b>functionality inside chain</b></u>
>>> 1. **NewChain** that help to create a cnew Chain
>>> 2. **Print** that helps to print entire chain
>>> 3. **CreateBlock** that helps to Create a block inside existing chain
>>> 4. ***LastBlock** that helps to get the last block of chains
>>> 5. ***AddTransactions** that helps to add the transaction into Blockchain

>> <u><b> **transaction.go** </b></u><br/>
    It's a main structure for transaction in blockchain.that help to maintain the transations.
>>> <U><b>functionality inside transactions</b></u>
>>> 1. **NewTransaction** that help to create a new Transaction
>>> 2. **Print** that helps to print entire transaction
>>> 3. ***MarshalJSON** that overrides the transactions
 -->


> # Functionality
In version 3.0 we have Hash value in string instead of sha256 encoding. also added lot's of things like wallet concept is there and Utils pkg there that help to know the signatures public signature and private signate 