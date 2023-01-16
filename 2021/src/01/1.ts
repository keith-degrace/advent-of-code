import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

export const run = () => {
    let values = fs
        .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
        .split("\n")
        .map((value) => parseInt(value));

    let increaseCount = 0;

    for (let i = 1; i < values.length; i++) {
        if (values[i - 1] < values[i]) {
            increaseCount++;
        }
    }

    log(increaseCount);
};
