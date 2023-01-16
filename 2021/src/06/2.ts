import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

export const run = () => {
    const initialFishes: number[] = fs
        .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
        .split(",")
        .map((value) => parseInt(value));

    let fishes = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
    for (const initialFish of initialFishes) {
        fishes[initialFish]++;
    }

    for (let day = 0; day < 256; day++) {
        fishes[7] += fishes[0];
        fishes[9] = fishes[0];

        fishes = fishes.slice(1);
    }

    let sum = 0;
    for (const fish of fishes) {
        sum += fish;
    }

    log(sum);
};
