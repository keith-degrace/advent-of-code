import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";
import { Position } from "../utils/position";

interface Position3D extends Position {
    z: number;
}

interface Beacon {
    position: Position3D;
    relativePositions: Position3D[];
}

interface Scanner {
    id: number;
    orientations: Beacon[][];
}

const rotateX = (position: Position3D): Position3D => {
    return { x: position.x, y: -position.z, z: position.y };
};

const rotateY = (position: Position3D): Position3D => {
    return { x: -position.z, y: position.y, z: position.x };
};

const rotatePosition = (position: Position3D): Position3D[] => {
    // I snatched this from stack overflow.  Someone described the series of X/Y rotations to get all 24 orientations of a cube in 3D.
    return [
        position,

        rotateX(position),
        rotateY(position),

        rotateX(rotateX(position)),
        rotateX(rotateY(position)),
        rotateY(rotateX(position)),
        rotateY(rotateY(position)),

        rotateX(rotateX(rotateX(position))),
        rotateX(rotateX(rotateY(position))),
        rotateX(rotateY(rotateX(position))),
        rotateX(rotateY(rotateY(position))),
        rotateY(rotateX(rotateX(position))),
        rotateY(rotateX(rotateX(position))),
        rotateY(rotateY(rotateY(position))),

        rotateX(rotateX(rotateX(rotateY(position)))),
        rotateX(rotateX(rotateY(rotateX(position)))),
        rotateX(rotateX(rotateY(rotateY(position)))),
        rotateX(rotateY(rotateX(rotateX(position)))),
        rotateX(rotateY(rotateY(rotateY(position)))),
        rotateY(rotateX(rotateX(rotateX(position)))),
        rotateY(rotateY(rotateY(rotateX(position)))),

        rotateX(rotateX(rotateX(rotateY(rotateX(position))))),
        rotateX(rotateY(rotateX(rotateX(rotateX(position))))),
        rotateX(rotateY(rotateY(rotateY(rotateX(position))))),
    ];
};

const load = (): Scanner[] => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    let scanners: Scanner[] = [];
    let id = 0;

    for (const line of input) {
        if (line.startsWith("---")) {
            scanners.push({ id: id++, orientations: [[], [], [], [], [], [], [], [], [], [], [], [], [], [], [], [], [], [], [], [], [], [], [], []] });
        } else if (line.length > 0) {
            const coordinates = line.split(",").map((value) => parseInt(value));
            const position = { x: coordinates[0], y: coordinates[1], z: coordinates[2] };
            const positions = rotatePosition(position);

            for (let i = 0; i < positions.length; i++) {
                scanners[scanners.length - 1].orientations[i].push({
                    position: positions[i],
                    relativePositions: [],
                });
            }
        }
    }

    for (const scanner of scanners) {
        for (const beacons of scanner.orientations) {
            for (const beacon of beacons) {
                for (const otherBeacon of beacons) {
                    beacon.relativePositions.push({
                        x: beacon.position.x - otherBeacon.position.x,
                        y: beacon.position.y - otherBeacon.position.y,
                        z: beacon.position.z - otherBeacon.position.z,
                    });
                }
            }
        }
    }

    return scanners;
};

const getRelativePosition = (scanner1: Scanner, scanner2: Scanner): Position3D => {
    for (const scanner1Beacons of scanner1.orientations) {
        for (const scanner1Beacon of scanner1Beacons) {
            for (const scanner2Beacons of scanner2.orientations) {
                for (const scanner2Beacon of scanner2Beacons) {
                    let overlapCount = 0;

                    for (const relativePosition1 of scanner1Beacon.relativePositions) {
                        for (const relativePosition2 of scanner2Beacon.relativePositions) {
                            if (relativePosition1.x === relativePosition2.x && relativePosition1.y === relativePosition2.y && relativePosition1.z === relativePosition2.z) {
                                overlapCount++;

                                if (overlapCount === 12) {
                                    return {
                                        x: scanner1Beacon.position.x - scanner2Beacon.position.x,
                                        y: scanner1Beacon.position.y - scanner2Beacon.position.y,
                                        z: scanner1Beacon.position.z - scanner2Beacon.position.z,
                                    };
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    return undefined;
};

const getOverlappingScanners = (scanner: Scanner, scannersToCheck: Scanner[]) => {
    const overlappingScanners: Scanner[] = [];

    for (const scannerToCheck of scannersToCheck) {
        if (scanner === scannerToCheck) {
            continue;
        }

        const relativePosition = getRelativePosition(scanner, scannerToCheck);
        if (relativePosition) {
            log(`${scannerToCheck.id}: ${JSON.stringify(relativePosition)}`);
            overlappingScanners.push(scannerToCheck);
        }
    }

    return overlappingScanners;
};

export const run = () => {
    const scanners = load();

    const stack = [scanners[0]];
    let scannersToCheck = scanners.slice(1);

    while (stack.length > 0) {
        const current = stack.shift();
        log(`Checking Scanner ${current.id} against ${scannersToCheck.length} scanners`);

        const results = getOverlappingScanners(current, scannersToCheck);
        log(`  Found: ${results.map((scanner) => scanner.id)}`);

        for (const result of results) {
            if (!stack.includes(result) && scannersToCheck.includes(result)) {
                stack.push(result);
            }
        }

        scannersToCheck = scannersToCheck.filter((scannerToCheck) => scannerToCheck !== current);
    }
};
