import * as fs from "fs";
import * as path from "path";
import { __values } from "tslib";
import { updatePartiallyEmittedExpression } from "typescript";
import { log } from "../utils/log";

type Decoder = Record<string, string>;

const getAllPossibleDecoders = (): Decoder[] => {
    const decoders: Decoder[] = [];

    const permutate = (current: string[], rest: string[]): void => {
        if (rest.length === 0) {
            decoders.push({
                a: current[0],
                b: current[1],
                c: current[2],
                d: current[3],
                e: current[4],
                f: current[5],
                g: current[6],
            });
            return;
        }

        for (const letter of rest) {
            const subRest = rest.filter((value) => value !== letter);
            permutate([...current, letter], subRest);
        }
    };

    permutate([], "abcdefg".split(""));

    return decoders;
};

const decode = (decoder: Decoder, encodedValue: string): string => {
    return encodedValue
        .split("")
        .map((value) => decoder[value])
        .sort()
        .join("");
};

const getDigit = (decoder: Decoder, encodedValue: string): number => {
    const patterns = ["abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"];
    const decodedValue = decode(decoder, encodedValue);
    return patterns.findIndex((pattern) => pattern === decodedValue);
};

const findDecoder = (decoders: Decoder[], encodedValues: string[]): Decoder => {
    for (const decoder of decoders) {
        let allDecoded = true;

        for (const encodedValue of encodedValues) {
            const digit = getDigit(decoder, encodedValue);
            if (digit === -1) {
                allDecoded = false;
                break;
            }
        }

        if (allDecoded) {
            return decoder;
        }
    }

    throw "huh?";
};

export const run = () => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\n");
    const allDecoders = getAllPossibleDecoders();

    let sum = 0;

    for (const line of input) {
        const signals = line.split(" | ")[0].trim().split(" ");
        const output = line.split(" | ")[1].trim().split(" ");

        const decoder = findDecoder(allDecoders, [...signals, ...output]);

        let decodedOutput = "";
        for (const encodedValue of output) {
            decodedOutput += getDigit(decoder, encodedValue);
        }

        sum += parseInt(decodedOutput);
    }

    log(sum);
};
