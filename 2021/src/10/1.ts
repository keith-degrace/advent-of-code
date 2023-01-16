import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

interface Pair {
    open: string;
    close: string;
}

const ValidPairs: Pair[] = [
    { open: "(", close: ")" },
    { open: "[", close: "]" },
    { open: "{", close: "}" },
    { open: "<", close: ">" },
];

const getPair = (open: string): Pair => {
    return ValidPairs.find((pair) => pair.open === open);
};

const getIllegalCharacterScore = (char: string): number => {
    switch (char) {
        case ")":
            return 3;
        case "]":
            return 57;
        case "}":
            return 1197;
        case ">":
            return 25137;
    }

    throw "huh?";
};

const getErrorScore = (line: string): number => {
    let firstPair = getPair(line[0]);
    if (!firstPair) {
        throw "huh?";
    }

    let stack = [firstPair];
    let index = 1;

    while (stack.length > 0) {
        const currentPair = stack[stack.length - 1];

        if (line[index] === currentPair.close) {
            stack.pop();
        } else {
            let newPair = getPair(line[index]);
            if (!newPair) {
                return getIllegalCharacterScore(line[index]);
            }

            stack.push(newPair);
        }

        index++;
        if (index >= line.length) {
            break;
        }
    }

    return 0;
};

export const run = () => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    let score = 0;

    for (const line of input) {
        score += getErrorScore(line);
    }

    log(score);
};
