import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

export const run = () => {
    const fishes: number[] = fs
        .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
        .split(",")
        .map((value) => parseInt(value));

    for (let day = 0; day < 80; day++) {
        const fishCount = fishes.length;
        for (let i = 0; i < fishCount; i++) {
            if (fishes[i] === 0) {
                fishes[i] = 6;
                fishes.push(8);
            } else {
                fishes[i]--;
            }
        }
    }

    log(fishes.length);
};
