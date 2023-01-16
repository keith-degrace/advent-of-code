import * as fs from "fs";
import * as path from "path";
import { createGrid, getPixel, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";
import { isPositionsEqual, Position } from "../utils/position";

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

const isLowPoint = (grid: Grid<number>, x: number, y: number): boolean => {
    const value = getPixel(grid, { x, y });

    const left = getPixel(grid, { x: x - 1, y });
    if (left <= value) {
        return false;
    }

    const right = getPixel(grid, { x: x + 1, y });
    if (right <= value) {
        return false;
    }

    const top = getPixel(grid, { x, y: y - 1 });
    if (top <= value) {
        return false;
    }

    const bottom = getPixel(grid, { x, y: y + 1 });
    if (bottom <= value) {
        return false;
    }

    return true;
};

const isInExistingBasin = (basins: Position[][], x: number, y: number): boolean => {
    for (const basin of basins) {
        for (const position of basin) {
            if (position.x === x && position.y === y) {
                return true;
            }
        }
    }

    return false;
};

const getBasin = (grid: Grid<number>, x: number, y: number): Position[] => {
    const basin: Position[] = [];

    const search = (position: Position) => {
        if (!!basin.find((value) => isPositionsEqual(value, position))) {
            return;
        }

        const value = getPixel(grid, position);
        if (value === undefined || value === 9) {
            return;
        }

        basin.push(position);

        search({ x: position.x - 1, y: position.y });
        search({ x: position.x + 1, y: position.y });
        search({ x: position.x, y: position.y - 1 });
        search({ x: position.x, y: position.y + 1 });
    };

    search({ x, y });

    return basin;
};

export const run = () => {
    const grid = load();

    let basins: Position[][] = [];

    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            if (isLowPoint(grid, x, y) && !isInExistingBasin(basins, x, y)) {
                basins.push(getBasin(grid, x, y));
            }
        }
    }

    basins.sort((a, b) => b.length - a.length);

    log(basins[0].length * basins[1].length * basins[2].length);
};
