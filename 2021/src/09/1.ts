import * as fs from "fs";
import * as path from "path";
import { createGrid, getPixel, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";

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

export const run = () => {
    const grid = load();

    let sum = 0;

    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            if (isLowPoint(grid, x, y)) {
                sum += getPixel(grid, { x, y }) + 1;
            }
        }
    }

    log(sum);
};
