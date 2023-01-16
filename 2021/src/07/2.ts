import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

export const run = () => {
    const crabs: number[] = fs
        .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
        .split(",")
        .map((value) => parseInt(value));

    const min = Math.min(...crabs);
    const max = Math.max(...crabs);

    let bestCost = Infinity;

    for (let position = min; position <= max; position++) {
        let cost = 0;

        for (const crab of crabs) {
            const dx = Math.abs(crab - position);
            for (let fuel = 1; fuel <= dx; fuel++) {
                cost += fuel;
            }
        }

        bestCost = Math.min(cost, bestCost);
    }

    log(bestCost);
};
