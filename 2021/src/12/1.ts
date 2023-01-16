import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

const load = (): Record<string, string[]> => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    const pathes: Record<string, string[]> = {};
    for (const line of input) {
        const parts = line.split("-");

        pathes[parts[0]] = [...(pathes[parts[0]] ?? []), parts[1]];
        pathes[parts[1]] = [...(pathes[parts[1]] ?? []), parts[0]];
    }

    return pathes;
};

const getPathCount = (pathes: Record<string, string[]>, from: string, to: string, visited: string[]): number => {
    if (from === to) {
        return 1;
    }

    let count = 0;

    for (const option of pathes[from]) {
        if (option === option.toUpperCase() || !visited.includes(option)) {
            count += getPathCount(pathes, option, to, [...visited, from]);
        }
    }

    return count;
};

export const run = () => {
    const pathes = load();

    const count = getPathCount(pathes, "start", "end", []);
    log(count);
};
