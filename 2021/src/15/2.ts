import * as fs from "fs";
import * as path from "path";
import { createGrid, getPixel, Grid, setPixel } from "../utils/grid";
import { log } from "../utils/log";
import { getPositionKey, isPositionsEqual, Position } from "../utils/position";

const load = (): Grid<number> => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    const baseWidth = input[0].length;
    const baseHeight = input.length;

    const baseGrid = createGrid<number>();
    for (let x = 0; x < baseWidth; x++) {
        for (let y = 0; y < baseHeight; y++) {
            setPixel(baseGrid, { x, y }, parseInt(input[y][x]));
        }
    }

    const scale = 5;
    const scaledWidth = baseWidth * scale;
    const scaledHeight = baseHeight * scale;

    const grid = createGrid<number>();
    for (let x = 0; x < scaledWidth; x++) {
        for (let y = 0; y < scaledHeight; y++) {
            const baseValue = getPixel(baseGrid, { x: x % baseWidth, y: y % baseHeight });
            let scaledValue = baseValue + Math.floor(x / baseWidth) + Math.floor(y / baseHeight);
            if (scaledValue > 9) {
                scaledValue -= 9;
            }

            setPixel(grid, { x, y }, scaledValue);
        }
    }

    return grid;
};

interface Node {
    position: Position;
    cost: number;
}

const getLowestRisk = (grid: Grid<number>, start: Position, end: Position): number => {
    const open: Node[] = [{ position: start, cost: 0 }];
    const closed: string[] = [];

    while (open.length > 0) {
        open.sort((a, b) => b.cost - a.cost);

        const current: Node = open.pop();
        if (isPositionsEqual(current.position, end)) {
            return current.cost;
        }

        closed.push(getPositionKey(current.position));

        const neighbors = [
            { x: current.position.x + 1, y: current.position.y },
            { x: current.position.x - 1, y: current.position.y },
            { x: current.position.x, y: current.position.y + 1 },
            { x: current.position.x, y: current.position.y - 1 },
        ];

        for (const neighbor of neighbors) {
            if (neighbor.x < 0 || neighbor.x >= grid.width || neighbor.y < 0 || neighbor.y >= grid.width) {
                continue;
            }

            if (closed.includes(getPositionKey(neighbor))) {
                continue;
            }

            const cost: number = current.cost + getPixel(grid, neighbor);

            const neighborInOpen = open.find((node) => isPositionsEqual(node.position, neighbor));
            if (neighborInOpen && neighborInOpen.cost <= cost) {
                continue;
            }

            open.push({ position: neighbor, cost });
        }
    }
};

export const run = () => {
    const grid = load();

    const start = { x: grid.min.x, y: grid.min.y };
    const end = { x: grid.max.x, y: grid.max.y };

    log("Takes about 10 minutes...");

    const lowestRisk = getLowestRisk(grid, start, end);

    log(lowestRisk);
};
