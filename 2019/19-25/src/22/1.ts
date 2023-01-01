import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

export const run = () => {
    const input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    let deck: number[] = [];
    for (let i = 0; i < 10007; i++) {
        deck.push(i);
    }

    log(deck);

    for (const instruction of input) {
        if (instruction === "deal into new stack") {
            deck = deck.reverse();
        } else if (instruction.startsWith("cut")) {
            const size = parseInt(instruction.split(" ")[1]);
            deck = [...deck.slice(size, deck.length), ...deck.slice(0, size)];
        } else if (instruction.startsWith("deal with")) {
            const increment = parseInt(instruction.split(" ")[3]);

            let newDeck = [...deck];
            let position = 0;
            for (let i = 0; i < deck.length; i++) {
                newDeck[position] = deck[i];
                position = (position + increment) % deck.length;
            }
            deck = newDeck;
        }
    }

    const position = deck.findIndex((card) => card === 2019);
    log(position);
};
