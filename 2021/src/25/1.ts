import * as fs from "fs";
import * as path from "path";
import { createGrid, getPixel, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";

const load = (): Grid => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    const width = input[0].length;
    const height = input.length;

    const grid = createGrid();

    for (let x = 0; x < width; x++) {
        for (let y = 0; y < height; y++) {
            setPixel(grid, { x, y }, input[y][x]);
        }
    }

    return grid;
};

const advance = (grid: Grid): Grid => {
    const newGrid = createGrid();

    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            setPixel(newGrid, { x, y }, ".");
        }
    }

    // Move horizontal ones.
    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            if (getPixel(grid, { x, y }) === ">") {
                const nextPosition = { x: (x + 1) % grid.width, y };

                if (getPixel(grid, nextPosition) === ".") {
                    setPixel(newGrid, nextPosition, ">");
                } else {
                    setPixel(newGrid, { x, y }, ">");
                }
            }
        }
    }

    // Move vertical ones.
    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            if (getPixel(grid, { x, y }) === "v") {
                const nextPosition = { x, y: (y + 1) % grid.height };

                if (getPixel(grid, nextPosition) !== "v" && getPixel(newGrid, nextPosition) === ".") {
                    setPixel(newGrid, nextPosition, "v");
                } else {
                    setPixel(newGrid, { x, y }, "v");
                }
            }
        }
    }

    let noMovement = true;
    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            if (getPixel(grid, { x, y }) !== getPixel(newGrid, { x, y })) {
                noMovement = false;
                break;
            }
        }
    }

    if (noMovement) {
        return undefined;
    }

    return newGrid;
};

export const run = () => {
    let grid = load();

    for (let step = 1; ; step++) {
        grid = advance(grid);
        if (!grid) {
            log(step);
            break;
        }
    }
};
