import * as fs from "fs";
import * as path from "path";
import { lcm } from "../utils/lcm";
import { log } from "../utils/log";
import { getManhattanDistance, isPositionsEqual, Position } from "../utils/position";
import { startTimer, stopTimer } from "../utils/timer";

type GridValue = "#" | Set<string>;

interface Grid {
    width: number;
    height: number;
    data: GridValue[];
}

export const setPixel = (grid: Grid, position: Position, value: GridValue) => {
    // if (position.x < 0 || position.y < 0 || position.x >= grid.width || position.y >= grid.height) {
    //     throw `Bad setPixel position: ${position.x},${position.y}`;
    // }

    grid.data[position.y * grid.width + position.x] = value;
};

export const getPixel = (grid: Grid, position: Position): GridValue => {
    // if (position.x < 0 || position.y < 0 || position.x >= grid.width || position.y >= grid.height) {
    //     throw `Bad getPixel position: ${position.x},${position.y}`;
    // }

    return grid.data[position.y * grid.width + position.x];
};

const load = (): Grid => {
    let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    const grid = {
        width: input[0].length,
        height: input.length,
        data: [],
    };

    for (let y = 0; y < input.length; y++) {
        for (let x = 0; x < input[0].length; x++) {
            if (input[y][x] === "#") {
                grid.data.push("#");
            } else if (input[y][x] === ".") {
                grid.data.push(new Set());
            } else {
                grid.data.push(new Set([input[y][x]]));
            }
        }
    }

    return grid;
};

const printGrid = (grid: Grid): void => {
    for (let y = 0; y < grid.height; y++) {
        let line = "";

        for (let x = 0; x < grid.width; x++) {
            const values: GridValue = getPixel(grid, { x, y });
            if (values === "#") {
                line += "#";
            } else if (values !== undefined) {
                if (values.size === 0) {
                    line += ".";
                } else if (values.size === 1) {
                    line += Array.from(values)[0];
                } else {
                    line += values.size;
                }
            } else {
                line += "?";
            }
        }

        log(line);
    }
};

const addPixel = (grid: Grid, position: Position, value: string) => {
    let current = getPixel(grid, position);
    if (!current) {
        current = new Set<string>();
        setPixel(grid, position, current);
    }

    if (typeof current === "string") {
        throw `Bad addPixel position: ${position.x},${position.y}`;
    }

    current.add(value);
};

const generateGrids = (initialGrid: Grid): Grid[] => {
    const width = initialGrid.width;
    const height = initialGrid.height;

    const gridStates = lcm(width - 2, height - 2);

    const grids: Grid[] = [initialGrid];
    let previousGrid: Grid = initialGrid;

    for (let time = 1; time < gridStates; time++) {
        const grid: Grid = {
            width,
            height,
            data: [],
        };

        for (let y = 0; y < height; y++) {
            for (let x = 0; x < width; x++) {
                const position: Position = { x, y };
                const value = getPixel(previousGrid, position);
                grid.data.push(value === "#" ? "#" : new Set());
            }
        }

        for (let y = 0; y < height; y++) {
            for (let x = 0; x < width; x++) {
                const position: Position = { x, y };
                const values: GridValue = getPixel(previousGrid, position);
                if (values === "#") {
                    continue;
                }

                for (const value of values) {
                    if (value === ">") {
                        const newPosition: Position = { x: position.x + 1, y: position.y };
                        if (newPosition.x === width - 1) {
                            newPosition.x = 1;
                        }

                        addPixel(grid, newPosition, value);
                    } else if (value === "<") {
                        const newPosition: Position = { x: position.x - 1, y: position.y };
                        if (newPosition.x === 0) {
                            newPosition.x = width - 2;
                        }

                        addPixel(grid, newPosition, value);
                    } else if (value === "v") {
                        const newPosition: Position = { x: position.x, y: position.y + 1 };
                        if (newPosition.y === height - 1) {
                            newPosition.y = 1;
                        }

                        addPixel(grid, newPosition, value);
                    } else if (value === "^") {
                        const newPosition: Position = { x: position.x, y: position.y - 1 };
                        if (newPosition.y === 0) {
                            newPosition.y = height - 2;
                        }

                        addPixel(grid, newPosition, value);
                    }
                }
            }
        }

        grids.push(grid);
        previousGrid = grid;
    }

    return grids;
};

interface TemporalPosition extends Position {
    time: number;
}

const isTemporalPositionsEqual = (a: TemporalPosition, b: TemporalPosition): boolean => {
    if (!a || !b) {
        return !a == !b;
    }

    return a.x == b.x && a.y == b.y && a.time == b.time;
};

interface Node {
    parent?: Node;
    position: TemporalPosition;
    g: number;
    h: number;
    f: number;
}

const getShortestPath = (grids: Grid[], start: TemporalPosition, end: Position): number => {
    const open: Node[] = [{ position: start, f: 0, h: getManhattanDistance(start, end), g: 0 }];
    const closed: Node[] = [];

    const getNeighbors = (position: TemporalPosition): TemporalPosition[] => {
        const neighbors: TemporalPosition[] = [];

        const isBlocked = (position: TemporalPosition): boolean => {
            if (position.x < 0 || position.y < 0 || position.x >= grids[0].width || position.y >= grids[0].height) {
                return true;
            }

            const value = getPixel(grids[position.time % grids.length], position);
            return value === "#" || value.size !== 0;
        };

        const addNeighbor = (neighbor: TemporalPosition) => {
            if (!isBlocked(neighbor)) {
                neighbors.push(neighbor);
            }
        };

        const nextTime = position.time + 1;
        addNeighbor({ x: position.x + 1, y: position.y, time: nextTime });
        addNeighbor({ x: position.x - 1, y: position.y, time: nextTime });
        addNeighbor({ x: position.x, y: position.y + 1, time: nextTime });
        addNeighbor({ x: position.x, y: position.y - 1, time: nextTime });
        addNeighbor({ x: position.x, y: position.y, time: nextTime });

        return neighbors;
    };

    while (open.length > 0) {
        open.sort((a, b) => b.f - a.f);
        const current: Node = open.pop();
        closed.push(current);

        if (isPositionsEqual(current.position, end)) {
            return current.g;
        }

        for (const neighbor of getNeighbors(current.position)) {
            if (closed.find((node) => isTemporalPositionsEqual(node.position, neighbor))) {
                continue;
            }

            const g: number = current.g + 1;
            const h: number = getManhattanDistance(neighbor, end);
            const f: number = g + h;

            const neighborInOpen = open.find((node) => isTemporalPositionsEqual(node.position, neighbor));
            if (neighborInOpen && neighborInOpen.f <= f) {
                continue;
            }

            open.push({ parent: current, position: neighbor, g, h, f });
        }

        closed.push(current);
    }
};

const main = () => {
    const initialGrid = load();

    log(`Generating grids...`);
    const grids = generateGrids(initialGrid);
    log(`${grids.length} unique grids`);

    const start: TemporalPosition = { x: 1, y: 0, time: 0 };
    const end: Position = { x: initialGrid.width - 2, y: initialGrid.height - 1 };

    log(`Finding shortest path...`);
    let steps = getShortestPath(grids, start, end);

    log(steps);
};

startTimer();
main();
stopTimer();
