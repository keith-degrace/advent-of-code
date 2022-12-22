import * as fs from "fs";
import * as path from "path";

const input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split(/\n/);

interface ElfGroup {
    elf1: string;
    elf2: string;
    elf3: string;
}

const elfGroups: ElfGroup[] = [];
for (let i = 0; i < input.length / 3; i++) {
    elfGroups.push({
        elf1: input[i * 3 + 0],
        elf2: input[i * 3 + 1],
        elf3: input[i * 3 + 2],
    });
}

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

const getBadge = (group: ElfGroup): string => {
    for (const item1 of group.elf1) {
        for (const item2 of group.elf2) {
            for (const item3 of group.elf3) {
                if (item1 === item2 && item1 === item3) {
                    return item1;
                }
            }
        }
    }

    return undefined;
};

let prioritySum: number = 0;
for (const group of elfGroups) {
    const badge: string = getBadge(group);
    if (badge) {
        prioritySum += getPriority(badge);
    }
}

console.log(prioritySum);
