import * as fs from "fs";
import * as path from "path";
import { createGrid, getPixel, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";
import { Position } from "../utils/position";

const load = (): Grid => {
    const input: string = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

    const match = input.match(/target area: x=(.+)\.\.(.+), y=(.+)\.\.(.+)/);
    const x1 = parseInt(match[1]);
    const x2 = parseInt(match[2]);
    const y1 = parseInt(match[3]);
    const y2 = parseInt(match[4]);

    const grid = createGrid();

    for (let x = x1; x <= x2; x++) {
        for (let y = y1; y <= y2; y++) {
            setPixel(grid, { x, y }, "T");
        }
    }

    return grid;
};

const shoot = (grid: Grid, vx: number, vy: number): number => {
    const position: Position = { x: 0, y: 0 };

    let highest = 0;

    while (true) {
        position.x += vx;
        position.y += vy;

        if (position.x > grid.max.x || position.y < grid.min.y) {
            return 0;
        }

        highest = Math.max(highest, position.y);

        if (getPixel(grid, position) === "T") {
            return highest;
        }

        vx = Math.max(vx - 1, 0);
        vy = vy - 1;
    }
};

export const run = () => {
    const grid = load();
    setPixel(grid, { x: 0, y: 0 }, "S");

    let highest = 0;

    for (let vx = 0; vx < grid.width; vx++) {
        for (let vy = 0; vy < grid.height; vy++) {
            highest = Math.max(highest, shoot(grid, vx, vy));
        }
    }

    log(highest);
};
