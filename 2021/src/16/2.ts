import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

export const run = () => {
    const hexValue: string = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

    let index = 0;

    const parsePacket = (bits: string): number => {
        const version = parseInt(bits.slice(index, index + 3), 2);
        index += 3;

        const type = parseInt(bits.slice(index, index + 3), 2);
        index += 3;

        if (type === 4) {
            let numberBits = "";
            while (true) {
                const lastGroup = bits[index++] === "0";

                numberBits += bits.slice(index, index + 4);
                index += 4;

                if (lastGroup) {
                    break;
                }
            }

            return parseInt(numberBits, 2);
        } else {
            const lengthType = parseInt(bits[index++], 2);

            let values: number[] = [];
            if (lengthType === 0) {
                const subPacketLength = parseInt(bits.slice(index, index + 15), 2);
                index += 15;

                const end = index + subPacketLength;
                while (index < end) {
                    values.push(parsePacket(bits));
                }
            } else if (lengthType === 1) {
                const subPacketCount = parseInt(bits.slice(index, index + 11), 2);
                index += 11;

                for (let i = 0; i < subPacketCount; i++) {
                    values.push(parsePacket(bits));
                }
            } else {
                throw `Invalid Length Type: ${lengthType}`;
            }

            switch (type) {
                case 0:
                    return values.reduce((a, b) => a + b, 0);
                case 1:
                    return values.reduce((a, b) => a * b, 1);
                case 2:
                    return Math.min(...values);
                case 3:
                    return Math.max(...values);
                case 5:
                    return values[0] > values[1] ? 1 : 0;
                case 6:
                    return values[0] < values[1] ? 1 : 0;
                case 7:
                    return values[0] === values[1] ? 1 : 0;
                default:
                    throw `Unknown type: ${type}`;
            }
        }
    };

    let bits = "";
    for (const char of hexValue) {
        bits += Number.parseInt(char, 16).toString(2).padStart(4, "0");
    }

    const result = parsePacket(bits);
    log(result);
};
