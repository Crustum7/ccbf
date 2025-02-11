# Brainf*ck interpreter

## Bytecode instructions

| Operation    | Description                                                       | Bytecode | Parameters    |
| ------------ | ----------------------------------------------------------------- | -------- | ------------- |
| >            | Move the pointer to the right one step                            | 1        | -             |
| <            | Move the pointer to the left one step                             | 2        | -             |
| +            | Increment the memory cell at the pointer                          | 3        | -             |
| -            | Decrement the memory cell at the pointer                          | 4        | -             |
| .            | Output the character signified by the cell at the pointer         | 5        | -             |
| ,            | Input a character and store it in the cell at the pointer         | 6        | -             |
| [            | Jump past the matching ] if the cell at the pointer is 0          | 7        | Jump location |
| ]            | Jump back to the matching [ if the cell at the pointer is nonzero | 8        | Jump location |
| + (repeated) | Increase value of memory cell with number of repetitions          | 9        | Repetitions   |
| - (repeated) | Decrease value of memory cell with number of repetitions          | 10       | Repetitions   |
| > (repeated) | Move the pointer to the right as many steps as repetitions        | 11       | Repetitions   |
| < (repeated) | Move the pointer to the left as many steps as repetitions         | 12       | Repetitions   |
| [-]>         | Reset memory cell and move pointer one step right                 | 13       |               |
| [-]          | Reset memory cell                                                 | 14       |               |
