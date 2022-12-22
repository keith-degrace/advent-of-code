import * as fs from "fs";
import * as path from "path";
import { getPositionKey, Position } from "../utils/position";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\n");

const grid: Record<string, string> = {};
let minGrid: Position = { x: 500, y: 0 };
let maxGrid: Position = { x: 500, y: 0 };

let floor: number = 0;

const getGridItem = (position: Position): string => {
    if (position.y === floor) {
        return "-";
    }

    return grid[getPositionKey(position)];
};

const setGridItem = (position: Position, type: string) => {
    grid[getPositionKey(position)] = type;

    minGrid.x = Math.min(minGrid.x, position.x);
    minGrid.y = Math.min(minGrid.y, position.y);

    maxGrid.x = Math.max(maxGrid.x, position.x);
    maxGrid.y = Math.max(maxGrid.y, position.y);
};

const initializeGrid = () => {
    for (const rockPath of input) {
        const points: Position[] = rockPath.split(" -> ").map((point) => {
            const parts = point.split(",");
            return { x: Number.parseInt(parts[0]), y: Number.parseInt(parts[1]) };
        });

        for (let i = 1; i < points.length; i++) {
            const start: Position = points[i - 1];
            const end: Position = points[i];

            const stepsX = start.x < end.x ? 1 : start.x > end.x ? -1 : 0;
            const stepsY = start.y < end.y ? 1 : start.y > end.y ? -1 : 0;

            const current: Position = { x: start.x, y: start.y };
            while (true) {
                setGridItem(current, "#");

                if (current.x === end.x && current.y === end.y) {
                    break;
                }

                current.x += stepsX;
                current.y += stepsY;
            }
        }
    }

    floor = maxGrid.y + 2;
};

const addUnitOfSand = (): Position => {
    const sand: Position = { x: 500, y: 0 };

    while (true) {
        if (getGridItem({ x: sand.x, y: sand.y + 1 }) === undefined) {
            sand.y++;
        } else if (getGridItem({ x: sand.x - 1, y: sand.y + 1 }) === undefined) {
            sand.x--;
            sand.y++;
        } else if (getGridItem({ x: sand.x + 1, y: sand.y + 1 }) === undefined) {
            sand.x++;
            sand.y++;
        } else {
            // Resting.
            break;
        }
    }

    return sand;
};

initializeGrid();

let units = 0;

while (true) {
    const restingPosition: Position = addUnitOfSand();
    setGridItem(restingPosition, "o");
    units++;

    if (restingPosition.x === 500 && restingPosition.y === 0) {
        break;
    }
}

console.log(units);
