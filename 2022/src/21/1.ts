import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

interface Operation {
    monkey1: string;
    monkey2: string;
    operation: "+" | "-" | "*" | "/";
}

interface Monkey {
    name: string;
    job: number | Operation;
}

const loadMonkeys = (): Record<string, Monkey> => {
    let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\r\n");

    const monkeys = {};
    for (const line of input) {
        const parts = line.split(": ");

        const monkey = {
            name: parts[0],
            job: undefined,
        };

        const match = parts[1].match(/(.+) (.) (.+)/);
        if (match) {
            monkey.job = {
                monkey1: match[1],
                monkey2: match[3],
                operation: match[2],
            };
        } else {
            monkey.job = parseInt(parts[1]);
        }

        monkeys[monkey.name] = monkey;
    }

    return monkeys;
};

const monkeys = loadMonkeys();

for (let pass = 0; typeof monkeys["root"].job !== "number"; pass++) {
    for (const monkey of Object.values(monkeys)) {
        if (typeof monkey.job !== "number") {
            const monkey1: Monkey = monkeys[monkey.job.monkey1];
            const monkey2: Monkey = monkeys[monkey.job.monkey2];

            if (typeof monkey1.job === "number" && typeof monkey2.job === "number") {
                const value1: number = monkey1.job;
                const value2: number = monkey2.job;

                switch (monkey.job.operation) {
                    case "+":
                        monkey.job = value1 + value2;
                        break;

                    case "-":
                        monkey.job = value1 - value2;
                        break;

                    case "*":
                        monkey.job = value1 * value2;
                        break;

                    case "/":
                        monkey.job = value1 / value2;
                        break;
                }
            }
        }
    }
}

log(monkeys["root"].job);
