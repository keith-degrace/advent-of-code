import * as fs from "fs";
import * as path from "path";
import { runSimpleIntCode } from "../utils/intcode";
import { log } from "../utils/log";

const isAffectedByBeam = (program: number[], x: number, y: number): boolean => {
    return runSimpleIntCode({ ...program }, [x, y])[0] === 1;
};

export const run = () => {
    let program = fs
        .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
        .split(",")
        .map((value) => parseInt(value));

    let affectedCount = 0;
    for (let x = 0; x < 50; x++) {
        for (let y = 0; y < 50; y++) {
            if (isAffectedByBeam(program, x, y)) {
                affectedCount++;
            }
        }
    }

    log(affectedCount);
};
