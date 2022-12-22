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

const isOccupied = (occupiedPositions: Set<number>, x: number, y: number): boolean => {
    if (x < 0 || x >= ChamberWidth || y < 0) {
        return true;
    }

    return occupiedPositions.has(y * 7 + x);
};

const setOccupied = (occupiedPositions: Set<number>, x: number, y: number): void => {
    if (occupiedPositions.has(y * 7 + x)) {
        throw `(${x},${y})`;
    }

    occupiedPositions.add(y * 7 + x);
};

const addRockToOccupied = (occupiedPositions: Set<number>, rock: Rock): void => {
    for (const point of rock.type.points) {
        setOccupied(occupiedPositions, rock.position.x + point.x, rock.position.y + point.y);
    }
};

const canRockMoveTo = (x: number, y: number, rockType: RockType, occupiedPositions: Set<number>): boolean => {
    return !rockType.points.some((point) => isOccupied(occupiedPositions, point.x + x, point.y + y));
};

const canRockMoveLeft = (rock: Rock, occupiedPositions: Set<number>): boolean => {
    return canRockMoveTo(rock.position.x - 1, rock.position.y, rock.type, occupiedPositions);
};

const pushRockLeft = (rock: Rock, occupiedPositions: Set<number>): boolean => {
    if (!canRockMoveLeft(rock, occupiedPositions)) {
        return false;
    }

    rock.position.x--;
    return true;
};

const canRockMoveRight = (rock: Rock, occupiedPositions: Set<number>): boolean => {
    return canRockMoveTo(rock.position.x + 1, rock.position.y, rock.type, occupiedPositions);
};

const pushRockRight = (rock: Rock, occupiedPositions: Set<number>): boolean => {
    if (!canRockMoveRight(rock, occupiedPositions)) {
        return false;
    }

    rock.position.x++;
    return true;
};

const canRockDrop = (rock: Rock, occupiedPositions: Set<number>): boolean => {
    return canRockMoveTo(rock.position.x, rock.position.y - 1, rock.type, occupiedPositions);
};

const dropRock = (rock: Rock): void => {
    rock.position.y--;
};

const drawChamber = (occupiedPositions: Set<number>, towerHeight: number, currentRock?: Rock): void => {
    const occupiedCopy = new Set<number>(occupiedPositions);
    if (currentRock) {
        addRockToOccupied(occupiedCopy, currentRock);
    }

    let height = Math.max(towerHeight, currentRock ? currentRock.position.y + currentRock.type.height : 0);

    for (let y = height - 1; y >= 0; y--) {
        let row = "";
        for (let x = 0; x < 7; x++) {
            row += isOccupied(occupiedCopy, x, y) ? "#" : ".";
        }
        log(`|${row}|`);
    }

    log(`+-------+`);
};

let diff1 = 0;
let diff2 = 0;
let diff3 = 0;

const simulate = (): void => {
    let fullTowerHeight = 0;
    let currentTowerHeight = 0;
    let currentRockTypeIndex = 0;
    let currentJetIndex = 0;
    let occupiedPositions = new Set<number>();

    for (let i = 0; i < 1000000000000; i++) {
        const currentRock: Rock = {
            type: rockTypes[currentRockTypeIndex],
            position: { x: 2, y: currentTowerHeight + 3 },
        };

        while (true) {
            const jet = input[currentJetIndex];
            currentJetIndex = (currentJetIndex + 1) % input.length;

            if (jet === "<") {
                pushRockLeft(currentRock, occupiedPositions);
            } else if (jet === ">") {
                pushRockRight(currentRock, occupiedPositions);
            }

            if (!canRockDrop(currentRock, occupiedPositions)) {
                break;
            }

            dropRock(currentRock);
        }

        addRockToOccupied(occupiedPositions, currentRock);

        for (const point of currentRock.type.points) {
            currentTowerHeight = Math.max(currentTowerHeight, currentRock.position.y + point.y + 1);
        }

        const y = currentTowerHeight - 1 - 3;
        if (y > 0) {
            const isDead = (x: number, y: number): boolean => {
                if (isOccupied(occupiedPositions, x, y)) {
                    return true;
                }

                if (isOccupied(occupiedPositions, x - 1, y) && isOccupied(occupiedPositions, x + 1, y)) {
                    if (isOccupied(occupiedPositions, x, y + 1) || isOccupied(occupiedPositions, x, y + 2) || isOccupied(occupiedPositions, x, y + 3)) {
                        return true;
                    }
                }

                return false;
            };

            let allDead = true;
            for (let x = 0; x < ChamberWidth; x++) {
                if (!isDead(x, y)) {
                    allDead = false;
                    break;
                }
            }

            if (allDead) {
                //drawChamber(occupiedPositions, currentTowerHeight);
                fullTowerHeight = fullTowerHeight + (y + 1);

                let newOccupiedPositions = new Set<number>();
                let newTowerHeight = currentTowerHeight - y - 1;

                for (let xx = 0; xx < 7; xx++) {
                    for (let yy = 0; yy < newTowerHeight; yy++) {
                        if (isOccupied(occupiedPositions, xx, y + yy + 1)) {
                            setOccupied(newOccupiedPositions, xx, yy);
                        }
                    }
                }

                // log(`[${i}] Chopped at ${y}, ${currentJetIndex}, ${currentTowerHeight}`);

                occupiedPositions = newOccupiedPositions;
                currentTowerHeight = newTowerHeight;

                if (i === 2244) {
                    const rocksLeft = 1000000000000 - i;
                    const times = Math.floor(rocksLeft / 1745);
                    const rocksToSkip = 1745 * times;
                    const towerHeightToAdd = 2752 * times;

                    i += rocksToSkip;
                    fullTowerHeight += towerHeightToAdd;

                    log(`Skipping ahead to ${i}`);
                }

                // if (y === 184) {
                //     log(`[${i}] Height ${fullTowerHeight + currentTowerHeight} d1:${fullTowerHeight - diff1} d2:${i - diff2}`);

                //     diff1 = fullTowerHeight;
                //     diff2 = i;
                // }

                // drawChamber(occupiedPositions, currentTowerHeight);
            }
        }

        currentRockTypeIndex = (currentRockTypeIndex + 1) % rockTypes.length;
    }

    log(fullTowerHeight + currentTowerHeight);
};

/**
 * This one was super cheesy. My first thought was to detect a dead line, and chop tower to save on memory.  When I did
 * this, I noticed that the new truncated tower repeated itself.  I realized this was one of those puzzles where you can
 * detect a pattern and skip ahead.
 *
 * The cheesy part is I detected the pattern by hand, and calculated the multiplication factors.  You can see the magic
 * numbers in the code, e.g. 2244, 1745, 2752.
 *
 * I suppose there might be a way to auto detect this?
 *
 */

simulate();
