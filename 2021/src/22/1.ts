import * as fs from "fs";
import * as path from "path";
import { createGrid, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";
import { Position3D } from "../utils/position";

interface Cube {
    min: Position3D;
    max: Position3D;
}

interface Step {
    state: "on" | "off";
    cube: Cube;
}

const load = (): Step[] => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    const steps: Step[] = [];

    let id = 0;
    for (const line of input) {
        const match = line.match(/(.+) x=([0-9]+)..([0-9]+),y=([0-9]+)..([0-9]+),z=([0-9]+)..([0-9]+)/);
        if (!match) {
            continue;
        }

        steps.push({
            state: match[1] as Step["state"],
            cube: {
                min: {
                    x: parseInt(match[2]),
                    y: parseInt(match[4]),
                    z: parseInt(match[6]),
                },
                max: {
                    x: parseInt(match[3]),
                    y: parseInt(match[5]),
                    z: parseInt(match[7]),
                },
            },
        });
    }

    return steps;
};

const getCentroid = (cube: Cube): Position3D => {
    return {
        x: (cube.min.x + cube.max.x) / 2,
        y: (cube.min.y + cube.max.y) / 2,
        z: (cube.min.z + cube.max.z) / 2,
    };
};

const containsPoint = (cube: Cube, point: Position3D): boolean => {
    return point.x >= cube.min.x && point.x <= cube.max.x && point.y >= cube.min.y && point.y <= cube.max.y && point.z >= cube.min.z && point.z <= cube.max.z;
};

const intersects = (cube1: Cube, cube2: Cube): boolean => {
    return (
        cube1.max.x >= cube2.min.x &&
        cube2.max.x >= cube1.min.x &&
        cube1.max.y >= cube2.min.y &&
        cube2.max.y >= cube1.min.y &&
        cube1.max.z >= cube2.min.z &&
        cube2.max.z >= cube1.min.z
    );
};

const empty = (cube: Cube): boolean => {
    return cube.min.x > cube.max.x || cube.min.y > cube.max.y || cube.min.z > cube.max.z;
};

const union = (cube1: Cube, cube2: Cube): Cube[] => {
    const xs = Array.from(new Set([cube1.min.x, cube1.max.x, cube2.min.x, cube2.max.x].sort((a, b) => a - b)));
    const ys = Array.from(new Set([cube1.min.y, cube1.max.y, cube2.min.y, cube2.max.y].sort((a, b) => a - b)));
    const zs = Array.from(new Set([cube1.min.z, cube1.max.z, cube2.min.z, cube2.max.z].sort((a, b) => a - b)));

    log(xs);
    log(ys);
    log(zs);

    let subCubes: Record<string, Cube> = {};

    // Create all the sub-cubes.
    for (let i = 0; i < xs.length - 1; i++) {
        for (let j = 0; j < ys.length - 1; j++) {
            for (let k = 0; k < zs.length - 1; k++) {
                const subCube: Cube = {
                    min: { x: xs[i], y: ys[j], z: zs[k] },
                    max: { x: xs[i + 1], y: ys[j + 1], z: zs[k + 1] },
                };

                // Only keep cubes that are part of the original cubes
                const centroid = getCentroid(subCube);
                if (containsPoint(cube1, centroid) || containsPoint(cube2, centroid)) {
                    subCubes[`${i},${j},${k}`] = subCube;
                }
            }
        }
    }

    // // Eliminate overlapping boundaries.
    // for (let i = 0; i < xs.length - 1; i++) {
    //     for (let j = 0; j < ys.length - 1; j++) {
    //         for (let k = 0; k < zs.length - 1; k++) {
    //             const subCube = subCubes[`${i},${j},${k}`];
    //             if (!subCube) {
    //                 continue;
    //             }

    //             // Avoid overlapping boundaries.  If there's a cube previously next to us, then advance our mins.
    //             if (subCubes[`${i - 1},${j},${k}`] !== undefined) {
    //                 subCube.min.x++;
    //             }

    //             if (subCubes[`${i},${j - 1},${k}`] !== undefined) {
    //                 subCube.min.y++;
    //             }

    //             if (subCubes[`${i},${j},${k - 1}`] !== undefined) {
    //                 subCube.min.z++;
    //             }
    //         }
    //     }
    // }

    return Object.values(subCubes);
};

const drawCube = (grid: Grid, cube: Cube, value: string): void => {
    for (let x = cube.min.x; x <= cube.max.x; x++) {
        for (let y = cube.min.y; y <= cube.max.y; y++) {
            setPixel(grid, { x, y }, value);
        }
    }
};

export const run = () => {
    let steps = load();

    const grid = createGrid();
    setPixel(grid, { x: 0, y: 0 }, ".");
    setPixel(grid, { x: 22, y: 22 }, ".");

    // for (let i = 0; i < steps.length; i++) {
    //     drawCube(grid, steps[i].cube, `${i}`);
    // }

    // printGrid(grid);
    // log("-----------------------------");

    // // while (true) {
    // let foundIntersection = false;

    // for (let i = 0; i < steps.length && !foundIntersection; i++) {
    //     for (let j = i + 1; j < steps.length && !foundIntersection; j++) {
    //         if (intersects(steps[i].cube, steps[j].cube)) {
    //             for (const subCube of union(steps[i].cube, steps[j].cube)) {
    //                 steps.push({ state: steps[i].state, cube: subCube });
    //             }

    //             steps = steps.filter((step) => step !== steps[i] && step !== steps[j]);

    //             foundIntersection = true;
    //             break;
    //         }
    //     }
    // }

    // // if (!foundIntersection) {
    // //     break;
    // // }
    // // }

    const newSteps = [];
    for (const subCube of union(steps[0].cube, steps[1].cube)) {
        newSteps.push({ state: steps[0].state, cube: subCube });
    }
    steps = newSteps;

    // for (let i = 0; i < steps.length; i++) {
    //     drawCube(grid, steps[i].cube, String.fromCharCode(97 + i));
    // }

    drawCube(grid, steps[0].cube, "a");
    drawCube(grid, steps[2].cube, "b");

    printGrid(grid);

    for (let i = 0; i < steps.length; i++) {
        for (let j = i + 1; j < steps.length; j++) {
            if (intersects(steps[i].cube, steps[j].cube)) {
                log(`Intersection:`);
                log(`  ${i}: (${steps[i].cube.min.x},${steps[i].cube.min.y},${steps[i].cube.min.z}) - (${steps[i].cube.max.x},${steps[i].cube.max.y},${steps[i].cube.max.z})`);
                log(`  ${j}: (${steps[j].cube.min.x},${steps[j].cube.min.y},${steps[j].cube.min.z}) - (${steps[j].cube.max.x},${steps[j].cube.max.y},${steps[j].cube.max.z})`);
            }
        }
    }
};
