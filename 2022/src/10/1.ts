import * as fs from "fs";
import * as path from "path";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\n");

let register: number = 1;
let cycle: number = 1;
let strengthSum: number = 0;

const nextCycle = () => {
    if (cycle === 20 || (cycle > 20 && (cycle - 20) % 40 === 0)) {
        strengthSum += cycle * register;
    }

    cycle++;
};

for (const instruction of input) {
    if (instruction === "noop") {
        nextCycle();
    } else if (instruction.startsWith("addx")) {
        nextCycle();
        nextCycle();

        const value: number = Number.parseInt(instruction.split(" ")[1]);
        register += value;
    }
}

console.log(strengthSum);
