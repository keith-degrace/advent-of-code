import * as fs from "fs";
import * as path from "path";
import { Node } from "../utils/astar";
import { createGrid, getPixel, Grid, setPixel } from "../utils/grid";
import { log } from "../utils/log";
import { getManhattanDistance, isPositionsEqual, Position } from "../utils/position";

const load = (): Grid<number> => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    const width = input[0].length;
    const height = input.length;

    const grid = createGrid<number>();

    for (let x = 0; x < width; x++) {
        for (let y = 0; y < height; y++) {
            setPixel(grid, { x, y }, parseInt(input[y][x]));
        }
    }

    return grid;
};

const getNeighbors = (grid: Grid<number>, position: Position): Position[] => {
    const neighbors: Position[] = [];

    const addNeighbor = (neighbor: Position) => {
        if (neighbor.x < grid.min.x || neighbor.x > grid.max.x || neighbor.y < grid.min.y || neighbor.y > grid.max.y) {
            return;
        }

        neighbors.push(neighbor);
    };

    addNeighbor({ x: position.x + 1, y: position.y });
    addNeighbor({ x: position.x, y: position.y + 1 });
    addNeighbor({ x: position.x - 1, y: position.y });
    addNeighbor({ x: position.x, y: position.y - 1 });

    return neighbors;
};

const getLowestRisk = (grid: Grid<number>, start: Position, end: Position): number => {
    const open: Node[] = [{ position: start, f: 0, h: getManhattanDistance(start, end), g: 0 }];
    const closed: Node[] = [];

    while (open.length > 0) {
        open.sort((a, b) => b.f - a.f);
        const current: Node = open.pop();
        closed.push(current);

        if (isPositionsEqual(current.position, end)) {
            return current.g;
        }

        for (const neighbor of getNeighbors(grid, current.position)) {
            if (closed.find((node) => isPositionsEqual(node.position, neighbor))) {
                continue;
            }

            const g: number = current.g + getPixel(grid, neighbor);
            const h: number = getManhattanDistance(neighbor, end);
            const f: number = g + h;

            const neighborInOpen = open.find((node) => isPositionsEqual(node.position, neighbor));
            if (neighborInOpen && neighborInOpen.f <= f) {
                continue;
            }

            open.push({ parent: current, position: neighbor, g, h, f });
        }

        closed.push(current);
    }
};

export const run = () => {
    const grid = load();

    const start = { x: grid.min.x, y: grid.min.y };
    const end = { x: grid.max.x, y: grid.max.y };

    const lowestRisk = getLowestRisk(grid, start, end);

    log(lowestRisk);
};
