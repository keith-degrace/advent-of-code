import * as fs from "fs";
import * as path from "path";

const input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split(/\n/);

/*

            [J]             [B] [W]
            [T]     [W] [F] [R] [Z]
        [Q] [M]     [J] [R] [W] [H]
    [F] [L] [P]     [R] [N] [Z] [G]
[F] [M] [S] [Q]     [M] [P] [S] [C]
[L] [V] [R] [V] [W] [P] [C] [P] [J]
[M] [Z] [V] [S] [S] [V] [Q] [H] [M]
[W] [B] [H] [F] [L] [F] [J] [V] [B]
 1   2   3   4   5   6   7   8   9 

 */

const stacks: string[][] = [
    "WMLF".split(""),
    "BZVMF".split(""),
    "HVRSLQ".split(""),
    "FSVQPMTJ".split(""),
    "LSW".split(""),
    "FVPMRJW".split(""),
    "JQCPNRF".split(""),
    "VHPSZWRB".split(""),
    "BMJCGHZW".split(""),
];

for (const instruction of input) {
    const parts = instruction.split(" ");

    const count: number = Number.parseInt(parts[1]);
    const from: number = Number.parseInt(parts[3]) - 1;
    const to: number = Number.parseInt(parts[5]) - 1;

    const crates: string[] = [];

    for (let i = 0; i < count; i++) {
        crates.push(stacks[from].pop());
    }

    stacks[to].push(...crates.reverse());
}

let tops: string = "";
for (const stack of stacks) {
    tops += stack.pop();
}

console.log(tops);
