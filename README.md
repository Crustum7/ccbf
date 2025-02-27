# Brainf*ck interpreter

This program is a Brainf*ck interpreter inspired by the [coding challenge](https://codingchallenges.fyi/challenges/challenge-brainfuck) created by John Crickett. The interpreter creates bytecode based on the Brainf\*ck program and runs the bytecode.

## Usage

Build the project by running:

```bash
go build -o bin/ccbf ./ccbf/cli/interpreter/main.go
```

Then use the interpreter on a `.bf` file like:

```bash
./bin/ccbf examples/mandelbrot.bf
```

Or run directly without building:

```bash
go run ./ccbf/cli/interpreter/main.go examples/mandelbrot.bf
```

## Bytecode instructions

The following is a complete list of the supported operations and their bytecode representation.

| Operation         | Description                                                       | Bytecode | Parameters      |
| ----------------- | ----------------------------------------------------------------- | -------- | --------------- |
| >                 | Move the pointer to the right one step                            | 1        | -               |
| <                 | Move the pointer to the left one step                             | 2        | -               |
| +                 | Increment the memory cell at the pointer                          | 3        | -               |
| -                 | Decrement the memory cell at the pointer                          | 4        | -               |
| .                 | Output the character signified by the cell at the pointer         | 5        | -               |
| ,                 | Input a character and store it in the cell at the pointer         | 6        | -               |
| [                 | Jump past the matching ] if the cell at the pointer is 0          | 7        | Jump location   |
| ]                 | Jump back to the matching [ if the cell at the pointer is nonzero | 8        | Jump location   |
| + (repeated)      | Increase value of memory cell with number of repetitions          | 9        | Repetitions     |
| - (repeated)      | Decrease value of memory cell with number of repetitions          | 10       | Repetitions     |
| > (repeated)      | Move the pointer to the right as many steps as repetitions        | 11       | Repetitions     |
| < (repeated)      | Move the pointer to the left as many steps as repetitions         | 12       | Repetitions     |
| [-]>              | Reset memory cell and move pointer one step right                 | 13       |                 |
| [-]               | Reset memory cell                                                 | 14       |                 |
| [->+<] (repeated) | Move value to the right a number of steps                         | 15       | Number of steps |
