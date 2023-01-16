import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

export const run = () => {
    let values = fs
        .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
        .split("\n")
        .map((value) => parseInt(value));

    let increaseCount = 0;

    let lastSum = values[0] + values[1] + values[2];
    for (let i = 3; i < values.length; i++) {
        const currentSum = values[i - 2] + values[i - 1] + values[i];
        if (lastSum < currentSum) {
            increaseCount++;
        }

        lastSum = currentSum;
    }

    log(increaseCount);
};
