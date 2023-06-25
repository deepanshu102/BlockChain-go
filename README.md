# Block-Chain
Block chain is a distributed database that maintains the records of transactions. It can be used to store

> ## <u>Models </u>
>> <u><b>**block.go**</b></u> <br/>
    It's a file  in which we are just create a class with Name **Block**.There is few methods using for Block. 
>>> <u><b>functionality of Block</b></u>
>>>   1. **NewBlock** create a new Block
>>>   2. **Print** print a block
>>>   3. **Hash** Hash a block using SHA256
>>>   4. *MashalJSON* it's a overridden method

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



> # Functionality
In version 1.0 we have Hash value in string instead of sha256 encoding.