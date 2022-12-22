import * as fs from "fs";
import * as path from "path";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\n");

interface Position {
    x: number;
    y: number;
}

const width: number = input[0].length;
const height: number = input.length;

const getElevation = (x: number, y: number): number => {
    const letter: string = input[y][x];
    switch (letter) {
        case "S":
            return "a".charCodeAt(0);
        case "E":
            return "z".charCodeAt(0);
        default:
            return input[y][x].charCodeAt(0);
    }
};

const findMarker = (marker: string) => {
    for (let x = 0; x < width; x++) {
        for (let y = 0; y < height; y++) {
            if (input[y][x] === marker) {
                return { x, y };
            }
        }
    }

    throw `Can't find ${marker}`;
};

const findMarkers = (marker: string): Position[] => {
    const positions: Position[] = [];
    for (let x = 0; x < width; x++) {
        for (let y = 0; y < height; y++) {
            if (input[y][x] === marker) {
                positions.push({ x, y });
            }
        }
    }

    return positions;
};

const isEqual = (a: Position, b: Position): boolean => {
    return a.x == b.x && a.y == b.y;
};

const isValidMove = (fromElevation: number, to: Position): boolean => {
    if (to.x < 0 || to.x >= width || to.y < 0 || to.y >= height) {
        return false;
    }

    const toElevation = getElevation(to.x, to.y);

    return toElevation - fromElevation <= 1;
};

interface Node {
    parent?: Node;
    position: Position;
    g: number;
    h: number;
    f: number;
}

const getDistanceToEnd = (position: Position): number => {
    return Math.abs(position.x - end.x) + Math.abs(position.y - end.y);
};

let bestPathOfAll: number = Infinity;

const getBestPath = (start: Position, end: Position, memo: Record<string, number>) => {
    const memoKey: string = `${start.x},${start.y}`;
    if (memo[memoKey] !== undefined) {
        return memo[memoKey];
    }

    const open: Node[] = [{ position: start, f: 0, h: getDistanceToEnd(start), g: 0 }];
    const closed: Node[] = [];

    while (open.length > 0) {
        open.sort((a, b) => b.f - a.f);
        const current: Node = open.pop();
        closed.push(current);

        if (isEqual(current.position, end)) {
            let steps: number = 0;
            for (let parent = current.parent; parent !== undefined; parent = parent.parent) {
                steps++;
                if (input[parent.position.y][parent.position.x] === "a") {
                    memo[`${parent.position.x},${parent.position.y}`] = steps;
                }
            }

            memo[memoKey] = steps;
            return steps;
        }

        const fromElevation = getElevation(current.position.x, current.position.y);

        const neighbors: Position[] = [
            { x: current.position.x + 1, y: current.position.y },
            { x: current.position.x - 1, y: current.position.y },
            { x: current.position.x, y: current.position.y + 1 },
            { x: current.position.x, y: current.position.y - 1 },
        ];

        for (const neighbor of neighbors) {
            if (!isValidMove(fromElevation, neighbor)) {
                continue;
            }

            if (input[neighbor.y][neighbor.x] === "a") {
                continue;
            }

            if (closed.find((node) => isEqual(node.position, neighbor))) {
                continue;
            }

            const g: number = current.g + 1;
            const h: number = getDistanceToEnd(neighbor);
            const f: number = g + h;

            if (f > bestPathOfAll) {
                continue;
            }

            const neighborInOpen = open.find((node) => isEqual(node.position, neighbor));

            if (neighborInOpen && neighborInOpen.f < f) {
                continue;
            }

            open.push({ parent: current, position: neighbor, g, h, f });
        }

        closed.push(current);
    }

    return Infinity;
};

const end: Position = findMarker("E");

const starts: Position[] = findMarkers("a");
starts.sort((a, b) => getDistanceToEnd(a) - getDistanceToEnd(b));

const memo: Record<string, number> = {};
for (let i = 0; i < starts.length; i++) {
    console.log(`${i + 1} of ${starts.length}`);

    const start = starts[i];
    const bestPath = getBestPath(start, end, memo);

    bestPathOfAll = Math.min(bestPathOfAll, bestPath);
}

console.log(bestPathOfAll);
