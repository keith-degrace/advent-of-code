import * as fs from "fs";
import * as path from "path";

const input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split(/\n/);

let exclusionCount: number = 0;

const isWithin = (value: number, range: number[]): boolean => value >= range[0] && value <= range[1];

for (const pair of input) {
    const ranges: string[] = pair.split(",");
    const range1: number[] = ranges[0].split("-").map((value) => Number.parseInt(value));
    const range2: number[] = ranges[1].split("-").map((value) => Number.parseInt(value));

    if (isWithin(range1[0], range2) || isWithin(range1[1], range2) || isWithin(range2[0], range1) || isWithin(range2[1], range1)) {
        exclusionCount++;
    }
}

console.log(exclusionCount);
