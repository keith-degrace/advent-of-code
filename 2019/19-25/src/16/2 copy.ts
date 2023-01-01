import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";
import { startTimer, stopTimer } from "../utils/timer";

const basePattern = [0, 1, 0, -1];

const getFactor = (x: number, y: number): number => {
    return basePattern[Math.floor((x + 1) / (y + 1)) % basePattern.length];
};

export const run = () => {
    let input = fs
        .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
        .split("")
        .map((value) => parseInt(value));

    const offset = parseInt(input.slice(0, 8).join(""));

    for (let repeat = 1; repeat < 10000; repeat++) {
        log(`Repeat: ${repeat}`);
        startTimer();

        let signal = [];
        for (let i = 0; i < repeat; i++) {
            signal.push(...input);
        }

        for (let phase = 0; phase < 100; phase++) {
            let newSignal = [];
            for (let y = 0; y < signal.length; y++) {
                let value = 0;

                for (let x = 0; x < signal.length; x++) {
                    value += signal[x] * getFactor(x, y);
                }

                newSignal.push(Math.abs(value) % 10);
            }

            signal = newSignal;
        }

        log(`Message: ${signal.slice(offset, 8)}`);

        stopTimer();
    }
};
