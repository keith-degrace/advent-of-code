import * as fs from "fs";
import * as path from "path";
import { log } from "../utils";

type CubeType = "Lava" | "Pocket" | "Displacing";

interface Grid {
    cubes: Map<string, CubeType>;
    width: number;
    height: number;
    depth: number;
}

const loadCubes = (grid: Grid) => {
    let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\r\n");

    for (const line of input) {
        const coords: number[] = line.split(",").map((coord) => Number.parseInt(coord.trim()));
        setCubeType(grid, coords[0], coords[1], coords[2], "Lava");
    }
};

const setCubeType = (grid: Grid, x: number, y: number, z: number, type: CubeType): void => {
    grid.cubes.set(`${x},${y},${z}`, type);
    grid.width = Math.max(grid.width, x + 1);
    grid.height = Math.max(grid.height, y + 1);
    grid.depth = Math.max(grid.depth, z + 1);
};

const getCubeType = (grid: Grid, x: number, y: number, z: number): CubeType => {
    if (x < 0 || x >= grid.width || y < 0 || y >= grid.height || z < 0 || z >= grid.depth) {
        return "Displacing";
    }

    return grid.cubes.get(`${x},${y},${z}`);
};

const isBlocked = (grid: Grid, x: number, y: number, z: number): boolean => {
    const cubeType = getCubeType(grid, x, y, z);
    return cubeType === "Lava" || cubeType === "Pocket";
};

const getCubeSurfaceArea = (grid: Grid, x: number, y: number, z: number): number => {
    let surfaceArea = 0;

    if (isBlocked(grid, x, y, z)) {
        surfaceArea += isBlocked(grid, x - 1, y, z) ? 0 : 1;
        surfaceArea += isBlocked(grid, x + 1, y, z) ? 0 : 1;
        surfaceArea += isBlocked(grid, x, y - 1, z) ? 0 : 1;
        surfaceArea += isBlocked(grid, x, y + 1, z) ? 0 : 1;
        surfaceArea += isBlocked(grid, x, y, z - 1) ? 0 : 1;
        surfaceArea += isBlocked(grid, x, y, z + 1) ? 0 : 1;
    }

    return surfaceArea;
};

const getTotalSurfaceArea = (grid: Grid): number => {
    let surfaceArea = 0;

    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            for (let z = 0; z < grid.depth; z++) {
                surfaceArea += getCubeSurfaceArea(grid, x, y, z);
            }
        }
    }

    return surfaceArea;
};

const setAllCubeTypes = (grid: Grid) => {
    const setDisplacingIfApplicable = (x: number, y: number, z: number): void => {
        const cubeType = getCubeType(grid, x, y, z);
        if (cubeType !== undefined) {
            return;
        }

        setCubeType(grid, x, y, z, "Displacing");

        setDisplacingIfApplicable(x - 1, y, z);
        setDisplacingIfApplicable(x + 1, y, z);
        setDisplacingIfApplicable(x, y - 1, z);
        setDisplacingIfApplicable(x, y + 1, z);
        setDisplacingIfApplicable(x, y, z - 1);
        setDisplacingIfApplicable(x, y, z + 1);
    };

    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            for (let z = 0; z < grid.depth; z++) {
                // Only interested in doing the exterior cubes.
                if (x !== 0 && x !== grid.width - 1 && y !== 0 && x !== grid.height - 1 && z !== 0 && x !== grid.depth - 1) {
                    continue;
                }

                setDisplacingIfApplicable(x, y, z);
            }
        }
    }

    // Anything left is a pocket.
    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            for (let z = 0; z < grid.depth; z++) {
                const cubeType = getCubeType(grid, x, y, z);
                if (cubeType === undefined) {
                    setCubeType(grid, x, y, z, "Pocket");
                }
            }
        }
    }
};

const grid: Grid = { cubes: new Map(), width: 0, height: 0, depth: 0 };
loadCubes(grid);
setAllCubeTypes(grid);

log(getTotalSurfaceArea(grid));
