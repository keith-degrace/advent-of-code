import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

const basePattern = [0, 1, 0, -1];

const getFactor = (x: number, y: number): number => {
    return basePattern[Math.floor((x + 1) / (y + 1)) % basePattern.length];
};

export const run = () => {
    let signal = fs
        .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
        .split("")
        .map((value) => parseInt(value));

    for (let phase = 0; phase < 100; phase++) {
        let newSignal = [];
        for (let y = 0; y < signal.length; y++) {
            let value = 0;
            for (let x = y; x < signal.length; x++) {
                value += signal[x] * getFactor(x, y);
            }

            newSignal.push(Math.abs(value) % 10);
        }

        signal = newSignal;
    }

    log(signal.slice(0, 8).join(""));
};
