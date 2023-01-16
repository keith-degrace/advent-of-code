import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";
import { Position } from "../utils/position";

interface Position3D extends Position {
    z: number;
}

interface Scanner {
    id: number;
    position?: Position3D;
    beacons: Position3D[];
}

type RotationFunction = (position: Position3D) => Position3D;

const getRotationFunctions = (): RotationFunction[] => {
    // I snatched this from stack overflow.  Someone described the series of X/Y rotations to get all 24 orientations of a cube in 3D.

    const rotateX = (position: Position3D): Position3D => {
        return { x: position.x, y: -position.z, z: position.y };
    };

    const rotateY = (position: Position3D): Position3D => {
        return { x: -position.z, y: position.y, z: position.x };
    };

    return [
        (position) => position,

        (position) => rotateX(position),
        (position) => rotateY(position),

        (position) => rotateX(rotateX(position)),
        (position) => rotateX(rotateY(position)),
        (position) => rotateY(rotateX(position)),
        (position) => rotateY(rotateY(position)),

        (position) => rotateX(rotateX(rotateX(position))),
        (position) => rotateX(rotateX(rotateY(position))),
        (position) => rotateX(rotateY(rotateX(position))),
        (position) => rotateX(rotateY(rotateY(position))),
        (position) => rotateY(rotateX(rotateX(position))),
        (position) => rotateY(rotateX(rotateX(position))),
        (position) => rotateY(rotateY(rotateY(position))),

        (position) => rotateX(rotateX(rotateX(rotateY(position)))),
        (position) => rotateX(rotateX(rotateY(rotateX(position)))),
        (position) => rotateX(rotateX(rotateY(rotateY(position)))),
        (position) => rotateX(rotateY(rotateX(rotateX(position)))),
        (position) => rotateX(rotateY(rotateY(rotateY(position)))),
        (position) => rotateY(rotateX(rotateX(rotateX(position)))),
        (position) => rotateY(rotateY(rotateY(rotateX(position)))),

        (position) => rotateX(rotateX(rotateX(rotateY(rotateX(position))))),
        (position) => rotateX(rotateY(rotateX(rotateX(rotateX(position))))),
        (position) => rotateX(rotateY(rotateY(rotateY(rotateX(position))))),
    ];
};

const load = (): Scanner[] => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    let scanners: Scanner[] = [];
    let id = 0;

    for (const line of input) {
        if (line.startsWith("---")) {
            scanners.push({ id: id++, beacons: [] });
        } else if (line.length > 0) {
            const coordinates = line.split(",").map((value) => parseInt(value));
            const position = { x: coordinates[0], y: coordinates[1], z: coordinates[2] };

            scanners[scanners.length - 1].beacons.push(position);
        }
    }

    return scanners;
};

const getRelativePositions = (beacon: Position3D, allBeacons: Position3D[]): Position3D[] => {
    const relativePositions: Position3D[] = [];

    for (const otherBeacon of allBeacons) {
        relativePositions.push({
            x: beacon.x - otherBeacon.x,
            y: beacon.y - otherBeacon.y,
            z: beacon.z - otherBeacon.z,
        });
    }

    return relativePositions;
};

const solveScanner = (referenceScanner: Scanner, scannerToSolve: Scanner): boolean => {
    for (const referenceBeacon of referenceScanner.beacons) {
        const referenceRelativePositions = getRelativePositions(referenceBeacon, referenceScanner.beacons);

        for (const rotate of getRotationFunctions()) {
            const rotatedBeacons = scannerToSolve.beacons.map((beacon) => rotate(beacon));

            for (const beacon2 of rotatedBeacons) {
                const relativePositions2 = getRelativePositions(beacon2, rotatedBeacons);

                let overlapCount = 0;

                for (const referenceRelativePosition of referenceRelativePositions) {
                    for (const relativePosition2 of relativePositions2) {
                        if (
                            referenceRelativePosition.x === relativePosition2.x &&
                            referenceRelativePosition.y === relativePosition2.y &&
                            referenceRelativePosition.z === relativePosition2.z
                        ) {
                            overlapCount++;

                            if (overlapCount === 12) {
                                scannerToSolve.position = {
                                    x: beacon2.x - (referenceBeacon.x + referenceScanner.position.x),
                                    y: beacon2.y - (referenceBeacon.y + referenceScanner.position.y),
                                    z: beacon2.z - (referenceBeacon.z + referenceScanner.position.z),
                                };
                                scannerToSolve.beacons = rotatedBeacons;

                                return true;
                            }
                        }
                    }
                }
            }
        }
    }

    return false;
};

const solveScanners = (referenceScanner: Scanner, scannersToSolve: Scanner[]): Scanner[] => {
    const solvedScanners: Scanner[] = [];

    for (const scannerToSolve of scannersToSolve) {
        if (solveScanner(referenceScanner, scannerToSolve)) {
            solvedScanners.push(scannerToSolve);
        }
    }

    return solvedScanners;
};

export const run = () => {
    const scanners = load();

    // Let's assume scanner 0 is at (0,0,0) and base everything off of this.
    scanners[0].position = { x: 0, y: 0, z: 0 };

    const stack = [scanners[0]];
    let scannersToSolve = scanners.slice(1);

    while (stack.length > 0) {
        const current = stack.shift();
        log(`Trying to solve ${scannersToSolve.length} scanners against Scanner ${current.id}`);

        const solvedScanners = solveScanners(current, scannersToSolve);
        log(`  Solved: ${solvedScanners.map((scanner) => scanner.id)}`);

        stack.push(...solvedScanners);

        scannersToSolve = scannersToSolve.filter((scannerToCheck) => scannerToCheck !== current);
    }

    const uniqueBeacons = new Set<string>();

    for (const scanner of scanners) {
        for (const beacon of scanner.beacons) {
            if (scanner.position === undefined) {
                log(`${scanner.id} has no position`);
                continue;
            }

            const x = scanner.position.x + beacon.x;
            const y = scanner.position.y + beacon.y;
            const z = scanner.position.z + beacon.z;

            uniqueBeacons.add(`${x},${y},${z}`);
        }
    }

    log(uniqueBeacons.size);
};

// 671 Too High
// 369 Too High
