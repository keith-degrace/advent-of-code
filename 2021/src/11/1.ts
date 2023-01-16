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

const increaseLevelsByOne = (grid: Grid<number>): void => {
    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            setPixel(grid, { x, y }, getPixel(grid, { x, y }) + 1);
        }
    }
};

const flashOctopuses = (grid: Grid<number>): number => {
    let flashCount = 0;

    const flashNeighbor = (x: number, y: number): void => {
        const value = getPixel(grid, { x, y });
        if (value > 0) {
            setPixel(grid, { x, y }, value + 1);
        }
    };

    while (true) {
        let flashed = false;

        for (let x = 0; x < grid.width; x++) {
            for (let y = 0; y < grid.height; y++) {
                if (getPixel(grid, { x, y }) > 9) {
                    setPixel(grid, { x, y }, 0);

                    flashNeighbor(x - 1, y - 1);
                    flashNeighbor(x, y - 1);
                    flashNeighbor(x + 1, y - 1);
                    flashNeighbor(x + 1, y);
                    flashNeighbor(x - 1, y);
                    flashNeighbor(x - 1, y + 1);
                    flashNeighbor(x, y + 1);
                    flashNeighbor(x + 1, y + 1);

                    flashed = true;
                    flashCount++;
                }
            }
        }

        if (!flashed) {
            break;
        }
    }

    return flashCount;
};

export const run = () => {
    const grid = load();

    let totalFlashes = 0;

    for (let steps = 0; steps < 100; steps++) {
        increaseLevelsByOne(grid);
        totalFlashes += flashOctopuses(grid);
    }

    log(totalFlashes);
};
