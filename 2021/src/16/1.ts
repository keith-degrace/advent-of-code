import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

export const run = () => {
    const hexValue: string = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

    let versionSum = 0;

    const parsePacket = (bits: string, index: number): number => {
        const version = parseInt(bits.slice(index, index + 3), 2);
        index += 3;

        versionSum += version;

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
        } else {
            const lengthType = parseInt(bits[index++], 2);

            if (lengthType === 0) {
                const subPacketLength = parseInt(bits.slice(index, index + 15), 2);
                index += 15;

                const end = index + subPacketLength;
                while (index < end) {
                    index = parsePacket(bits, index);
                }
            } else if (lengthType === 1) {
                const subPacketCount = parseInt(bits.slice(index, index + 11), 2);
                index += 11;

                for (let i = 0; i < subPacketCount; i++) {
                    index = parsePacket(bits, index);
                }
            } else {
                throw `Invalid Length Type: ${lengthType}`;
            }
        }

        return index;
    };

    let bits = "";
    for (const char of hexValue) {
        bits += Number.parseInt(char, 16).toString(2).padStart(4, "0");
    }

    parsePacket(bits, 0);

    log(versionSum);
};
