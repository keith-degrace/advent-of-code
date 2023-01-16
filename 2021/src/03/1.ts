import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

export const run = () => {
    let values = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\n");

    let bitCount = values[0].length;
    let gammaRate = "";
    let epsilonRate = "";

    for (let position = 0; position < bitCount; position++) {
        let ones = 0;
        let zeroes = 0;

        for (const value of values) {
            if (value[position] === "0") {
                ones++;
            } else {
                zeroes++;
            }
        }

        if (ones > zeroes) {
            gammaRate += "1";
            epsilonRate += "0";
        } else {
            gammaRate += "0";
            epsilonRate += "1";
        }
    }

    log(parseInt(gammaRate, 2) * parseInt(epsilonRate, 2));
};
