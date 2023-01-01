import * as fs from "fs";
import * as path from "path";
import { runSimpleIntCode } from "../utils/intcode";
import { log } from "../utils/log";

export const run = () => {
    let program = fs
        .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
        .split(",")
        .map((value) => parseInt(value));

    program[1] = 12;
    program[2] = 2;

    runSimpleIntCode(program);

    log(program[0]);
};
