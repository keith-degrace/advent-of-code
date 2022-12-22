import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";
import { startTimer, stopTimer } from "../utils/timer";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split(/\r\n/);

interface Valve {
    code: string;
    flowRate: number;
    tunnels: Valve[];
    isOpen: boolean;
    canOpen: boolean;
}

const loadValves = (): Valve[] => {
    // First pass, create all the valves.
    const valves: Valve[] = input.map((line) => {
        const match = line.match(/Valve (.+) has flow rate=(.+); tunnels? leads? to valves? (.+)/);

        const code: string = match[1];
        const flowRate: number = Number.parseInt(match[2]);

        return { code, flowRate, tunnels: [], isOpen: false, canOpen: false };
    });

    // Second pass, hook up the tunnels.
    input.forEach((line) => {
        const match = line.match(/Valve (.+) has flow rate=.+; tunnels? leads? to valves? (.+)/);

        const code: string = match[1];
        const tunnels: string[] = match[2].split(", ");

        const valve = valves.find((valve) => valve.code === code);

        for (const tunnel of tunnels) {
            const tunnelValve = valves.find((valve) => valve.code === tunnel);
            valve.tunnels.push(tunnelValve);
        }
    });

    return valves;
};

const getBestPressure = (valves: Valve[], current: Valve, timeLeft: number): number => {
    let memo: Map<string, number> = new Map();
    let openValveState: string = "";
    let openableValves = valves.filter((valve) => valve.canOpen);

    const updateState = () => {
        openValveState = "";

        for (const valve of openableValves) {
            openValveState += valve.code + (valve.isOpen ? "0" : "1");
        }
    };

    const openValve = (valve: Valve) => {
        valve.isOpen = true;
        updateState();
    };

    const closeValve = (valve: Valve) => {
        valve.isOpen = false;
        updateState();
    };

    const iterate = (current: Valve, previous: Valve, timeLeft: number): number => {
        const memoKey = `${timeLeft}${current.code}${openValveState}`;

        let bestPressure = memo.get(memoKey);
        if (bestPressure === undefined) {
            let pressureIfOpen = 0;
            if (!current.isOpen && current.canOpen) {
                openValve(current);

                pressureIfOpen += current.flowRate * (timeLeft - 1);

                if (timeLeft > 1) {
                    pressureIfOpen += iterate(current, current, timeLeft - 1);
                }

                closeValve(current);
            }

            let pressureIfClosed = 0;
            if (timeLeft > 1) {
                for (const tunnel of current.tunnels) {
                    if (tunnel !== previous) {
                        pressureIfClosed = Math.max(pressureIfClosed, iterate(tunnel, current, timeLeft - 1));
                    }
                }
            }

            bestPressure = Math.max(pressureIfOpen, pressureIfClosed);
            memo.set(memoKey, bestPressure);
        }

        return bestPressure;
    };

    updateState();

    return iterate(current, undefined, timeLeft);
};

const getOpenableValvePermutations = (valves: Valve[]): Valve[][] => {
    const permutations: Valve[][] = [];

    const permutate = (valvesToToggle: Valve[], openableValveCodes: Valve[]) => {
        if (valvesToToggle.length === 0) {
            permutations.push(openableValveCodes);
            return;
        }

        permutate(valvesToToggle.slice(1), [valvesToToggle[0], ...openableValveCodes]);
        permutate(valvesToToggle.slice(1), [...openableValveCodes]);
    };

    const valvesWithFlow = valves.filter((valve) => valve.flowRate > 0);
    permutate(valvesWithFlow, []);

    return permutations;
};

interface Solution {
    pressure: number;
    openableValves: Valve[];
}

const getSolutions = (valves: Valve[], openableValvePermutations: Valve[][]): Solution[] => {
    const solutions: Solution[] = [];

    const AA = valves.find((valve) => valve.code === "AA");
    for (const openableValves of openableValvePermutations) {
        // Apply the openable valve permutation.
        valves.forEach((valve) => (valve.canOpen = false));
        openableValves.forEach((valve) => (valve.canOpen = true));

        // Solve the solution
        solutions.push({
            openableValves,
            pressure: getBestPressure(valves, AA, 26),
        });
    }

    return solutions;
};

const getBestPressureWithElephant = () => {
    const openValvesIntersect = (solution1: Solution, solution2: Solution): boolean => {
        for (const valve1 of solution1.openableValves) {
            if (solution2.openableValves.includes(valve1)) {
                return true;
            }
        }

        return false;
    };

    let bestPressure = 0;

    for (let i = 0; i < solutions.length; i++) {
        for (let j = i + 1; j < solutions.length; j++) {
            if (!openValvesIntersect(solutions[i], solutions[j])) {
                bestPressure = Math.max(bestPressure, solutions[i].pressure + solutions[j].pressure);
            }
        }
    }

    return bestPressure;
};

startTimer();

const valves = loadValves();

log(`Permutating openable valves...`);
const openableValvePermutations = getOpenableValvePermutations(valves);

log(`Solving solutions for ${openableValvePermutations.length} permutations of openable valves...`);
log(`... sadly this takes 10-20 minutes.  But hey, at least it works!`);
const solutions = getSolutions(valves, openableValvePermutations);

log(`Finding best pair from ${solutions.length} solutions...`);
const bestPressureWithElephant = getBestPressureWithElephant();

log(`Best Pressure: ${bestPressureWithElephant}`);

stopTimer();
