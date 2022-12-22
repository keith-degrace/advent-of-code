import * as fs from "fs";
import * as path from "path";
import { log } from "../utils";

interface Face {
    id: number;
    origin: Position;
}

interface Map {
    pixels: Record<string, string>;
    width: number;
    height: number;

    faces: Face[];
    faceSize: number;
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
    const map: Map = {
        pixels: {},
        width: 0,
        height: input.length - 2,
        faces: [],
        faceSize: 50,
    };
    for (let y = 0; y < map.height; y++) {
        for (let x = 0; x < input[y].length; x++) {
            setPixel(map, { x, y }, input[y][x]);
        }

        map.width = Math.max(map.width, input[y].length);
    }

    // Determine the faces.
    const faces: Face[] = [
        { id: 0, origin: { x: map.faceSize, y: 0 } },
        { id: 1, origin: { x: map.faceSize * 2, y: 0 } },
        { id: 2, origin: { x: map.faceSize, y: map.faceSize } },
        { id: 3, origin: { x: 0, y: map.faceSize * 2 } },
        { id: 4, origin: { x: map.faceSize, y: map.faceSize * 2 } },
        { id: 5, origin: { x: 0, y: map.faceSize * 3 } },
    ];

    map.faces = faces;

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

const getFace = (map: Map, position: Position): Face => {
    for (const face of map.faces) {
        if (position.x >= face.origin.x && position.x < face.origin.x + map.faceSize && position.y >= face.origin.y && position.y < face.origin.y + map.faceSize) {
            return face;
        }
    }

    throw `No face at ${position.x},${position.y}.`;
};

const findStart = (map: Map): Position => {
    for (let x = 0; x < map.width; x++) {
        if (getPixel(map, { x, y: 0 }) !== " ") {
            return { x, y: 0 };
        }
    }
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

const printMap = (map: Map): void => {
    for (let y = 0; y < map.height; y++) {
        let line = "";
        for (let x = 0; x < map.width; x++) {
            line += getPixel(map, { x, y });
        }

        log(line);
    }
};

const getNextPosition = (map: Map, position: Position, orientation: Orientation): [Position, Orientation] => {
    const currentFace = getFace(map, position);

    if (
        (orientation == "<" && currentFace.origin.x === position.x) ||
        (orientation == "^" && currentFace.origin.y === position.y) ||
        (orientation == ">" && currentFace.origin.x + map.faceSize - 1 === position.x) ||
        (orientation == "v" && currentFace.origin.y + map.faceSize - 1 === position.y)
    ) {
        const offsetX = position.x - currentFace.origin.x;
        const offsetY = position.y - currentFace.origin.y;

        // Handle wrap cases.
        switch (currentFace.id) {
            case 0:
                switch (orientation) {
                    case "^":
                        return [
                            {
                                x: map.faces[5].origin.x,
                                y: map.faces[5].origin.y + offsetX,
                            },
                            ">",
                        ];
                    case "<":
                        return [
                            {
                                x: map.faces[3].origin.x,
                                y: map.faces[3].origin.y + (map.faceSize - 1) - offsetY,
                            },
                            ">",
                        ];
                }
                break;

            case 1:
                switch (orientation) {
                    case "v":
                        return [
                            {
                                x: map.faces[2].origin.x + (map.faceSize - 1),
                                y: map.faces[2].origin.y + offsetX,
                            },
                            "<",
                        ];
                    case "^":
                        return [
                            {
                                x: map.faces[5].origin.x + offsetX,
                                y: map.faces[5].origin.y + (map.faceSize - 1),
                            },
                            "^",
                        ];
                    case ">":
                        return [
                            {
                                x: map.faces[4].origin.x + (map.faceSize - 1),
                                y: map.faces[4].origin.y + (map.faceSize - 1) - offsetY,
                            },
                            "<",
                        ];
                }
                break;

            case 2:
                switch (orientation) {
                    case "<":
                        return [
                            {
                                x: map.faces[3].origin.x + offsetY,
                                y: map.faces[3].origin.y,
                            },
                            "v",
                        ];
                    case ">":
                        return [
                            {
                                x: map.faces[1].origin.x + offsetY,
                                y: map.faces[1].origin.y + (map.faceSize - 1),
                            },
                            "^",
                        ];
                }
                break;

            case 3:
                switch (orientation) {
                    case "^":
                        return [
                            {
                                x: map.faces[2].origin.x,
                                y: map.faces[2].origin.y + offsetX,
                            },
                            ">",
                        ];
                    case "<":
                        return [
                            {
                                x: map.faces[0].origin.x,
                                y: map.faces[0].origin.y + (map.faceSize - 1) - offsetY,
                            },
                            ">",
                        ];
                }
                break;

            case 4:
                switch (orientation) {
                    case "v":
                        return [
                            {
                                x: map.faces[5].origin.x + (map.faceSize - 1),
                                y: map.faces[5].origin.y + offsetX,
                            },
                            "<",
                        ];
                    case ">":
                        return [
                            {
                                x: map.faces[1].origin.x + (map.faceSize - 1),
                                y: map.faces[1].origin.y + (map.faceSize - 1) - offsetY,
                            },
                            "<",
                        ];
                }
                break;

            case 5:
                switch (orientation) {
                    case "v":
                        return [
                            {
                                x: map.faces[1].origin.x + offsetX,
                                y: map.faces[1].origin.y,
                            },
                            "v",
                        ];
                    case ">":
                        return [
                            {
                                x: map.faces[4].origin.x + offsetY,
                                y: map.faces[4].origin.y + (map.faceSize - 1),
                            },
                            "^",
                        ];
                    case "<":
                        return [
                            {
                                x: map.faces[0].origin.x + offsetY,
                                y: map.faces[0].origin.y,
                            },
                            "v",
                        ];
                }
                break;
        }
    }

    // Handle non-wrap cases.
    switch (orientation) {
        case ">":
            return [{ x: position.x + 1, y: position.y }, orientation];
        case "<":
            return [{ x: position.x - 1, y: position.y }, orientation];
        case "v":
            return [{ x: position.x, y: position.y + 1 }, orientation];
        case "^":
            return [{ x: position.x, y: position.y - 1 }, orientation];
    }
};

const move = (map: Map, position: Position, orientation: Orientation, steps: number): [Position, Orientation] => {
    let newPosition: Position = position;
    let newOrientation: Orientation = orientation;

    for (let step = 0; step < steps; step++) {
        const [nextPosition, nextOrientation] = getNextPosition(map, newPosition, newOrientation);
        if (getPixel(map, nextPosition) === "#") {
            break;
        }

        newPosition = nextPosition;
        newOrientation = nextOrientation;

        setPixel(map, newPosition, orientation);
    }

    return [newPosition, newOrientation];
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
        const [newPosition, newOrientation] = move(map, position, orientation, instruction);
        position = newPosition;
        orientation = newOrientation;
    } else {
        orientation = turn(orientation, instruction);
    }
}

const finalRow = position.y + 1;
const finalColumn = position.x + 1;
const finalFacing = orientation === ">" ? 0 : orientation === "v" ? 1 : orientation === "<" ? 2 : 3;

const password = 1000 * finalRow + 4 * finalColumn + finalFacing;

printMap(map);

log(`Password: ${password}`);
// 101240 too low

// const test = (position: Position, orientation: Orientation): void => {
//     const [newPosition, newOrientation] = getNextPosition(map, position, orientation);

//     setPixel(map, position, orientation);
//     setPixel(map, newPosition, newOrientation);

//     printMap(map);
// };

// Face 0
// test({ x: 6, y: 4 }, "v");
// test({ x: 6, y: 0 }, "^");
// test({ x: 5, y: 3 }, "<");
// test({ x: 9, y: 1 }, ">");

// Face 1
// test({ x: 11, y: 4 }, "v");
// test({ x: 11, y: 0 }, "^");
// test({ x: 10, y: 3 }, "<");
// test({ x: 14, y: 3 }, ">");

// Face 2
// test({ x: 6, y: 9 }, "v");
// test({ x: 6, y: 5 }, "^");
// test({ x: 5, y: 6 }, "<");
// test({ x: 9, y: 6 }, ">");

// Face 3
// test({ x: 1, y: 14 }, "v");
// test({ x: 1, y: 10 }, "^");
// test({ x: 0, y: 11 }, "<");
// test({ x: 4, y: 11 }, ">");

// Face 4
// test({ x: 6, y: 14 }, "v");
// test({ x: 6, y: 10 }, "^");
// test({ x: 5, y: 11 }, "<");
// test({ x: 9, y: 11 }, ">");

// Face 5
// test({ x: 1, y: 19 }, "v");
// test({ x: 1, y: 15 }, "^");
// test({ x: 0, y: 16 }, "<");
// test({ x: 4, y: 16 }, ">");
