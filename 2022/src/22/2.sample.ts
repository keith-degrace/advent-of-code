import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";
import { getPositionKey, Position } from "../utils/position";

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

type Orientation = "^" | "<" | "v" | ">";

const loadInput = (): [Map, Instruction[]] => {
    let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    // Load map
    const map: Map = {
        pixels: {},
        width: 0,
        height: input.length - 2,
        faces: [],
        faceSize: 4,
    };
    for (let y = 0; y < map.height; y++) {
        for (let x = 0; x < input[y].length; x++) {
            setPixel(map, { x, y }, input[y][x]);
        }

        map.width = Math.max(map.width, input[y].length);
    }

    // Determine the faces.
    const faces: Face[] = [
        { id: 0, origin: { x: 8, y: 0 } },
        { id: 1, origin: { x: 0, y: 4 } },
        { id: 2, origin: { x: 4, y: 4 } },
        { id: 3, origin: { x: 8, y: 4 } },
        { id: 4, origin: { x: 8, y: 8 } },
        { id: 5, origin: { x: 12, y: 8 } },
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

const getNextPosition = (position: Position, orientation: Orientation): [Position, Orientation] => {
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
                                x: map.faces[1].origin.x + (map.faceSize - 1) - offsetX,
                                y: map.faces[1].origin.y,
                            },
                            "v",
                        ];
                    case "<":
                        return [
                            {
                                x: map.faces[2].origin.x + offsetY,
                                y: map.faces[2].origin.y,
                            },
                            "v",
                        ];
                    case ">":
                        return [
                            {
                                x: map.faces[5].origin.x + (map.faceSize - 1),
                                y: map.faces[5].origin.y + (map.faceSize - 1) - offsetY,
                            },
                            "<",
                        ];
                }
                break;

            case 1:
                switch (orientation) {
                    case "v":
                        return [
                            {
                                x: map.faces[4].origin.x + (map.faceSize - 1) - offsetX,
                                y: map.faces[4].origin.y + (map.faceSize - 1),
                            },
                            "^",
                        ];
                    case "^":
                        return [
                            {
                                x: map.faces[0].origin.x + (map.faceSize - 1) - offsetX,
                                y: map.faces[0].origin.y,
                            },
                            "v",
                        ];
                    case "<":
                        return [
                            {
                                x: map.faces[5].origin.x + (map.faceSize - 1) - offsetY,
                                y: map.faces[5].origin.y + (map.faceSize - 1),
                            },
                            "^",
                        ];
                }
                break;

            case 2:
                switch (orientation) {
                    case "v":
                        return [
                            {
                                x: map.faces[4].origin.x,
                                y: map.faces[4].origin.y + (map.faceSize - 1) - offsetX,
                            },
                            ">",
                        ];
                    case "^":
                        return [
                            {
                                x: map.faces[0].origin.x,
                                y: map.faces[0].origin.y + offsetX,
                            },
                            ">",
                        ];
                }
                break;

            case 3:
                switch (orientation) {
                    case ">":
                        return [
                            {
                                x: map.faces[5].origin.x + (map.faceSize - 1) - offsetY,
                                y: map.faces[5].origin.y,
                            },
                            "v",
                        ];
                }
                break;

            case 4:
                switch (orientation) {
                    case "v":
                        return [
                            {
                                x: map.faces[1].origin.x + (map.faceSize - 1) - offsetX,
                                y: map.faces[1].origin.y + (map.faceSize - 1),
                            },
                            "^",
                        ];
                    case "<":
                        return [
                            {
                                x: map.faces[2].origin.x + (map.faceSize - 1) - offsetY,
                                y: map.faces[2].origin.y + (map.faceSize - 1),
                            },
                            "^",
                        ];
                }
                break;

            case 5:
                switch (orientation) {
                    case "v":
                        return [
                            {
                                x: map.faces[1].origin.x,
                                y: map.faces[1].origin.y + (map.faceSize - 1) - offsetX,
                            },
                            ">",
                        ];
                    case "^":
                        return [
                            {
                                x: map.faces[3].origin.x + (map.faceSize - 1),
                                y: map.faces[3].origin.y + (map.faceSize - 1) - offsetX,
                            },
                            "<",
                        ];
                    case ">":
                        return [
                            {
                                x: map.faces[0].origin.x + (map.faceSize - 1),
                                y: map.faces[0].origin.y + (map.faceSize - 1) - offsetY,
                            },
                            "<",
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

    return [position, orientation];
};

const move = (map: Map, position: Position, orientation: Orientation, steps: number): [Position, Orientation] => {
    let newPosition: Position = position;
    let newOrientation: Orientation = orientation;

    for (let step = 0; step < steps; step++) {
        const [nextPosition, nextOrientation] = getNextPosition(newPosition, newOrientation);
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

// for (let i = 0; i < map.faces.length; i++) {
//     for (let x = 0; x < map.faceSize; x++) {
//         for (let y = 0; y < map.faceSize; y++) {
//             map.pixels[getPositionKey({ x: map.faces[i].origin.x + x, y: map.faces[i].origin.y + y })] = `${i}`;
//         }
//     }
// }

// const test = (position: Position, orientation: Orientation): void => {
//     const [newPosition, newOrientation] = move3D(map, position, orientation);

//     setPixel(map, position, orientation);
//     setPixel(map, newPosition, newOrientation);

//     printMap(map);
// };

// Face 0
// test({ x: 9, y: 3 }, "v");
// test({ x: 10, y: 0 }, "^");
// test({ x: 8, y: 3 }, "<");
// test({ x: 11, y: 1 }, ">");

// Face 1
// test({ x: 1, y: 7 }, "v");
// test({ x: 1, y: 4 }, "^");
// test({ x: 0, y: 5 }, "<");
// test({ x: 3, y: 5 }, ">");

// Face 2
// test({ x: 5, y: 7 }, "v");
// test({ x: 5, y: 4 }, "^");
// test({ x: 4, y: 5 }, "<");
// test({ x: 7, y: 5 }, ">");

// Face 3
// test({ x: 9, y: 7 }, "v");
// test({ x: 10, y: 4 }, "^");
// test({ x: 8, y: 4 }, "<");
// test({ x: 11, y: 7 }, ">");

// Face 4
// test({ x: 9, y: 11 }, "v");
// test({ x: 9, y: 8 }, "^");
// test({ x: 8, y: 10 }, "<");
// test({ x: 11, y: 10 }, ">");

// Face 5
// test({ x: 13, y: 11 }, "v");
// test({ x: 13, y: 8 }, "^");
// test({ x: 12, y: 9 }, "<");
// test({ x: 15, y: 9 }, ">");
