import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

interface Constants {
    k1: number;
    k2: number;
    k3: number;
}

// The program is broken into 14 different nearly identical code chunks.  This extracts the constants that differ in each chunk.
const getConstants = (): Constants[] => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    const chunkCount = input.length / 18;

    const constants: Constants[] = [];

    for (let chunk = 0; chunk < chunkCount; chunk++) {
        const k1 = parseInt(input[chunk * 18 + 5].split(" ")[2]);
        const k2 = parseInt(input[chunk * 18 + 15].split(" ")[2]);
        const k3 = parseInt(input[chunk * 18 + 4].split(" ")[2]);

        constants.push({ k1, k2, k3 });
    }

    return constants;
};

// The program is broken into 14 different nearly identical code chunks.  This function is a result of analyzing
// the code, converting it to javascript, and simplifying it.
const executePart = (w: number, z: number, constants: Constants): number => {
    if ((z % 26) + constants.k1 !== w) {
        return Math.floor(z / constants.k3) * 26 + w + constants.k2;
    } else {
        return Math.floor(z / constants.k3);
    }
};

const solveInterestingInputs = (): Set<number>[] => {
    const constants = getConstants();

    const interestingInputs: Set<number>[] = [
        new Set(),
        new Set(),
        new Set(),
        new Set(),
        new Set(),
        new Set(),
        new Set(),
        new Set(),
        new Set(),
        new Set(),
        new Set(),
        new Set(),
        new Set(),
        new Set(),
        new Set([0]),
    ];

    // A bit of a cheezy strategy.  The end result that we want is a zero value for Z.  This works
    // backwards and figures out all Z inputs for a code chunk that will produce a result that will
    // ultimately result in a final zero value for Z.
    for (let chunk = 13; chunk >= 0; chunk--) {
        for (let inputZ = 0; inputZ < 1000000; inputZ++) {
            for (let w = 1; w <= 9; w++) {
                const outputZ = executePart(w, inputZ, constants[chunk]);

                if (interestingInputs[chunk + 1].has(outputZ)) {
                    interestingInputs[chunk].add(inputZ);
                }
            }
        }
    }

    return interestingInputs;
};

const getValidSerialNumbers = (): number[] => {
    const constants = getConstants();
    const interestingInputZ: Set<number>[] = solveInterestingInputs();

    // Once we solve all the Z inputs that can result in a final zero Z, we iteratively search for all valid serial numbers.

    const validSerialNumbers: number[] = [];

    const search = (chunk: number, previousZ: number, result: string) => {
        if (result.length === 14) {
            validSerialNumbers.push(parseInt(result));
            return;
        }

        for (let w = 1; w <= 9; w++) {
            const outputZ = executePart(w, previousZ, constants[chunk]);

            if (interestingInputZ[chunk + 1].has(outputZ)) {
                search(chunk + 1, outputZ, `${result}${w}`);
            }
        }
    };

    search(0, 0, "");

    return validSerialNumbers;
};

export const run = () => {
    const validSerialNumbers: number[] = getValidSerialNumbers();

    validSerialNumbers.sort((a, b) => b - a);

    log(validSerialNumbers[0]);
};
