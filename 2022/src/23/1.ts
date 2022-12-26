import * as fs from "fs";
import * as path from "path";
import { createGrid, getPixel, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";
import { getPositionKey, isPositionsEqual, Position } from "../utils/position";

const load = (): Grid => {
    let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    const grid = createGrid();

    for (let y = 0; y < input.length; y++) {
        for (let x = 0; x < input[0].length; x++) {
            setPixel(grid, { x, y }, input[y][x]);
        }
    }

    return grid;
};

const getElfPositions = (grid: Grid): Position[] => {
    const elves: Position[] = [];

    for (let x = grid.min.x; x <= grid.max.x; x++) {
        for (let y = grid.min.y; y <= grid.max.y; y++) {
            if (getPixel(grid, { x, y }) === "#") {
                elves.push({ x, y });
            }
        }
    }

    return elves;
};

type Direction = "North" | "South" | "West" | "East";

const getNextDirection = (direction: Direction): Direction => {
    switch (direction) {
        case "North":
            return "South";
        case "South":
            return "West";
        case "West":
            return "East";
        case "East":
            return "North";
    }
};

const hasElf = (grid: Grid, position: Position): boolean => {
    return getPixel(grid, position) === "#";
};

const canMove = (grid: Grid, position: Position, direction: Direction): boolean => {
    const hasElfNW = hasElf(grid, { x: position.x - 1, y: position.y - 1 });
    const hasElfN = hasElf(grid, { x: position.x, y: position.y - 1 });
    const hasElfNE = hasElf(grid, { x: position.x + 1, y: position.y - 1 });
    const hasElfSW = hasElf(grid, { x: position.x - 1, y: position.y + 1 });
    const hasElfS = hasElf(grid, { x: position.x, y: position.y + 1 });
    const hasElfSE = hasElf(grid, { x: position.x + 1, y: position.y + 1 });
    const hasElfW = hasElf(grid, { x: position.x - 1, y: position.y });
    const hasElfE = hasElf(grid, { x: position.x + 1, y: position.y });

    if (!hasElfNW && !hasElfN && !hasElfNE && !hasElfSW && !hasElfS && !hasElfSE && !hasElfW && !hasElfE) {
        return false;
    }

    switch (direction) {
        case "North":
            return !hasElfNE && !hasElfN && !hasElfNW;

        case "South":
            return !hasElfSE && !hasElfS && !hasElfSW;

        case "West":
            return !hasElfNW && !hasElfW && !hasElfSW;

        case "East":
            return !hasElfNE && !hasElfE && !hasElfSE;
    }
};

const move = (position: Position, direction: Direction): Position => {
    switch (direction) {
        case "North":
            return { x: position.x, y: position.y - 1 };

        case "South":
            return { x: position.x, y: position.y + 1 };

        case "West":
            return { x: position.x - 1, y: position.y };

        case "East":
            return { x: position.x + 1, y: position.y };
    }
};

const getDirectionsToTry = (startDirection: Direction): Direction[] => {
    switch (startDirection) {
        case "North":
            return ["North", "South", "West", "East"];
        case "South":
            return ["South", "West", "East", "North"];
        case "West":
            return ["West", "East", "North", "South"];
        case "East":
            return ["East", "North", "South", "West"];
    }
};

const getProposedPosition = (position: Position, startDirection: Direction): Position => {
    for (const direction of getDirectionsToTry(startDirection)) {
        if (canMove(grid, position, direction)) {
            return move(position, direction);
        }
    }

    return position;
};

const getProposedPositions = (positions: Position[], startDirection: Direction): Position[] => {
    const proposedPositions: Position[] = [];

    for (const position of positions) {
        proposedPositions.push(getProposedPosition(position, startDirection));
    }

    return proposedPositions;
};

const getNewPositions = (positions: Position[], proposedPositions: Position[]): Position[] => {
    const newPositions: Position[] = [...positions];

    for (let i = 0; i < proposedPositions.length; i++) {
        const isUnique = proposedPositions.filter((proposedPosition) => isPositionsEqual(proposedPositions[i], proposedPosition)).length === 1;
        if (isUnique) {
            newPositions[i] = proposedPositions[i];
        }
    }

    return newPositions;
};

const noMoves = (currentPositions: Position[], newPositions: Position[]): boolean => {
    for (let i = 0; i < currentPositions.length; i++) {
        if (!isPositionsEqual(currentPositions[i], newPositions[i])) {
            return false;
        }
    }

    return true;
};

let grid = load();
let currentPositions: Position[] = getElfPositions(grid);
let currentStartDirection: Direction = "North";
let round = 1;

for (let i = 0; i < 10; i++) {
    const proposedPositions: Position[] = getProposedPositions(currentPositions, currentStartDirection);
    const newPositions: Position[] = getNewPositions(currentPositions, proposedPositions);

    for (const currentPosition of currentPositions) {
        setPixel(grid, currentPosition, ".");
    }

    for (const newPosition of newPositions) {
        setPixel(grid, newPosition, "#");
    }

    currentPositions = newPositions;
    currentStartDirection = getNextDirection(currentStartDirection);
    round++;
}

const min: Position = { x: Infinity, y: Infinity };
const max: Position = { x: -Infinity, y: -Infinity };

for (const position of currentPositions) {
    min.x = Math.min(min.x, position.x);
    max.x = Math.max(max.x, position.x);
    min.y = Math.min(min.y, position.y);
    max.y = Math.max(max.y, position.y);
}

let emptyTiles = 0;
for (let y = min.y; y <= max.y; y++) {
    for (let x = min.x; x <= max.x; x++) {
        if (getPixel(grid, { x, y }) !== "#") {
            emptyTiles++;
        }
    }
}

log(emptyTiles);
