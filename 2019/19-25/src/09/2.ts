import * as fs from "fs";
import * as path from "path";
import { runSimpleIntCode } from "../utils/intcode";
import { log } from "../utils/log";

export const run = () => {
    let program = fs
        .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
        .split(",")
        .map((value) => parseInt(value));

    const output = runSimpleIntCode(program, [2]);

    log(output);
};
