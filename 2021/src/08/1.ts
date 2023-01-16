import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

export const run = () => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\n");

    let count = 0;

    for (const line of input) {
        const output = line.split(" | ")[1].trim().split(" ");

        for (const sequence of output) {
            if (sequence.length === 2 || sequence.length === 3 || sequence.length === 4 || sequence.length === 7) {
                count++;
            }
        }
    }

    log(count);
};
