import * as fs from "fs";
import * as path from "path";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\n");

interface Position {
    x: number;
    y: number;
}

const moveHead = (position: Position, direction: string): Position => {
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

const moveKnot = (knot: Position, knotInFront: Position): Position => {
    const deltaX = knotInFront.x - knot.x;
    const deltaY = knotInFront.y - knot.y;

    if (Math.abs(deltaX) < 2 && Math.abs(deltaY) < 2) {
        return knot;
    }

    if (deltaX === 0 && deltaY == 2) {
        return { x: knot.x, y: knot.y + 1 };
    }

    if (deltaX === 0 && deltaY == -2) {
        return { x: knot.x, y: knot.y - 1 };
    }

    if (deltaX === 2 && deltaY == 0) {
        return { x: knot.x + 1, y: knot.y };
    }

    if (deltaX === -2 && deltaY == 0) {
        return { x: knot.x - 1, y: knot.y };
    }

    if (deltaX < 0 && deltaY < 0) {
        return { x: knot.x - 1, y: knot.y - 1 };
    }

    if (deltaX > 0 && deltaY < 0) {
        return { x: knot.x + 1, y: knot.y - 1 };
    }

    if (deltaX < 0 && deltaY > 0) {
        return { x: knot.x - 1, y: knot.y + 1 };
    }

    if (deltaX > 0 && deltaY > 0) {
        return { x: knot.x + 1, y: knot.y + 1 };
    }

    throw `error: ${deltaX}, ${deltaY}`;
};

const uniqueTailPositions = new Set<string>();

let rope: Position[] = [
    { x: 0, y: 0 },
    { x: 0, y: 0 },
    { x: 0, y: 0 },
    { x: 0, y: 0 },
    { x: 0, y: 0 },
    { x: 0, y: 0 },
    { x: 0, y: 0 },
    { x: 0, y: 0 },
    { x: 0, y: 0 },
    { x: 0, y: 0 },
];

for (const instruction of input) {
    const direction: string = instruction.split(" ")[0];
    const steps: number = Number.parseInt(instruction.split(" ")[1]);

    for (let step = 0; step < steps; step++) {
        rope[0] = moveHead(rope[0], direction);

        for (let i = 1; i < rope.length; i++) {
            rope[i] = moveKnot(rope[i], rope[i - 1]);
        }

        const tail = rope[rope.length - 1];
        uniqueTailPositions.add(`${tail.x},${tail.y}`);
    }
}

console.log(uniqueTailPositions.size);
