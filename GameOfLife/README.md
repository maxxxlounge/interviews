# Conway's Game of Life.

The universe of the Game of Life is an infinite, two-dimensional orthogonal grid of square cells,
each of which is in one of two possible states, alive or dead, (or populated and unpopulated, respectively).
Every cell interacts with its eight neighbours, which are the cells that are horizontally, vertically, or diagonally adjacent.

At each step in time, the following transitions occur:

* Any live cell with two or three neighbors survives.
* Any dead cell with three live neighbors becomes a live cell.
* All other live cells die in the next generation.

Similarly, all other dead cells stay dead.

The initial pattern constitutes the seed of the system.

The first generation is created by applying the above rules simultaneously to every cell in the seed; births and deaths occur simultaneously,
and the discrete moment at which this happens is sometimes called a tick.
Each generation is a pure function of the preceding one.

The rules continue to be applied repeatedly to create further generations.


## Execution

start on port 8888

```sh
    ./run.sh
```

## Endpoint

 * GET /cells return the matrix and update every 1 second
 * POST /cells/generate change resource with width and height, randomize matrix

## Deployment strategy

Server rendering: go
Every cell need to has pointer reference for near cells, current state and a function that get the next state

output can be a json (ordered for definitions) with a sequence of cell status and array and rendered with a simple react application








