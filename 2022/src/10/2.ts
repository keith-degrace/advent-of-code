import * as fs from "fs";
import * as path from "path";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\n");

let crt: string[] = [" ".repeat(40), " ".repeat(40), " ".repeat(40), " ".repeat(40), " ".repeat(40), " ".repeat(40)];

let register: number = 1;
let cycle: number = 1;

const nextCycle = () => {
    const x: number = (cycle - 1) % 40;
    const y: number = Math.floor((cycle - 1) / 40);

    const spriteStart: number = register - 1;
    const spriteEnd: number = register + 1;

    if (x >= spriteStart && x <= spriteEnd) {
        crt[y] = crt[y].substr(0, x) + "|" + crt[y].substr(x + 1);
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

console.log(crt.join("\n"));
