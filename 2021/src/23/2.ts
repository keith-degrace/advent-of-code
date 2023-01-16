import * as fs from "fs";
import * as path from "path";
import { getShortestPath } from "../utils/astar";
import { createGrid, getPixel, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";
import { getPositionKey, isPositionsEqual, Position } from "../utils/position";

const load = (): Grid => {
    let input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    input = [input[0], input[1], input[2], "  #D#C#B#A#  ", "  #D#B#A#C#  ", input[3], input[4]];

    const width = input[0].length;
    const height = input.length;

    const grid = createGrid();

    for (let x = 0; x < width; x++) {
        for (let y = 0; y < height; y++) {
            setPixel(grid, { x, y }, input[y][x]);
        }
    }

    return grid;
};

const HallwaySpots: Position[] = [
    { x: 1, y: 1 },
    { x: 2, y: 1 },
    { x: 4, y: 1 },
    { x: 6, y: 1 },
    { x: 8, y: 1 },
    { x: 10, y: 1 },
    { x: 11, y: 1 },
];

const Rooms: Record<string, Position[]> = {
    A: [
        { x: 3, y: 2 },
        { x: 3, y: 3 },
        { x: 3, y: 4 },
        { x: 3, y: 5 },
    ],
    B: [
        { x: 5, y: 2 },
        { x: 5, y: 3 },
        { x: 5, y: 4 },
        { x: 5, y: 5 },
    ],
    C: [
        { x: 7, y: 2 },
        { x: 7, y: 3 },
        { x: 7, y: 4 },
        { x: 7, y: 5 },
    ],
    D: [
        { x: 9, y: 2 },
        { x: 9, y: 3 },
        { x: 9, y: 4 },
        { x: 9, y: 5 },
    ],
};

interface Amphipod {
    type: "A" | "B" | "C" | "D";
    position: Position;
    moveCost: number;
}

const isInHallway = (amphipod: Amphipod): boolean => {
    return HallwaySpots.some((spot) => isPositionsEqual(spot, amphipod.position));
};

const isInHappyPlace = (grid: Grid, amphipod: Amphipod): boolean => {
    const correctRoom = Rooms[amphipod.type];

    // Get the index of this amphipod's position in the room.
    const currentRoomIndex = correctRoom.findIndex((position) => isPositionsEqual(position, amphipod.position));
    if (currentRoomIndex === -1) {
        return false;
    }

    // We're only happy if any spots after us in the room are filled with the correct
    // type, otherwise we are blocking an unhappy amphipod.
    for (let i = currentRoomIndex + 1; i < correctRoom.length; i++) {
        if (getPixel(grid, correctRoom[i]) !== amphipod.type) {
            return false;
        }
    }

    return true;
};

const getNextPositionInHappyRoom = (grid: Grid, amphipod: Amphipod): Position => {
    const correctRoom = Rooms[amphipod.type];

    for (let i = correctRoom.length - 1; i >= 0; i--) {
        const value = getPixel(grid, correctRoom[i]);

        if (value === ".") {
            return correctRoom[i];
        }

        if (value !== amphipod.type) {
            return undefined;
        }
    }

    return undefined;
};

const moveAmphipod = (grid: Grid, amphipod: Amphipod, newPosition: Position): number => {
    const getNeighbors = (position: Position): Position[] => {
        const neighbors: Position[] = [];

        const addNeighbor = (neighbor: Position) => {
            if (getPixel(grid, neighbor) === ".") {
                neighbors.push(neighbor);
            }
        };

        addNeighbor({ x: position.x + 1, y: position.y });
        addNeighbor({ x: position.x, y: position.y + 1 });
        addNeighbor({ x: position.x - 1, y: position.y });
        addNeighbor({ x: position.x, y: position.y - 1 });

        return neighbors;
    };

    const path = getShortestPath(amphipod.position, newPosition, getNeighbors);
    if (path === undefined) {
        return undefined;
    }

    setPixel(grid, amphipod.position, ".");
    amphipod.position = newPosition;
    setPixel(grid, amphipod.position, amphipod.type);

    return (path.length - 1) * amphipod.moveCost;
};

const memo: Record<string, number> = {};

const move = (grid: Grid, amphipods: Amphipod[], level: number): number => {
    const memoKey = `${amphipods.map((amphipod) => getPositionKey(amphipod.position)).join("-")}`;
    if (memo[memoKey] !== undefined) {
        return memo[memoKey];
    }

    let bestCost = Infinity;

    const unhappyAmphipods = amphipods.filter((amphipod) => !isInHappyPlace(grid, amphipod));
    if (unhappyAmphipods.length === 0) {
        return 0;
    }

    for (const amphipod of unhappyAmphipods) {
        const savedPosition = { ...amphipod.position };

        // Try moving to correct room.
        const happyRoomPosition = getNextPositionInHappyRoom(grid, amphipod);
        if (happyRoomPosition) {
            const cost = moveAmphipod(grid, amphipod, happyRoomPosition);
            if (cost !== undefined) {
                bestCost = Math.min(bestCost, move(grid, amphipods, level + 1) + cost);
                moveAmphipod(grid, amphipod, savedPosition);
            }
        }

        // If in a room, try moving to all the different available spots in the hallway.
        if (!isInHallway(amphipod)) {
            for (const hallwaySpot of HallwaySpots) {
                const cost = moveAmphipod(grid, amphipod, hallwaySpot);
                if (cost !== undefined) {
                    // log(`[${amphipod.type}-${amphipod.position.x},${amphipod.position.y}] Moving to hallway`);
                    bestCost = Math.min(bestCost, move(grid, amphipods, level + 1) + cost);
                    moveAmphipod(grid, amphipod, savedPosition);
                }
            }
        }

        moveAmphipod(grid, amphipod, savedPosition);
    }

    memo[memoKey] = bestCost;
    return bestCost;
};

export const run = () => {
    const grid = load();

    const amphipods: Amphipod[] = [];
    for (let x = grid.min.x; x <= grid.max.x; x++) {
        for (let y = grid.min.y; y <= grid.max.y; y++) {
            const value = getPixel(grid, { x, y });
            switch (value) {
                case "A":
                    amphipods.push({ type: "A", position: { x, y }, moveCost: 1 });
                    break;

                case "B":
                    amphipods.push({ type: "B", position: { x, y }, moveCost: 10 });
                    break;

                case "C":
                    amphipods.push({ type: "C", position: { x, y }, moveCost: 100 });
                    break;

                case "D":
                    amphipods.push({ type: "D", position: { x, y }, moveCost: 1000 });
                    break;
            }
        }
    }

    printGrid(grid);

    const cost = move(grid, amphipods, 0);
    log(cost);
};
