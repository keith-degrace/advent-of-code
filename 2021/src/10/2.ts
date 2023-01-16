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

const getMissingCloseScore = (char: string): number => {
    switch (char) {
        case ")":
            return 1;
        case "]":
            return 2;
        case "}":
            return 3;
        case ">":
            return 4;
    }

    throw "huh?";
};

const getMissingCloses = (line: string): string[] => {
    let stack = [];

    for (let i = 0; i < line.length; i++) {
        let currentPair = stack[stack.length - 1];

        if (currentPair && line[i] == currentPair.close) {
            stack.pop();
        } else {
            let newPair = getPair(line[i]);
            if (!newPair) {
                return undefined;
            }

            stack.push(newPair);
        }
    }

    return stack.reverse().map((pair) => pair.close);
};

export const run = () => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    let scores = [];

    for (const line of input) {
        const missingCloses = getMissingCloses(line);
        if (!missingCloses) {
            continue;
        }

        let lineScore = 0;
        for (const missingClose of missingCloses) {
            lineScore *= 5;
            lineScore += getMissingCloseScore(missingClose);
        }

        scores.push(lineScore);
    }

    scores = scores.sort((a, b) => b - a);

    log(scores[(scores.length - 1) / 2]);
};
