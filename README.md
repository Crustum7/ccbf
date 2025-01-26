# Brainf*ck interpreter

## Idea for command handling

Every command has its own class that is initialized with number of repetitions. The class has a function returning bytecode produced by the command and number of source tokens have been processed.

## Bytecode

|Operation|Description|Bytecode|Parameters|
|-|-|-|-|
|>|	Move the pointer to the right|1|-|
|<|	Move the pointer to the left|2|-|
|+|	Increment the memory cell at the pointer|3|-|
|-|	Decrement the memory cell at the pointer|4|-|
|.|	Output the character signified by the cell at the pointer|5|-|
|,|	Input a character and store it in the cell at the pointer|6|-|
|[|	Jump past the matching ] if the cell at the pointer is 0|7|Jump location|
|]|	Jump back to the matching [ if the cell at the pointer is nonzero|8|Jump location|
