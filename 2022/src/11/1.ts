import * as fs from "fs";
import * as path from "path";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

interface Monkey {
    id: number;
    inspected: number;
    items: number[];
    operation: {
        kind: string;
        value: string;
    };
    testValue: number;
    ifTrue: number;
    ifFalse: number;
}

const getMonkeys = (): Record<number, Monkey> => {
    let regex =
        /Monkey ([0-9]+).*\n  Starting items: (.*)\n  Operation: new = old (.) (.+)\n  Test: divisible by ([0-9]+)\n    If true: throw to monkey ([0-9]+)\n    If false: throw to monkey ([0-9]+)/gm;

    const monkeys: Record<number, Monkey> = {};

    for (const match of input.matchAll(regex)) {
        const monkeyId: number = Number.parseInt(match[1]);

        monkeys[monkeyId] = {
            id: monkeyId,
            inspected: 0,
            items: match[2].split(",").map((value) => Number.parseInt(value.trim())),
            operation: {
                kind: match[3],
                value: match[4],
            },
            testValue: Number.parseInt(match[5]),
            ifTrue: Number.parseInt(match[6]),
            ifFalse: Number.parseInt(match[7]),
        };
    }

    return monkeys;
};

const monkeys: Record<number, Monkey> = getMonkeys();

const playRound = () => {
    for (const monkey of Object.values(monkeys)) {
        console.log(`Monkey ${monkey.id}:`);

        for (const item of monkey.items) {
            console.log(`  Monkey inspects an item with a worry level of ${item}.`);
            monkey.inspected++;

            let worry: number = item;

            let value: number = monkey.operation.value === "old" ? worry : Number.parseInt(monkey.operation.value);

            if (monkey.operation.kind === "*") {
                worry *= value;
                console.log(`    Worry level is multiplied by ${value} to ${worry}.`);
            } else if (monkey.operation.kind === "+") {
                worry += value;
                console.log(`    Worry level is increases by ${value} to ${worry}.`);
            } else {
                throw `Unknown operation: ${monkey.operation.kind}`;
            }

            worry = Math.floor(worry / 3);
            console.log(`    Monkey gets bored with item. Worry level is divided by 3 to ${worry}.`);

            if (worry % monkey.testValue === 0) {
                console.log(`    Current worry level is divisible by ${monkey.testValue}.`);
                console.log(`    Item with worry level ${worry} is thrown to monkey ${monkey.ifTrue}`);
                monkeys[monkey.ifTrue].items.push(worry);
            } else {
                console.log(`    Current worry level is not divisible by ${monkey.testValue}.`);
                console.log(`    Item with worry level ${worry} is thrown to monkey ${monkey.ifFalse}`);
                monkeys[monkey.ifFalse].items.push(worry);
            }
        }

        monkey.items = [];
    }
};

for (let i = 0; i < 20; i++) {
    playRound();
}

for (const monkey of Object.values(monkeys)) {
    console.log(`Monkey ${monkey.id}: ${monkey.items}`);
}

for (const monkey of Object.values(monkeys)) {
    console.log(`Monkey ${monkey.id} inspected items ${monkey.inspected} times.`);
}

const inspections: number[] = Object.values(monkeys)
    .map((monkey): number => monkey.inspected)
    .sort((a, b) => b - a);

console.log(inspections);

console.log(`Monkey business is ${inspections[0] * inspections[1]}`);
