import * as fs from "fs";
import * as path from "path";
import { log } from "../utils";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("");

interface Position {
    x: number;
    y: number;
}

interface RockType {
    name: string;
    points: Position[];
    width: number;
    height: number;
}

const createRockType = (name: string, points: Position[]): RockType => {
    return {
        name,
        points,
        width: Math.max(...points.map((point) => point.x)) + 1,
        height: Math.max(...points.map((point) => point.y)) + 1,
    };
};

const rockTypes: RockType[] = [
    createRockType("Horizontal Line", [
        { x: 0, y: 0 },
        { x: 1, y: 0 },
        { x: 2, y: 0 },
        { x: 3, y: 0 },
    ]),
    createRockType("Cross", [
        { x: 0, y: 1 },
        { x: 1, y: 1 },
        { x: 2, y: 1 },
        { x: 1, y: 0 },
        { x: 1, y: 2 },
    ]),
    createRockType("Flipped L", [
        { x: 0, y: 0 },
        { x: 1, y: 0 },
        { x: 2, y: 0 },
        { x: 2, y: 1 },
        { x: 2, y: 2 },
    ]),
    createRockType("Vertical Line", [
        { x: 0, y: 0 },
        { x: 0, y: 1 },
        { x: 0, y: 2 },
        { x: 0, y: 3 },
    ]),
    createRockType("Square", [
        { x: 0, y: 0 },
        { x: 1, y: 0 },
        { x: 0, y: 1 },
        { x: 1, y: 1 },
    ]),
];

interface Rock {
    type: RockType;
    position: Position;
}

const ChamberWidth = 7;

const isOccupied = (occupiedPositions: Set<string>, x: number, y: number): boolean => {
    if (x < 0 || x >= ChamberWidth || y < 0) {
        return true;
    }

    return occupiedPositions.has(`(${x},${y})`);
};

const setOccupied = (occupiedPositions: Set<string>, rock: Rock): void => {
    for (const point of rock.type.points) {
        const x = rock.position.x + point.x;
        const y = rock.position.y + point.y;

        if (occupiedPositions.has(`(${x},${y})`)) {
            throw `(${x},${y})`;
        }

        occupiedPositions.add(`(${x},${y})`);
    }
};

const canRockMoveTo = (x: number, y: number, rockType: RockType, occupiedPositions: Set<string>): boolean => {
    return !rockType.points.some((point) => isOccupied(occupiedPositions, point.x + x, point.y + y));
};

const canRockMoveLeft = (rock: Rock, occupiedPositions: Set<string>): boolean => {
    return canRockMoveTo(rock.position.x - 1, rock.position.y, rock.type, occupiedPositions);
};

const pushRockLeft = (rock: Rock, occupiedPositions: Set<string>): boolean => {
    if (!canRockMoveLeft(rock, occupiedPositions)) {
        return false;
    }

    rock.position.x--;
    return true;
};

const canRockMoveRight = (rock: Rock, occupiedPositions: Set<string>): boolean => {
    return canRockMoveTo(rock.position.x + 1, rock.position.y, rock.type, occupiedPositions);
};

const pushRockRight = (rock: Rock, occupiedPositions: Set<string>): boolean => {
    if (!canRockMoveRight(rock, occupiedPositions)) {
        return false;
    }

    rock.position.x++;
    return true;
};

const canRockDrop = (rock: Rock, occupiedPositions: Set<string>): boolean => {
    return canRockMoveTo(rock.position.x, rock.position.y - 1, rock.type, occupiedPositions);
};

const dropRock = (rock: Rock): void => {
    rock.position.y--;
};

const drawChamber = (occupiedPositions: Set<string>, towerHeight: number, currentRock?: Rock): void => {
    const occupiedCopy = new Set<string>(occupiedPositions);
    if (currentRock) {
        setOccupied(occupiedCopy, currentRock);
    }

    let height = Math.max(towerHeight, currentRock ? currentRock.position.y + currentRock.type.height : 0);

    for (let y = height; y >= 0; y--) {
        let row = "";
        for (let x = 0; x < 7; x++) {
            row += isOccupied(occupiedCopy, x, y) ? "#" : ".";
        }
        log(`|${row}|`);
    }

    log(`+-------+`);
};

let wef = 0;

const simulate = (): void => {
    let towerHeight = 0;
    let currentRockTypeIndex = 0;
    let currentJetIndex = 0;
    const occupiedPositions = new Set<string>();

    for (let i = 0; i < 2022; i++) {
        const currentRock: Rock = {
            type: rockTypes[currentRockTypeIndex],
            position: { x: 2, y: towerHeight + 3 },
        };

        //log(`${currentRock.type.name} begins falling from ${JSON.stringify(currentRock.position)}`);
        //log(`Tower height: ${towerHeight}`);
        //drawChamber(occupiedPositions, towerHeight, currentRock);

        while (true) {
            const jet = input[currentJetIndex];
            currentJetIndex = (currentJetIndex + 1) % input.length;

            if (jet === "<") {
                if (pushRockLeft(currentRock, occupiedPositions)) {
                    //log("Jet of gas pushes rock left");
                } else {
                    //log("Jet of gas pushes rock left, but nothing happens");
                }
            } else if (jet === ">") {
                if (pushRockRight(currentRock, occupiedPositions)) {
                    //log("Jet of gas pushes rock right");
                } else {
                    //log("Jet of gas pushes rock right, but nothing happens");
                }
            }

            //drawChamber(occupiedPositions, towerHeight, currentRock);

            if (canRockDrop(currentRock, occupiedPositions)) {
                dropRock(currentRock);
                //log("Rock falls 1 unit");
                //drawChamber(occupiedPositions, towerHeight, currentRock);
            } else {
                //log("Rock falls 1 unit, causing it to come to rest");
                break;
            }
        }

        setOccupied(occupiedPositions, currentRock);
        //drawChamber(occupiedPositions, towerHeight);

        for (const point of currentRock.type.points) {
            towerHeight = Math.max(towerHeight, currentRock.position.y + point.y + 1);
        }

        currentRockTypeIndex = (currentRockTypeIndex + 1) % rockTypes.length;
    }

    log(towerHeight);
};

simulate();
