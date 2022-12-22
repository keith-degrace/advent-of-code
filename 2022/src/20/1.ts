import * as fs from "fs";
import * as path from "path";
import { log } from "../utils";

interface Node {
    value: number;
    nextOriginal: Node;
    previous: Node;
    next: Node;
}

const loadNodes = (): [Node, number] => {
    let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\r\n");

    let root: Node = undefined;
    let previous: Node = undefined;
    for (const line of input) {
        const current = {
            value: Number.parseInt(line),
            nextOriginal: undefined,
            next: undefined,
            previous,
        };

        if (!root) {
            root = current;
        }

        if (previous) {
            previous.nextOriginal = current;
            previous.next = current;
        }

        previous = current;
    }

    previous.next = root;
    root.previous = previous;

    return [root, input.length];
};

const getNodeAt = (node: Node, steps: number): Node => {
    if (steps > 0) {
        while (steps--) {
            node = node.next;
        }
    } else if (steps < 0) {
        while (steps++ <= 0) {
            node = node.previous;
        }
    }

    return node;
};

const insertNode = (node: Node, afterNode: Node) => {
    let newPrevious = afterNode;
    let newNext = afterNode.next;

    node.previous = newPrevious;
    node.next = newNext;

    newPrevious.next = node;
    newNext.previous = node;
};

const removeNode = (node: Node): Node => {
    node.previous.next = node.next;
    node.next.previous = node.previous;

    node.next = undefined;
    node.previous = undefined;

    return node;
};

const findZero = (root: Node): Node => {
    let current = root;

    while (current.value !== 0) {
        current = current.next;
    }

    return current;
};

const mix = (root: Node, size: number) => {
    for (let current = root; current !== undefined; current = current.nextOriginal) {
        const insertAfter = getNodeAt(current, current.value % (size - 1));
        if (insertAfter === current) {
            continue;
        }

        removeNode(current);
        insertNode(current, insertAfter);
    }
};

const getCoordinateSum = (root: Node) => {
    const zero = findZero(root);

    const x = getNodeAt(zero, 1000).value;
    const y = getNodeAt(zero, 2000).value;
    const z = getNodeAt(zero, 3000).value;

    return x + y + z;
};

const solve = () => {
    let [root, size] = loadNodes();

    mix(root, size);

    log(`Sum: ${getCoordinateSum(root)}`);
};

solve();
