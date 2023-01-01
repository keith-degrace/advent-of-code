import * as fs from "fs";
import * as path from "path";
import { runSimpleIntCode } from "../utils/intcode";
import { log } from "../utils/log";

interface Row {
    x1: number;
    x2: number;
    y: number;
}

const isAffectedByBeam = (program: number[], x: number, y: number): boolean => {
    return runSimpleIntCode({ ...program }, [x, y])[0] === 1;
};

export const run = () => {
    let program = fs
        .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
        .split(",")
        .map((value) => parseInt(value));

    let x1 = 0;
    let x2 = 100;

    let lastRows: Row[] = [];

    for (let y = 13; ; y++) {
        // Find the new X1.
        while (isAffectedByBeam(program, x1, y)) x1--;
        while (!isAffectedByBeam(program, x1, y) && x1 < x2) x1++;

        // Find the new X2
        while (isAffectedByBeam(program, x2, y)) x2++;
        while (!isAffectedByBeam(program, x2, y) && x1 < x2) x2--;

        // Log this row.
        lastRows.push({ x1, x2, y });

        // Only need 100 rows.
        if (lastRows.length > 100) {
            lastRows.shift();
        }

        // Compute the maximum width of the current set of rows.  Measuring from X1 of last row to X2
        // of first row. Look at puzzle diagram.
        const maxWidth = lastRows[0].x2 - lastRows[lastRows.length - 1].x1 + 1;
        if (maxWidth === 100) {
            log(`${lastRows[lastRows.length - 1].x1 * 10000 + lastRows[0].y}`);
            break;
        }
    }
};
