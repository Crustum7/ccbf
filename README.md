# Brainf*ck interpreter

## Bytecode

|Operation|Description|Bytecode|Parameters|
|-|-|-|-|
|>|	Move the pointer to the right|0|-|
|<|	Move the pointer to the left|1|-|
|+|	Increment the memory cell at the pointer|2|-|
|-|	Decrement the memory cell at the pointer|3|-|
|.|	Output the character signified by the cell at the pointer|4|-|
|,|	Input a character and store it in the cell at the pointer|5|-|
|[|	Jump past the matching ] if the cell at the pointer is 0|6|Jump location|
|]|	Jump back to the matching [ if the cell at the pointer is nonzero|7|Jump location|
