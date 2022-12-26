import * as fs from "fs";
import * as path from "path";
import { getShortestPath } from "../utils/astar";
import { createGrid, findValueInGrid, getPixel, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";
import { getPositionKey, Position } from "../utils/position";

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
            if (/[a-z0-9]/.test(value)) {
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

const isRobotKey = (key: Key) => {
    return ["1", "2", "3", "4"].includes(key.value);
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

const initRobots = (map: Grid): void => {
    const start = findValueInGrid(map, "@");

    setPixel(map, { x: start.x - 1, y: start.y - 1 }, "1");
    setPixel(map, { x: start.x + 1, y: start.y - 1 }, "2");
    setPixel(map, { x: start.x - 1, y: start.y + 1 }, "3");
    setPixel(map, { x: start.x + 1, y: start.y + 1 }, "4");

    setPixel(map, { x: start.x, y: start.y }, "#");
    setPixel(map, { x: start.x - 1, y: start.y }, "#");
    setPixel(map, { x: start.x + 1, y: start.y }, "#");
    setPixel(map, { x: start.x, y: start.y - 1 }, "#");
    setPixel(map, { x: start.x, y: start.y + 1 }, "#");
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
            if (key1 === key2 || isRobotKey(key2)) {
                continue;
            }

            const path = getBestPath(map, key1.position, key2.position);
            if (path) {
                const requiredKeys = getKeysRequiredForPath(map, path, keys);
                key1.pathToKeys.push({ key: key2, path, requiredKeys });
            }
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

const solve = (map: Grid, robotCurrentKeys: Key[], keyCount: number, foundKeys: string[]): number => {
    let best = Infinity;

    if (foundKeys.length === keyCount) {
        return 0;
    }

    const memoKey = `${robotCurrentKeys.map((key) => key.value).join(",")}-${foundKeys.join(",")}`;
    if (memo[memoKey] !== undefined) {
        return memo[memoKey];
    }

    for (let i = 0; i < robotCurrentKeys.length; i++) {
        for (const pathToKey of robotCurrentKeys[i].pathToKeys) {
            if (foundKeys.includes(pathToKey.key.value)) {
                continue;
            }

            if (isBlockedByDoor(map, pathToKey.path, foundKeys)) {
                continue;
            }

            const newRobotCurrentKeys = [...robotCurrentKeys];
            newRobotCurrentKeys[i] = pathToKey.key;

            const newFoundKeys: string[] = [...foundKeys, pathToKey.key.value].sort((a, b) => (a < b ? -1 : 1));

            const stepsToKey = pathToKey.path.length - 1;
            const stepsToRest = solve(map, newRobotCurrentKeys, keyCount, newFoundKeys);

            best = Math.min(best, stepsToKey + stepsToRest);
        }
    }

    memo[memoKey] = best;

    return best;
};

export const run = () => {
    const map = loadMap();

    initRobots(map);

    const keys = getKeys(map);

    const robotKeys = keys.filter(isRobotKey);
    for (const robot of robotKeys) {
        blockDeadEnds(map, robot.position);
    }

    calculatePathBetweenKeys(map, keys);

    const best = solve(map, robotKeys, keys.length - 4, []);

    log(best);
};
