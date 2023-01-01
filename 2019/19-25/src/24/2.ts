import * as fs from "fs";
import * as path from "path";
import { createGrid, getPixel, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";

interface Space {
    grids: Record<number, Grid>;
    minLevel: number;
    maxLevel: number;
}

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

const hasBug = (space: Space, level: number, x: number, y: number): boolean => {
    const grid = space.grids[level];
    if (!grid) {
        return false;
    }

    return getPixel(grid, { x, y }) === "#";
};

const getAdjacentBugCount = (space: Space, level: number, x: number, y: number): number => {
    let count = 0;

    // Left
    {
        const xx = x - 1;
        const yy = y;

        if (xx == -1) {
            // Outside Left edge
            if (hasBug(space, level - 1, 1, 2)) {
                count++;
            }
        } else if (xx === 2 && yy == 2) {
            // Center
            for (let yyy = 0; yyy < 5; yyy++) {
                if (hasBug(space, level + 1, 4, yyy)) {
                    count++;
                }
            }
        } else if (hasBug(space, level, xx, yy)) {
            // Standard.
            count++;
        }
    }

    // Right
    {
        const xx = x + 1;
        const yy = y;

        if (xx == 5) {
            // Outside Right edge
            if (hasBug(space, level - 1, 3, 2)) {
                count++;
            }
        } else if (xx === 2 && yy == 2) {
            // Center
            for (let yyy = 0; yyy < 5; yyy++) {
                if (hasBug(space, level + 1, 0, yyy)) {
                    count++;
                }
            }
        } else if (hasBug(space, level, xx, yy)) {
            // Standard.
            count++;
        }
    }

    // Top
    {
        const xx = x;
        const yy = y - 1;

        if (yy == -1) {
            // Outside Top edge
            if (hasBug(space, level - 1, 2, 1)) {
                count++;
            }
        } else if (xx === 2 && yy == 2) {
            // Center
            for (let xxx = 0; xxx < 5; xxx++) {
                if (hasBug(space, level + 1, xxx, 4)) {
                    count++;
                }
            }
        } else if (hasBug(space, level, xx, yy)) {
            // Standard.
            count++;
        }
    }

    // Bottom
    {
        const xx = x;
        const yy = y + 1;

        if (yy == 5) {
            // Outside Bottom edge
            if (hasBug(space, level - 1, 2, 3)) {
                count++;
            }
        } else if (xx === 2 && yy == 2) {
            // Center
            for (let xxx = 0; xxx < 5; xxx++) {
                if (hasBug(space, level + 1, xxx, 0)) {
                    count++;
                }
            }
        } else if (hasBug(space, level, xx, yy)) {
            // Standard.
            count++;
        }
    }

    return count;
};

const advance = (space: Space): Space => {
    const newSpace = {
        grids: {},
        minLevel: space.minLevel - 1,
        maxLevel: space.maxLevel + 1,
    };

    for (let level = newSpace.minLevel; level <= newSpace.maxLevel; level++) {
        const newGrid = createGrid();

        for (let y = 0; y < 5; y++) {
            for (let x = 0; x < 5; x++) {
                if (x === 2 && y === 2) {
                    continue;
                }

                const adjacentBugCount = getAdjacentBugCount(space, level, x, y);

                if (hasBug(space, level, x, y)) {
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

        newSpace.grids[level] = newGrid;
    }

    return newSpace;
};

const getBugCount = (space: Space): number => {
    let count = 0;

    for (let level = space.minLevel; level <= space.maxLevel; level++) {
        for (let y = 0; y < 5; y++) {
            for (let x = 0; x < 5; x++) {
                if (hasBug(space, level, x, y)) {
                    count++;
                }
            }
        }
    }

    return count;
};

export const run = () => {
    let space: Space = {
        grids: {
            0: load(),
        },
        minLevel: 0,
        maxLevel: 0,
    };

    for (let minute = 0; minute < 200; minute++) {
        space = advance(space);
    }

    log(getBugCount(space));
};
