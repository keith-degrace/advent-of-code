import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

type Pair = [number | Pair, number | Pair];

const parse = (value: string): Pair => {
    let index = 0;

    const readPair = (): Pair => {
        if (value[index] !== "[") {
            throw "wut?";
        }

        // Skip over opening bracket
        index++;

        // Read the left value.
        let left: Pair | number;

        if (value[index] === "[") {
            left = readPair();
            index++;
        } else {
            const end = value.indexOf(",", index);
            left = parseInt(value.slice(index, end));
            index = end + 1;
        }

        // Read the right value.
        let right: Pair | number;

        if (value[index] === "[") {
            right = readPair();
            index++;
        } else {
            const end = value.indexOf("]", index);
            right = parseInt(value.slice(index, end));
            index = end + 1;
        }

        return [left, right];
    };

    return readPair();
};

const add = (pair1: Pair, pair2: Pair): Pair => {
    return [pair1, pair2];
};

const addToFirstNumberToLeftOfMarker = (pair: Pair, value: number): void => {
    let foundMarker = false;
    let foundValue = false;

    const search = (currentPair: Pair) => {
        if (foundValue) {
            return;
        }

        if (typeof currentPair[1] !== "number") {
            search(currentPair[1]);
        } else {
            if (foundMarker) {
                if (!foundValue) {
                    currentPair[1] += value;
                    foundValue = true;
                    return;
                }
            } else if (currentPair[1] === Infinity) {
                foundMarker = true;
            }
        }

        if (typeof currentPair[0] !== "number") {
            search(currentPair[0]);
        } else {
            if (foundMarker) {
                if (!foundValue) {
                    currentPair[0] += value;
                    foundValue = true;
                    return;
                }
            } else if (currentPair[0] === Infinity) {
                foundMarker = true;
            }
        }
    };

    search(pair);
};

const addToFirstNumberToRightOfMarker = (pair: Pair, value: number): void => {
    let foundMarker = false;
    let foundValue = false;

    const search = (currentPair: Pair) => {
        if (foundValue) {
            return;
        }

        if (typeof currentPair[0] !== "number") {
            search(currentPair[0]);
        } else {
            if (foundMarker) {
                if (!foundValue) {
                    currentPair[0] += value;
                    foundValue = true;
                    return;
                }
            } else if (currentPair[0] === Infinity) {
                foundMarker = true;
            }
        }

        if (typeof currentPair[1] !== "number") {
            search(currentPair[1]);
        } else {
            if (foundMarker) {
                if (!foundValue) {
                    currentPair[1] += value;
                    foundValue = true;
                    return;
                }
            } else if (currentPair[1] === Infinity) {
                foundMarker = true;
            }
        }
    };

    search(pair);
};

const reduce = (pair: Pair): Pair => {
    const reduceNestedPair = (nestedPair: Pair, level: number): boolean => {
        if (level !== 4) {
            if (typeof nestedPair[0] !== "number" && reduceNestedPair(nestedPair[0], level + 1)) {
                return true;
            }

            if (typeof nestedPair[1] !== "number" && reduceNestedPair(nestedPair[1], level + 1)) {
                return true;
            }

            return false;
        }

        if (typeof nestedPair[0] !== "number") {
            const leftValue: number = nestedPair[0][0] as number;
            const rightValue: number = nestedPair[0][1] as number;

            nestedPair[0] = Infinity;

            addToFirstNumberToLeftOfMarker(pair, leftValue);
            addToFirstNumberToRightOfMarker(pair, rightValue);

            nestedPair[0] = 0;

            return true;
        }

        if (typeof nestedPair[1] !== "number") {
            const leftValue: number = nestedPair[1][0] as number;
            const rightValue: number = nestedPair[1][1] as number;

            nestedPair[1] = Infinity;

            addToFirstNumberToLeftOfMarker(pair, leftValue);
            addToFirstNumberToRightOfMarker(pair, rightValue);

            nestedPair[1] = 0;

            return true;
        }

        return false;
    };

    const splitNumber = (nestedPair: Pair): boolean => {
        if (typeof nestedPair[0] !== "number" && splitNumber(nestedPair[0])) {
            return true;
        }

        if (typeof nestedPair[0] === "number" && nestedPair[0] >= 10) {
            const leftValue = Math.floor(nestedPair[0] / 2);
            const rightValue = nestedPair[0] - leftValue;
            nestedPair[0] = [leftValue, rightValue];
            return true;
        }

        if (typeof nestedPair[1] !== "number" && splitNumber(nestedPair[1])) {
            return true;
        }

        if (typeof nestedPair[1] === "number" && nestedPair[1] >= 10) {
            const leftValue = Math.floor(nestedPair[1] / 2);
            const rightValue = nestedPair[1] - leftValue;
            nestedPair[1] = [leftValue, rightValue];
            return true;
        }

        return false;
    };

    while (reduceNestedPair(pair, 1) || splitNumber(pair)) {}

    return pair;
};

const getMagnitude = (pair: Pair): number => {
    let leftMagnitude = 0;
    if (typeof pair[0] === "number") {
        leftMagnitude = pair[0];
    } else {
        leftMagnitude = getMagnitude(pair[0]);
    }

    let rightMagnitude = 0;
    if (typeof pair[1] === "number") {
        rightMagnitude = pair[1];
    } else {
        rightMagnitude = getMagnitude(pair[1]);
    }

    return leftMagnitude * 3 + rightMagnitude * 2;
};

export const run = () => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    let pair = parse(input[0]);
    for (let i = 1; i < input.length; i++) {
        pair = add(pair, parse(input[i]));
        pair = reduce(pair);
    }

    log(getMagnitude(pair));
};
