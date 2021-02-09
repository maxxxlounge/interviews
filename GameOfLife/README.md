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

 * GET /cells return the matrix json and update every 1 second
 * GET /cells.html return the html page for showing situation
 * POST /cells/generate change resource with new width and height, randomize matrix seed

## Deployment strategy

Every cell need to has pointer 8 reference for near cells for each direction, current and next status, a label for inspection
every left cell retrieve the left,topleft,top and top right references from previous cell or row first cell
if present exchanging the itself reference ad bottom cell for his top


NO coordinates needed, because Grid is intend an ordered array oof cell status(bool)
First generation population is done from 2 concurrent function that sets cell status true and false at the "same" time.
The tick function evaluate the next function for all cell than sets for all with a concurrent function that work concurrently from server
Output is json (ordered for definitions) with a sequence of cell status and array and rendered with a simple react application

The UI can generate new seed and define width and height, only not integer parsable input are checked
grid size input is not limited but not more thant 100X100 is suggested
error are managed on log.console

