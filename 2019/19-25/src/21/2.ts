import * as fs from "fs";
import * as path from "path";
import { runSimpleIntCode } from "../utils/intcode";
import { log } from "../utils/log";

const program = fs
    .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
    .split(",")
    .map((value) => parseInt(value));

const runProgram = (instructions: string[]): string | number => {
    // Convert text instructions to ASCII.
    let input: number[] = [];
    for (const instruction of instructions) {
        for (let i = 0; i < instruction.length; i++) {
            input.push(instruction.codePointAt(i));
        }

        input.push(10);
    }

    let output: number[] = runSimpleIntCode({ ...program }, input);

    // Convert ASCII output to text.
    try {
        return String.fromCodePoint(...output);
    } catch {
        // This happens if output contains "a single giant integer outside the normal ASCII range"
        return output[output.length - 1];
    }
};

export const run = () => {
    // J = (!A & D) | (!B && D) | (!C && D && H)
    const script = `
        NOT A J
        AND D J
        NOT B T
        AND D T
         OR T J
        NOT C T
        AND D T
        AND H T
         OR T J
        `
        .split("\n")
        .map((value) => value.trim())
        .filter((value) => value.length > 0);

    const result = runProgram([...script, "RUN"]);
    log(result);
};
