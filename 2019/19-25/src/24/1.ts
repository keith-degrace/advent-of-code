import * as fs from "fs";
import * as path from "path";
import { createGrid, getPixel, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";

const load = (): Grid => {
    let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    let map: Grid = createGrid();

    for (let y = 0; y < input.length; y++) {
        for (let x = 0; x < input[0].length; x++) {
            setPixel(map, { x, y }, input[y][x]);
        }
    }

    return map;
};

const serializeGrid = (grid: Grid): string => {
    let serialized = "";

    for (const entry of Object.entries(grid.pixels)) {
        serialized += `${entry[1]}`;
    }

    if (serialized.length != 25) {
        throw serialized.length;
    }

    return serialized;
};

const hasBug = (grid: Grid, x: number, y: number) => {
    return getPixel(grid, { x, y }) === "#";
};

const advance = (grid: Grid): Grid => {
    const newGrid = createGrid();

    for (let y = grid.min.y; y <= grid.max.y; y++) {
        for (let x = grid.min.x; x <= grid.max.x; x++) {
            let adjacentBugCount = 0;
            if (hasBug(grid, x - 1, y)) {
                adjacentBugCount++;
            }
            if (hasBug(grid, x + 1, y)) {
                adjacentBugCount++;
            }
            if (hasBug(grid, x, y - 1)) {
                adjacentBugCount++;
            }
            if (hasBug(grid, x, y + 1)) {
                adjacentBugCount++;
            }

            if (hasBug(grid, x, y)) {
                if (adjacentBugCount === 1) {
                    setPixel(newGrid, { x, y }, "#");
                } else {
                    setPixel(newGrid, { x, y }, ".");
                }
            } else {
                if (adjacentBugCount === 1 || adjacentBugCount === 2) {
                    setPixel(newGrid, { x, y }, "#");
                } else {
                    setPixel(newGrid, { x, y }, ".");
                }
            }
        }
    }

    return newGrid;
};

const getBiodiversityRating = (grid: Grid): number => {
    let exponent = 0;
    let rating = 0;

    for (let y = grid.min.y; y <= grid.max.y; y++) {
        for (let x = grid.min.x; x <= grid.max.x; x++) {
            if (hasBug(grid, x, y)) {
                rating += Math.pow(2, exponent);
            }

            exponent++;
        }
    }

    return rating;
};

export const run = () => {
    let grid = load();

    const seenGrids = new Set<string>();
    seenGrids.add(serializeGrid(grid));

    for (let minute = 1; ; minute++) {
        grid = advance(grid);

        const serializedGrid = serializeGrid(grid);
        if (seenGrids.has(serializedGrid)) {
            log(getBiodiversityRating(grid));
            return;
        }

        seenGrids.add(serializedGrid);
    }
};
