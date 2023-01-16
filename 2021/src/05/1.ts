import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

const isHorizontal = (point1: number[], point2: number[]): boolean => {
    return point1[1] === point2[1];
};

const isVertical = (point1: number[], point2: number[]): boolean => {
    return point1[0] === point2[0];
};

export const run = () => {
    const lines = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\n");

    const coverage: Record<string, number> = {};

    for (const line of lines) {
        const parts = line.split(" -> ");
        const point1 = parts[0].split(",").map((value) => parseInt(value));
        const point2 = parts[1].split(",").map((value) => parseInt(value));

        if (isHorizontal(point1, point2)) {
            const x1 = Math.min(point1[0], point2[0]);
            const x2 = Math.max(point1[0], point2[0]);
            const y = point1[1];

            for (let x = x1; x <= x2; x++) {
                coverage[`${x},${y}`] = (coverage[`${x},${y}`] ?? 0) + 1;
            }
        } else if (isVertical(point1, point2)) {
            const x = point1[0];
            const y1 = Math.min(point1[1], point2[1]);
            const y2 = Math.max(point1[1], point2[1]);

            for (let y = y1; y <= y2; y++) {
                coverage[`${x},${y}`] = (coverage[`${x},${y}`] ?? 0) + 1;
            }
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
