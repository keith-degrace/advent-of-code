import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

const load = (): [string, Record<string, string>] => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    const template = input[0];

    const rules: Record<string, string> = {};
    for (let i = 2; i < input.length; i++) {
        const parts = input[i].split(" -> ");
        rules[parts[0]] = parts[1];
    }

    return [template, rules];
};

const evolve = (template: string, rules: Record<string, string>): string => {
    let nextTemplate = "";

    for (let i = 0; i < template.length - 1; i++) {
        const char1 = template[i];
        const char2 = template[i + 1];

        nextTemplate += char1 + rules[`${char1}${char2}`];
    }
    nextTemplate += template[template.length - 1];

    return nextTemplate;
};

const memo: Record<string, Record<string, number>> = {};

const getCounts = (template: string, steps: number, rules: Record<string, string>): Record<string, number> => {
    const memoKey = `${template}-${steps}`;
    if (memo[memoKey]) {
        return { ...memo[memoKey] };
    }

    let counts: Record<string, number> = {};

    if (steps > 0) {
        const nextTemplate = evolve(template, rules);

        for (let i = 0; i < nextTemplate.length - 1; i++) {
            const char1 = nextTemplate[i];
            const char2 = nextTemplate[i + 1];

            const subCounts = getCounts(`${char1}${char2}`, steps - 1, rules);
            if (i !== 0) {
                subCounts[char1]--;
            }

            for (const entry of Object.entries(subCounts)) {
                const char = entry[0];
                let count = entry[1];

                counts[char] = (counts[char] ?? 0) + count;
            }
        }
    } else {
        for (const char of template) {
            counts[char] = (counts[char] ?? 0) + 1;
        }
    }

    memo[memoKey] = counts;
    return { ...counts };
};

export const run = () => {
    let [template, rules] = load();

    const counts = getCounts(template, 40, rules);

    const leastCommon = Math.min(...Object.values(counts));
    const mostCommon = Math.max(...Object.values(counts));

    log(mostCommon - leastCommon);
};
