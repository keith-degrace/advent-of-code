import * as fs from "fs";
import * as path from "path";
import { createGrid, getPixel, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";
import { getManhattanDistance, getPositionKey, isPositionsEqual, Position } from "../utils/position";

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

type PortalKind = "Inner" | "Outer";

interface Portal {
    name: string;
    position: Position;
    kind: PortalKind;
    link: Portal;
}

const getPortals = (grid: Grid): Record<string, Portal> => {
    const portals: Portal[] = [];

    log("HACK: inner/outer boundaries are hardcoded");
    // Sample
    // const outerMin: Position = { x: 2, y: 2 };
    // const outerMax: Position = { x: 42, y: 34 };
    // const innerMin: Position = { x: 8, y: 8 };
    // const innerMax: Position = { x: 36, y: 28 };

    // Real Input
    const outerMin: Position = { x: 2, y: 2 };
    const outerMax: Position = { x: 102, y: 108 };
    const innerMin: Position = { x: 26, y: 26 };
    const innerMax: Position = { x: 78, y: 84 };

    const getPortalKind = (position: Position): PortalKind => {
        if (position.x === outerMin.x || position.x === outerMax.x || position.y === outerMin.y || position.y === outerMax.y) {
            return "Outer";
        } else if (position.x === innerMin.x || position.x === innerMax.x || position.y === innerMin.y || position.y === innerMax.y) {
            return "Inner";
        } else {
            setPixel(grid, position, "@");
            printGrid(grid);
            throw JSON.stringify(position);
        }
    };

    for (let x = 0; x < grid.width; x++) {
        for (let y = 0; y < grid.height; y++) {
            const value = getPixel(grid, { x, y });

            if (isLetter(value)) {
                const below = getPixel(grid, { x, y: y + 1 });
                if (isLetter(below)) {
                    const name = `${value}${below}`;
                    const position = getPortalPosition(grid, { x, y });
                    const kind = getPortalKind(position);
                    portals.push({ name, position, kind, link: undefined });
                }

                const right = getPixel(grid, { x: x + 1, y });
                if (isLetter(right)) {
                    const name = `${value}${right}`;
                    const position = getPortalPosition(grid, { x, y });
                    const kind = getPortalKind(position);
                    portals.push({ name, position, kind, link: undefined });
                }
            }
        }
    }

    for (const portal of portals) {
        portal.link = portals.find((candidate) => candidate != portal && candidate.name === portal.name);
    }

    const portalMap = {};
    for (const portal of portals) {
        portalMap[getPositionKey(portal.position)] = portal;
    }

    return portalMap;
};

interface Position3D extends Position {
    level: number;
}

export interface Node {
    parent?: Node;
    position: Position3D;
    g: number;
    h: number;
    f: number;
}

export const isPositionsEqual3D = (a: Position3D, b: Position3D): boolean => {
    if (!a || !b) {
        return !a === !b;
    }

    return a.x === b.x && a.y === b.y && a.level === b.level;
};

export const getShortestPath = (start: Position3D, end: Position3D, getNeighbors: (position: Position3D) => Position3D[]): Position[] => {
    const open: Node[] = [{ position: start, f: 0, h: getManhattanDistance(start, end), g: 0 }];
    const closed: Node[] = [];

    while (open.length > 0) {
        open.sort((a, b) => b.f - a.f);
        const current: Node = open.pop();
        closed.push(current);

        if (isPositionsEqual3D(current.position, end)) {
            return buildPath(current);
        }

        for (const neighbor of getNeighbors(current.position)) {
            if (closed.find((node) => isPositionsEqual3D(node.position, neighbor))) {
                continue;
            }

            const g: number = current.g + 1;
            const h: number = getManhattanDistance(neighbor, end) + neighbor.level;
            const f: number = g + h;

            const neighborInOpen = open.find((node) => isPositionsEqual3D(node.position, neighbor));
            if (neighborInOpen && neighborInOpen.f <= f) {
                continue;
            }

            open.push({ parent: current, position: neighbor, g, h, f });
        }

        closed.push(current);
    }
};

const buildPath = (end: Node): Position[] => {
    const path: Position[] = [];

    for (let current = end; current !== undefined; current = current.parent) {
        path.push(current.position);
    }

    return path.reverse();
};

export const run = () => {
    const grid = load();
    const portals = getPortals(grid);

    const start = { ...Object.values(portals).find((portal) => portal.name === "AA").position, level: 0 };
    delete portals[getPositionKey(start)];

    const end = { ...Object.values(portals).find((portal) => portal.name === "ZZ").position, level: 0 };
    delete portals[getPositionKey(end)];

    const getNeighbors = (position: Position3D): Position3D[] => {
        const neighbors: Position3D[] = [];

        const addNeighbor = (neighbor: Position3D) => {
            if (neighbor.level > 25) {
                return;
            }

            if (getPixel(grid, neighbor) === ".") {
                neighbors.push(neighbor);
            }
        };

        addNeighbor({ x: position.x + 1, y: position.y, level: position.level });
        addNeighbor({ x: position.x - 1, y: position.y, level: position.level });
        addNeighbor({ x: position.x, y: position.y + 1, level: position.level });
        addNeighbor({ x: position.x, y: position.y - 1, level: position.level });

        const portal = portals[getPositionKey(position)];
        if (portal) {
            if (portal.kind === "Inner") {
                addNeighbor({ ...portal.link.position, level: position.level + 1 });
            } else if (position.level !== 0) {
                addNeighbor({ ...portal.link.position, level: position.level - 1 });
            }
        }

        return neighbors;
    };

    const path = getShortestPath(start, end, getNeighbors);
    log(path?.length - 1);
};
