import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

const getMostCommon = (values: string[], position: number): string => {
    let ones = 0;
    let zeroes = 0;

    for (const value of values) {
        if (value[position] === "0") {
            ones++;
        } else {
            zeroes++;
        }
    }

    if (ones > zeroes) {
        return "1";
    } else {
        return "0";
    }
};

const getLeastCommon = (values: string[], position: number): string => {
    return getMostCommon(values, position) === "0" ? "1" : "0";
};

const getOxygenGeneratorRating = (values: string[]): number => {
    const bitCount = values[0].length;

    let filteredValues = [...values];

    for (let position = 0; position < bitCount; position++) {
        const mostCommon = getMostCommon(filteredValues, position);

        filteredValues = filteredValues.filter((value) => value[position] === mostCommon);

        if (filteredValues.length === 1) {
            return parseInt(filteredValues[0], 2);
        }
    }

    return undefined;
};

const getScrubberRating = (values: string[]): number => {
    const bitCount = values[0].length;

    let filteredValues = [...values];

    for (let position = 0; position < bitCount; position++) {
        const leastCommon = getLeastCommon(filteredValues, position);

        filteredValues = filteredValues.filter((value) => value[position] === leastCommon);

        if (filteredValues.length === 1) {
            return parseInt(filteredValues[0], 2);
        }
    }

    return undefined;
};

export const run = () => {
    const values = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\n");

    log(getOxygenGeneratorRating(values) * getScrubberRating(values));
};
