import * as fs from "fs";
import * as path from "path";
import { log } from "../utils";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split(/\r\n/);

interface Valve {
    code: string;
    flowRate: number;
    tunnels: Valve[];
    isOpen: boolean;
}

const valves: Valve[] = input.map((line) => {
    const match = line.match(/Valve (.+) has flow rate=(.+); tunnels? leads? to valves? (.+)/);

    const code: string = match[1];
    const flowRate: number = Number.parseInt(match[2]);

    return { code, flowRate, tunnels: [], isOpen: false };
});

input.forEach((line) => {
    const match = line.match(/Valve (.+) has flow rate=.+; tunnels? leads? to valves? (.+)/);

    const code: string = match[1];
    const tunnels: string[] = match[2].split(", ");

    const valve = valves.find((valve) => valve.code === code);

    for (const tunnel of tunnels) {
        const tunnelValve = valves.find((valve) => valve.code === tunnel);
        valve.tunnels.push(tunnelValve);
    }

    valve.tunnels = valve.tunnels.sort((a, b) => b.flowRate - a.flowRate);
});

let valveState: string = "";

const updateValveState = () => {
    valveState = "";
    for (const valve of valves) {
        if (valve.isOpen) {
            valveState += valve.code;
        }
    }
};

const openValves = (valves: Valve[]) => {
    valves.forEach((valve) => (valve.isOpen = true));
    updateValveState();
};

const closeValves = (valves: Valve[]) => {
    valves.forEach((valve) => (valve.isOpen = false));
    updateValveState();
};

const memo: Record<string, number> = {};

const getBestPressure = (current: Valve, timeLeft: number): number => {
    if (timeLeft === 0) {
        return 0;
    }

    const memoKey = `${timeLeft}-${current.code}-${valveState}`;
    if (memo[memoKey] === undefined) {
        let pressureIfOpen = 0;
        if (!current.isOpen && current.flowRate > 0) {
            openValves([current]);
            pressureIfOpen = current.flowRate * (timeLeft - 1) + getBestPressure(current, timeLeft - 1);
            closeValves([current]);
        }

        let pressureIfClosed = 0;
        for (const tunnel of current.tunnels) {
            pressureIfClosed = Math.max(pressureIfClosed, getBestPressure(tunnel, timeLeft - 1));
        }

        memo[memoKey] = Math.max(pressureIfOpen, pressureIfClosed);
    }

    return memo[memoKey];
};

const AA = valves.find((valve) => valve.code === "AA");

log("Start...");
log(getBestPressure(AA, 30));
