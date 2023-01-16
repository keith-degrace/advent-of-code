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

export const run = () => {
    let [template, rules] = load();

    for (let steps = 0; steps < 10; steps++) {
        let nextTemplate = "";
        for (let i = 0; i < template.length - 1; i++) {
            const char1 = template[i];
            const char2 = template[i + 1];

            nextTemplate += char1 + rules[`${char1}${char2}`];
        }
        nextTemplate += template[template.length - 1];

        template = nextTemplate;
    }

    const counts: Record<string, number> = {};
    for (const char of template) {
        counts[char] = (counts[char] ?? 0) + 1;
    }

    const leastCommon = Math.min(...Object.values(counts));
    const mostCommon = Math.max(...Object.values(counts));

    log(mostCommon - leastCommon);
};
