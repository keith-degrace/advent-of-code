import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

interface Player {
    score: number;
    position: number;
}

const rollDice = (dice: number): number => {
    dice++;
    if (dice > 100) {
        dice = 1;
    }

    return dice;
};

const advance = (position: number, steps: number): number => {
    for (let i = 0; i < steps; i++) {
        position++;
        if (position > 10) {
            position = 1;
        }
    }

    return position;
};

export const run = () => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    const players: Player[] = [
        {
            score: 0,
            position: parseInt(input[0].split(": ")[1]),
        },
        {
            score: 0,
            position: parseInt(input[1].split(": ")[1]),
        },
    ];

    let dice: number = 0;
    let rolls: number = 0;
    let player: number = 0;

    while (true) {
        dice = rollDice(dice);
        rolls++;
        players[player].position = advance(players[player].position, dice);

        dice = rollDice(dice);
        rolls++;
        players[player].position = advance(players[player].position, dice);

        dice = rollDice(dice);
        rolls++;
        players[player].position = advance(players[player].position, dice);

        players[player].score += players[player].position;
        if (players[player].score >= 1000) {
            break;
        }

        player = player == 1 ? 0 : 1;
    }

    let loser = player === 1 ? 0 : 1;

    log(rolls * players[loser].score);
};
