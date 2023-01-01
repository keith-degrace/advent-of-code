import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

const fromSnafu = (snafu: string): number => {
    let value = 0;

    for (let i = snafu.length - 1; i >= 0; i--) {
        let snafuDigitValue: number;
        if (snafu[i] === "=") {
            snafuDigitValue = -2;
        } else if (snafu[i] === "-") {
            snafuDigitValue = -1;
        } else {
            snafuDigitValue = parseInt(snafu[i]);
        }

        value += snafuDigitValue * Math.pow(5, snafu.length - 1 - i);
    }

    return value;
};

const toSnafuDigit = (number: number): string => {
    if (number == -2) {
        return "=";
    } else if (number == -1) {
        return "-";
    } else {
        return number.toString();
    }
};

const toSnafu = (number: number): string => {
    let snafu: string = "";

    for (let position = 0; ; position++) {
        let remainder = number % 5;

        if (remainder > 2) {
            remainder -= 5;
            number += 5;
        }

        snafu = toSnafuDigit(remainder) + snafu;

        number = Math.floor(number / 5);
        if (number === 0) {
            break;
        }
    }

    return snafu;
};

let sum = 0;
for (const entry of input) {
    sum += fromSnafu(entry);
}

log(toSnafu(sum));
