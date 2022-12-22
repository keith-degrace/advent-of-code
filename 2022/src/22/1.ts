import * as fs from "fs";
import * as path from "path";
import { log } from "../utils";

interface Map {
    pixels: Record<string, string>;
    width: number;
    height: number;
}

type Instruction = "R" | "L" | number;

interface Position {
    x: number;
    y: number;
}

type Orientation = "^" | "<" | "v" | ">";

const loadInput = (): [Map, Instruction[]] => {
    let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    // Load map
    const map = {
        pixels: {},
        width: 0,
        height: input.length - 2,
    };
    for (let y = 0; y < map.height; y++) {
        for (let x = 0; x < input[y].length; x++) {
            setPixel(map, { x, y }, input[y][x]);
        }

        map.width = Math.max(map.width, input[y].length);
    }

    // Load instructions
    const instructions: Instruction[] = [];
    for (const match of input[input.length - 1].matchAll(/([0-9]+|R|L)/g)) {
        if (match[1] === "R" || match[1] === "L") {
            instructions.push(match[1]);
        } else {
            instructions.push(parseInt(match[1]));
        }
    }

    return [map, instructions];
};

const findStart = (map: Map): Position => {
    return { x: getFirstNonEmptyOnRow(map, 0), y: 0 };
};

const getPositionKey = (position: Position): string => {
    return `${position.x},${position.y}`;
};

const setPixel = (map: Map, position: Position, value: string): void => {
    const positionKey = getPositionKey(position);

    if (getPixel(map, position) === "#") {
        console.trace();
    }

    if (value === " ") {
        delete map.pixels[positionKey];
    } else {
        map.pixels[positionKey] = value;
    }
};

const getPixel = (map: Map, position: Position): string => {
    const positionKey = getPositionKey(position);
    return map.pixels[positionKey] ?? " ";
};

const getFirstNonEmptyOnRow = (map: Map, y: number): number => {
    for (let x = 0; x < map.width; x++) {
        if (getPixel(map, { x, y }) !== " ") {
            return x;
        }
    }

    // throw `getFirstPositionOnRow(${position.x}, ${position.y}) failed`;
};

const getLastNonEmptyOnRow = (map: Map, y: number): number => {
    for (let x = map.width; x >= 0; x--) {
        if (getPixel(map, { x, y }) !== " ") {
            return x;
        }
    }

    //  throw `getLastPositionOnRow(${position.x}, ${position.y}) failed`;
};

const getFirstNonEmptyOnColumn = (map: Map, x: number): number => {
    for (let y = 0; y < map.height; y++) {
        if (getPixel(map, { x, y }) !== " ") {
            return y;
        }
    }

    // throw `getFirstNonEmptyOnColumn(${position.x}, ${position.y}) failed`;
};

const getLastNonEmptyOnColumn = (map: Map, x: number): number => {
    for (let y = map.height; y >= 0; y--) {
        if (getPixel(map, { x, y }) !== " ") {
            return y;
        }
    }

    // throw `getFirstNonEmptyOnColumn(${position.x}, ${position.y}) failed`;
};

const printMap = (map: Map): void => {
    for (let y = 0; y < map.height; y++) {
        let line = "";
        for (let x = 0; x < map.width; x++) {
            line += getPixel(map, { x, y });
        }

        log(line);
    }
};

const move = (map: Map, position: Position, orientation: Orientation, steps: number): Position => {
    let newPosition: Position = position;

    for (let step = 0; step < steps; step++) {
        const nextPosition = { ...newPosition };

        switch (orientation) {
            case ">":
                nextPosition.x += 1;
                if (getPixel(map, nextPosition) === " ") {
                    nextPosition.x = getFirstNonEmptyOnRow(map, nextPosition.y);
                }
                break;

            case "v":
                nextPosition.y += 1;
                if (getPixel(map, nextPosition) === " ") {
                    nextPosition.y = getFirstNonEmptyOnColumn(map, nextPosition.x);
                }
                break;

            case "<":
                nextPosition.x -= 1;
                if (getPixel(map, nextPosition) === " ") {
                    nextPosition.x = getLastNonEmptyOnRow(map, nextPosition.y);
                }
                break;

            case "^":
                nextPosition.y -= 1;
                if (getPixel(map, nextPosition) === " ") {
                    nextPosition.y = getLastNonEmptyOnColumn(map, nextPosition.x);
                }
                break;
        }

        if (getPixel(map, nextPosition) === "#") {
            return newPosition;
        }

        setPixel(map, newPosition, orientation);

        newPosition = nextPosition;
    }

    return newPosition;
};

const turn = (orientation: Orientation, direction: "R" | "L"): Orientation => {
    switch (orientation) {
        case ">":
            return direction === "R" ? "v" : "^";
        case "v":
            return direction === "R" ? "<" : ">";
        case "<":
            return direction === "R" ? "^" : "v";
        case "^":
            return direction === "R" ? ">" : "<";
    }
};

const [map, instructions] = loadInput();

let position: Position = findStart(map);
let orientation: Orientation = ">";

for (const instruction of instructions) {
    if (typeof instruction === "number") {
        position = move(map, position, orientation, instruction);
    } else {
        orientation = turn(orientation, instruction);
    }
}

const finalRow = position.y + 1;
const finalColumn = position.x + 1;
const finalFacing = orientation === ">" ? 0 : orientation === "v" ? 1 : orientation === "<" ? 2 : 3;

const password = 1000 * finalRow + 4 * finalColumn + finalFacing;

log(password);

printMap(map);
