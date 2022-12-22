import * as fs from "fs";
import * as path from "path";
import { getTsBuildInfoEmitOutputFilePath, isTemplateSpan } from "typescript";

let input = fs
    .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
    .trim()
    .split("\n")
    .filter((value) => value.length !== 0)
    .map(eval);

const isInOrder = (a, b): boolean => {
    const aIsNumber: boolean = typeof a === "number";
    const bIsNumber: boolean = typeof b === "number";

    if (aIsNumber && bIsNumber) {
        if (a < b) {
            return true;
        } else if (a > b) {
            return false;
        }

        return undefined;
    }

    if (aIsNumber && !bIsNumber) {
        return isInOrder([a], b);
    }

    if (!aIsNumber && bIsNumber) {
        return isInOrder(a, [b]);
    }

    let index: number = 0;

    while (true) {
        if (index === a.length && index === b.length) {
            return undefined;
        }

        if (index === a.length) {
            return true;
        }

        if (index === b.length) {
            return false;
        }

        const inOrder: boolean = isInOrder(a[index], b[index]);
        if (inOrder !== undefined) {
            return inOrder;
        }

        index++;
    }
};

input.push([[2]]);
input.push([[6]]);

input.sort((a, b) => (isInOrder(a, b) ? -1 : 1));

const marker1 = input.findIndex((value) => JSON.stringify(value) === "[[2]]") + 1;
const marker2 = input.findIndex((value) => JSON.stringify(value) === "[[6]]") + 1;

console.log(marker1 * marker2);
