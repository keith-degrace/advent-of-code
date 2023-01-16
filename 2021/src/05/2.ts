import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

export const run = () => {
    const lines = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\n");

    const coverage: Record<string, number> = {};

    for (const line of lines) {
        const parts = line.split(" -> ");
        const point1 = parts[0].split(",").map((value) => parseInt(value));
        const point2 = parts[1].split(",").map((value) => parseInt(value));

        const dx = point1[0] === point2[0] ? 0 : point1[0] < point2[0] ? 1 : -1;
        const dy = point1[1] === point2[1] ? 0 : point1[1] < point2[1] ? 1 : -1;

        let x = point1[0];
        let y = point1[1];

        while (x !== point2[0] + dx || y !== point2[1] + dy) {
            coverage[`${x},${y}`] = (coverage[`${x},${y}`] ?? 0) + 1;

            x += dx;
            y += dy;
        }
    }

    let count = 0;
    for (const value of Object.values(coverage)) {
        if (value > 1) {
            count++;
        }
    }

    log(count);
};
