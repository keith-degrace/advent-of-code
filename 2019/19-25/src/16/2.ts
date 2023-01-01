import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

export const run = () => {
    let input = fs
        .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
        .split("")
        .map((value) => parseInt(value));

    const offset = parseInt(input.slice(0, 7).join(""));

    let signal = [];
    for (let i = 0; i < 10000; i++) {
        signal.push(...input);
    }

    for (let phase = 0; phase < 100; phase++) {
        let newSignal = new Array(signal.length);
        let lastSum;

        for (let y = offset; y < signal.length; y++) {
            const patternRepeat = y + 1;
            const patternWidth = patternRepeat * 4;
            const addIndex = y;
            const subIndex = addIndex + patternRepeat + patternRepeat;

            // Leverage previous sum.  The only different on the next row is removing that first value, and adding two more values at the end of the pattern width.
            let previousSum = 0;

            if (y === offset) {
                let value = 0;

                // Adds
                for (let i = addIndex; i < signal.length; i += patternWidth) {
                    for (let j = 0; j < patternRepeat; j++) {
                        if (i + j >= signal.length) {
                            break;
                        }

                        value += signal[i + j];
                    }
                }

                previousSum = value;

                // Subs
                for (let i = subIndex; i < signal.length; i += patternWidth) {
                    for (let j = 0; j < patternRepeat; j++) {
                        if (i + j >= signal.length) {
                            break;
                        }

                        value -= signal[i + j];
                    }
                }

                newSignal[y] = Math.abs(value) % 10;
            } else {
                let value = lastSum;

                // Adds
                for (let i = addIndex; i < signal.length; i += patternWidth) {
                    const from = i;
                    const to = i + patternRepeat;

                    value -= signal[from - 1];

                    if (to - 2 < signal.length) {
                        value += signal[to - 2];
                    }
                    if (to - 1 < signal.length) {
                        value += signal[to - 1];
                    }
                }

                previousSum = value;

                // Subs
                for (let i = subIndex; i < signal.length; i += patternWidth) {
                    for (let j = 0; j < patternRepeat; j++) {
                        if (i + j >= signal.length) {
                            break;
                        }

                        value -= signal[i + j];
                    }
                }

                newSignal[y] = Math.abs(value) % 10;
            }

            lastSum = previousSum;
        }

        signal = newSignal;
    }

    log(signal.slice(offset, offset + 8).join(""));
};
