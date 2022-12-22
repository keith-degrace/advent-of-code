import * as fs from "fs";
import * as path from "path";
import { __values } from "tslib";
import { log } from "../utils";

type Operator = "+" | "-" | "*" | "/";

interface Operation {
    monkey1: Monkey;
    monkey2: Monkey;
    operator: Operator;
}

interface Monkey {
    name: string;
    job?: Operation;
    value: number;
}

const loadMonkeys = (): Record<string, Monkey> => {
    let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\r\n");

    const monkeys: Record<string, Monkey> = {};
    const jobMonkeys: Record<string, string[]> = {};

    for (const line of input) {
        const parts = line.split(": ");

        let name: string = parts[0];
        let job: Operation;
        let value: number;

        const match = parts[1].match(/(.+) (.) (.+)/);
        if (match) {
            job = {
                monkey1: undefined,
                monkey2: undefined,
                operator: match[2] as Operator,
            };

            jobMonkeys[name] = [match[1], match[3]];
        } else {
            value = parseInt(parts[1]);
        }

        monkeys[name] = { name, job, value };
    }

    for (const monkey of Object.values(monkeys)) {
        const jobMonkey = jobMonkeys[monkey.name];
        if (jobMonkey) {
            monkey.job.monkey1 = monkeys[jobMonkey[0]];
            monkey.job.monkey2 = monkeys[jobMonkey[1]];
        }
    }

    return monkeys;
};

const solveConstants = (monkey: Monkey): void => {
    if (monkey.value !== undefined) {
        return;
    }

    if (monkey.job !== undefined) {
        solveConstants(monkey.job.monkey1);
        solveConstants(monkey.job.monkey2);

        const value1 = monkey.job.monkey1.value;
        const value2 = monkey.job.monkey2.value;

        if (value1 !== undefined && value2 !== undefined) {
            switch (monkey.job.operator) {
                case "+":
                    monkey.value = value1 + value2;
                    break;

                case "-":
                    monkey.value = value1 - value2;
                    break;

                case "*":
                    monkey.value = value1 * value2;
                    break;

                case "/":
                    monkey.value = value1 / value2;
                    break;
            }
        }
    }
};

const expressionToString = (monkey: Monkey): string => {
    if (monkey.value !== undefined) {
        return `${monkey.value}`;
    }

    if (monkey.job) {
        return `(${expressionToString(monkey.job.monkey1)}${monkey.job.operator}${expressionToString(monkey.job.monkey2)})`;
    }

    return "X";
};

const monkeys = loadMonkeys();

const root = monkeys["root"];

monkeys["humn"].job = undefined;
monkeys["humn"].value = undefined;

solveConstants(root);

log("Feed the following equation into a math solver (e.g. Wolfram):");
log(`${expressionToString(root.job.monkey1)}=${expressionToString(root.job.monkey2)}`);
