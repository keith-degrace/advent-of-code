import * as fs from "fs";
import * as path from "path";

const input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split(/\n/);

interface Rucksack {
    firstCompartment: string;
    secondCompartment: string;
}

const rucksacks: Rucksack[] = input.map(
    (contents: string): Rucksack => ({
        firstCompartment: contents.slice(0, contents.length / 2),
        secondCompartment: contents.slice(contents.length / 2),
    })
);

const getPriority = (item: string): number => {
    const charCode = item.charCodeAt(0);

    if (charCode >= 97) {
        // a-z = 97-122
        return charCode - 96;
    } else {
        // A-Z = 65-90
        return charCode - 38;
    }

    return charCode;
};

const getCommonItem = (rucksack: Rucksack): string => {
    for (const item1 of rucksack.firstCompartment) {
        for (const item2 of rucksack.secondCompartment) {
            if (item1 === item2) {
                return item1;
            }
        }
    }

    return undefined;
};

let prioritySum: number = 0;
for (const rucksack of rucksacks) {
    const commonItem: string = getCommonItem(rucksack);
    if (commonItem) {
        prioritySum += getPriority(commonItem);
    }
}

console.log(prioritySum);
