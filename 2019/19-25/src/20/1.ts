import * as fs from "fs";
import * as path from "path";
import { getShortestPath } from "../utils/astar";
import { createGrid, getPixel, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";
import { getPositionKey, Position } from "../utils/position";

const load = (): Grid => {
    let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    let grid: Grid = createGrid();

    for (let y = 0; y < input.length; y++) {
        for (let x = 0; x < input[0].length; x++) {
            setPixel(grid, { x, y }, input[y][x]);
        }
    }

    return grid;
};

const isLetter = (value: string) => {
    return /[A-Z]/.test(value);
};

const getPortalPosition = (grid: Grid, firstLetter: Position): Position => {
    // Above first letter
    if (getPixel(grid, { x: firstLetter.x, y: firstLetter.y - 1 }) === ".") {
        return { x: firstLetter.x, y: firstLetter.y - 1 };
    }

    // Below second letter
    if (getPixel(grid, { x: firstLetter.x, y: firstLetter.y + 2 }) === ".") {
        return { x: firstLetter.x, y: firstLetter.y + 2 };
    }

    // Left of first letter
    if (getPixel(grid, { x: firstLetter.x - 1, y: firstLetter.y }) === ".") {
        return { x: firstLetter.x - 1, y: firstLetter.y };
    }

    // Right of first letter
    if (getPixel(grid, { x: firstLetter.x + 2, y: firstLetter.y }) === ".") {
        return { x: firstLetter.x + 2, y: firstLetter.y };
    }
};

const getStart = (grid: Grid): Position => {
    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            const value = getPixel(grid, { x, y });
            if (value !== "A") {
                continue;
            }

            if (getPixel(grid, { x, y: y + 1 }) === "A" || getPixel(grid, { x: x + 1, y }) === "A") {
                return getPortalPosition(grid, { x, y });
            }
        }
    }
};

const getEnd = (grid: Grid): Position => {
    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            const value = getPixel(grid, { x, y });
            if (value !== "Z") {
                continue;
            }

            if (getPixel(grid, { x, y: y + 1 }) === "Z" || getPixel(grid, { x: x + 1, y }) === "Z") {
                return getPortalPosition(grid, { x, y });
            }
        }
    }
};

const getPortals = (grid: Grid): Record<string, Position> => {
    const portalPositions: Record<string, Position[]> = {};

    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            const value = getPixel(grid, { x, y });

            if (isLetter(value)) {
                const below = getPixel(grid, { x, y: y + 1 });

                if (isLetter(below)) {
                    const portalName = `${value}${below}`;
                    if (portalName !== "AA" && portalName !== "ZZ") {
                        const portalPosition = getPortalPosition(grid, { x, y });

                        if (!portalPositions[portalName]) {
                            portalPositions[portalName] = [];
                        }

                        portalPositions[portalName].push(portalPosition);
                    }
                }

                const right = getPixel(grid, { x: x + 1, y });
                if (isLetter(right)) {
                    const portalName = `${value}${right}`;
                    if (portalName !== "AA" && portalName !== "ZZ") {
                        const portalPosition = getPortalPosition(grid, { x, y });

                        if (!portalPositions[portalName]) {
                            portalPositions[portalName] = [];
                        }

                        portalPositions[portalName].push(portalPosition);
                    }
                }
            }
        }
    }

    const portals: Record<string, Position> = {};

    for (const value of Object.values(portalPositions)) {
        portals[getPositionKey(value[0])] = value[1];
        portals[getPositionKey(value[1])] = value[0];
    }

    return portals;
};

export const run = () => {
    const grid = load();
    const start = getStart(grid);
    const end = getEnd(grid);
    const portals: Record<string, Position> = getPortals(grid);

    const getNeighbors = (position: Position): Position[] => {
        const neighbors: Position[] = [];

        const addNeighbor = (neighbor: Position) => {
            if (getPixel(grid, neighbor) === ".") {
                neighbors.push(neighbor);
            }
        };

        addNeighbor({ x: position.x + 1, y: position.y });
        addNeighbor({ x: position.x - 1, y: position.y });
        addNeighbor({ x: position.x, y: position.y + 1 });
        addNeighbor({ x: position.x, y: position.y - 1 });

        const portal = portals[getPositionKey(position)];
        if (portal) {
            neighbors.push(portal);
        }

        return neighbors;
    };

    const path = getShortestPath(start, end, getNeighbors, { getCost: () => 0 });
    for (const position of path) {
        setPixel(grid, position, "+");
    }

    printGrid(grid);

    log(path.length - 1);
};
