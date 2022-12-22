import * as fs from "fs";
import * as path from "path";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\n");

interface Position {
    x: number;
    y: number;
}

let head: Position = { x: 0, y: 0 };
let tail: Position = { x: 0, y: 0 };

const move = (position: Position, direction: string): Position => {
    switch (direction) {
        case "U":
            return { x: position.x, y: position.y + 1 };
        case "D":
            return { x: position.x, y: position.y - 1 };
        case "L":
            return { x: position.x - 1, y: position.y };
        case "R":
            return { x: position.x + 1, y: position.y };
    }

    throw "huh?";
};

const moveTail = (tail: Position, head: Position): Position => {
    if (head.x - tail.x > 1) {
        return { x: head.x - 1, y: head.y };
    }

    if (tail.x - head.x > 1) {
        return { x: head.x + 1, y: head.y };
    }

    if (head.y - tail.y > 1) {
        return { x: head.x, y: head.y - 1 };
    }

    if (tail.y - head.y > 1) {
        return { x: head.x, y: head.y + 1 };
    }

    return tail;
};

const uniqueTailPositions = new Set<string>();

for (const instruction of input) {
    const direction: string = instruction.split(" ")[0];
    const steps: number = Number.parseInt(instruction.split(" ")[1]);

    for (let step = 0; step < steps; step++) {
        head = move(head, direction);
        tail = moveTail(tail, head);

        uniqueTailPositions.add(`${tail.x},${tail.y}`);
    }
}

console.log(uniqueTailPositions.size);
