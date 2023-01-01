import * as fs from "fs";
import * as path from "path";
import { getShortestPath } from "../utils/astar";
import { createGrid, findValueInGrid, getPixel, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";
import { getPositionKey, Position } from "../utils/position";
import { startTimer, stopTimer } from "../utils/timer";

const loadMap = (): Grid => {
    let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    let map: Grid = createGrid();

    for (let y = 0; y < input.length; y++) {
        for (let x = 0; x < input[0].length; x++) {
            setPixel(map, { x, y }, input[y][x]);
        }
    }

    return map;
};

interface Key {
    value: string;
    position: Position;
    pathToKeys?: PathToKey[];
}

interface PathToKey {
    key: Key;
    path: Position[];
    requiredKeys: Key[];
}

const getKeys = (map: Grid): Key[] => {
    const keys = [];

    for (let y = 0; y < map.height; y++) {
        for (let x = 0; x < map.width; x++) {
            const value = getPixel(map, { x, y });
            if (/[a-z]/.test(value)) {
                keys.push({
                    value,
                    position: { x, y },
                });
            }
        }
    }

    return keys;
};

const getKey = (keys: Key[], value: string): Key => {
    const key = keys.find((key) => key.value === value.toLowerCase());
    if (!key) {
        throw "Bad key: " + value;
    }

    return key;
};

const getBestPath = (map: Grid, start: Position, end: Position): Position[] => {
    const getNeighbors = (position: Position): Position[] => {
        const neighbors: Position[] = [];

        const addNeighbor = (neighbor: Position) => {
            if (getPixel(map, neighbor) !== "#") {
                neighbors.push(neighbor);
            }
        };

        addNeighbor({ x: position.x + 1, y: position.y });
        addNeighbor({ x: position.x - 1, y: position.y });
        addNeighbor({ x: position.x, y: position.y + 1 });
        addNeighbor({ x: position.x, y: position.y - 1 });

        return neighbors;
    };

    return getShortestPath(start, end, getNeighbors);
};

const blockDeadEnds = (map: Grid, position: Position, visited: Set<string> = new Set()): void => {
    try {
        visited.add(getPositionKey(position));

        const neighbors: Position[] = [
            { x: position.x + 1, y: position.y },
            { x: position.x - 1, y: position.y },
            { x: position.x, y: position.y + 1 },
            { x: position.x, y: position.y - 1 },
        ].filter((neighbor) => !visited.has(getPositionKey(neighbor)));

        const isWall = (position: Position) => {
            return getPixel(map, position) === "#";
        };

        for (const neighbor of neighbors) {
            if (!isWall(neighbor)) {
                blockDeadEnds(map, neighbor, visited);
            }
        }

        if (getPixel(map, position) === "." && neighbors.every(isWall)) {
            setPixel(map, position, "#");
        }
    } finally {
        visited.delete(getPositionKey(position));
    }
};

const getKeysRequiredForPath = (map: Grid, path: Position[], keys: Key[]): Key[] => {
    const positionsToCheck = path.slice(1, path.length - 1);
    if (path.length - positionsToCheck.length !== 2) {
        throw path.length + " " + positionsToCheck.length;
    }

    const requiredKeys: Key[] = [];

    for (const position of path) {
        const value = getPixel(map, position);

        if (/[A-Z]/.test(value)) {
            requiredKeys.push(getKey(keys, value.toLowerCase()));
        }
    }

    return requiredKeys;
};

const calculatePathBetweenKeys = (map: Grid, keys: Key[]): void => {
    for (const key1 of keys) {
        key1.pathToKeys = [];

        for (const key2 of keys) {
            if (key1 === key2) {
                continue;
            }

            const path = getBestPath(map, key1.position, key2.position);
            const requiredKeys = getKeysRequiredForPath(map, path, keys);

            key1.pathToKeys.push({ key: key2, path, requiredKeys });
        }
    }
};

const isBlockedByDoor = (map: Grid, path: Position[], availableKeys: string[]): boolean => {
    for (const position of path) {
        const value = getPixel(map, position);

        if (/[A-Z]/.test(value)) {
            if (!availableKeys.includes(value.toLowerCase())) {
                return true;
            }
        }
    }

    return false;
};

const memo: Record<string, number> = {};

const solve = (map: Grid, key: Key, foundKeys: string[], remainingKeys: string[]): number => {
    let bestSteps = Infinity;

    if (remainingKeys.length === 0) {
        return 0;
    }

    const memoKey = `${key.value}-${foundKeys.join(",")}`;
    if (memo[memoKey] !== undefined) {
        return memo[memoKey];
    }

    for (const pathToKey of key.pathToKeys) {
        if (foundKeys.includes(pathToKey.key.value)) {
            continue;
        }

        if (isBlockedByDoor(map, pathToKey.path, foundKeys)) {
            continue;
        }

        const subFoundKeys: string[] = [...foundKeys, pathToKey.key.value].sort((a, b) => (a < b ? -1 : 1));
        const subRemainingKeys = remainingKeys.filter((key) => !subFoundKeys.includes(key));

        const stepsToKey = pathToKey.path.length - 1;
        const stepsToRest = solve(map, pathToKey.key, subFoundKeys, subRemainingKeys);

        bestSteps = Math.min(bestSteps, stepsToKey + stepsToRest);
    }

    memo[memoKey] = bestSteps;

    return bestSteps;
};

export const run = () => {
    const map = loadMap();
    const keys = getKeys(map);
    const start = findValueInGrid(map, "@");

    blockDeadEnds(map, start);

    calculatePathBetweenKeys(map, keys);

    let best = Infinity;

    log(JSON.stringify(keys.map((key) => key.value)));

    for (const key of keys) {
        const foundKeys: string[] = [key.value];
        const remainingKeys = keys.filter((aKey) => aKey !== key).map((aKey) => aKey.value);

        const pathToKey = getBestPath(map, start, key.position);

        const stepsToKey = pathToKey.length - 1;
        const stepsToRest = solve(map, key, foundKeys, remainingKeys);

        log(`[${key.value}] ${stepsToKey + stepsToRest}`);

        best = Math.min(best, stepsToKey + stepsToRest);
    }

    // 5392 Too High
    log(best);
};
