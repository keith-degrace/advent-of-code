import * as fs from "fs";
import * as path from "path";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim();

export const isMarker = (input: string, position: number, packetLength: number): boolean => {
    for (let i = 0; i < packetLength; i++) {
        for (let j = i + 1; j < packetLength; j++) {
            if (input[position - packetLength + i] === input[position - packetLength + j]) {
                return false;
            }
        }
    }

    return true;
};

export const getMarker = (input: string, packetLength: number): number => {
    for (let i = packetLength; i < input.length; i++) {
        if (isMarker(input, i, packetLength)) {
            return i;
        }
    }

    return -1;
};

console.log(getMarker(input, 4));
