import * as fs from "fs";
import * as path from "path";
import { log } from "../utils";

interface Grid {
    data: Set<string>;
    width: number;
    height: number;
    depth: number;
}

const loadCubes = (grid: Grid) => {
    let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\r\n");

    for (const line of input) {
        const coords: number[] = line.split(",").map((coord) => Number.parseInt(coord.trim()));
        addCube(grid, coords[0], coords[1], coords[2]);
    }
};

const addCube = (grid: Grid, x: number, y: number, z: number): void => {
    grid.data.add(`${x},${y},${z}`);
    grid.width = Math.max(grid.width, x + 1);
    grid.height = Math.max(grid.height, y + 1);
    grid.depth = Math.max(grid.depth, z + 1);
};

const hasCube = (grid: Grid, x: number, y: number, z: number): boolean => {
    return grid.data.has(`${x},${y},${z}`);
};

const getCubeSurfaceArea = (grid: Grid, x: number, y: number, z: number): number => {
    let surfaceArea = 0;

    surfaceArea += hasCube(grid, x - 1, y, z) ? 0 : 1;
    surfaceArea += hasCube(grid, x + 1, y, z) ? 0 : 1;
    surfaceArea += hasCube(grid, x, y - 1, z) ? 0 : 1;
    surfaceArea += hasCube(grid, x, y + 1, z) ? 0 : 1;
    surfaceArea += hasCube(grid, x, y, z - 1) ? 0 : 1;
    surfaceArea += hasCube(grid, x, y, z + 1) ? 0 : 1;

    return surfaceArea;
};

const getTotalSurfaceArea = (grid: Grid): number => {
    let surfaceArea = 0;

    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            for (let z = 0; z < grid.depth; z++) {
                if (hasCube(grid, x, y, z)) {
                    surfaceArea += getCubeSurfaceArea(grid, x, y, z);
                }
            }
        }
    }

    return surfaceArea;
};

const grid: Grid = { data: new Set(), width: 0, height: 0, depth: 0 };
loadCubes(grid);

log(getTotalSurfaceArea(grid));
