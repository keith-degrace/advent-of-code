import * as fs from "fs";
import * as path from "path";
import { resourceUsage } from "process";
import { log } from "../utils";

interface Blueprint {
    id: number;
    oreRobotRecipe: {
        ore: number;
    };
    clayRobotRecipe: {
        ore: number;
    };
    obsidianRobotRecipe: {
        ore: number;
        clay: number;
    };
    geodeRobotRecipe: {
        ore: number;
        obsidian: number;
    };
}

const loadBlueprints = (): Blueprint[] => {
    let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim();

    const regex =
        /Blueprint ([0-9]+):\r?\n? *Each ore robot costs ([0-9]+) ore.\r?\n? *Each clay robot costs ([0-9]+) ore.\r?\n? *Each obsidian robot costs ([0-9]+) ore and ([0-9]+) clay.\r?\n? *Each geode robot costs ([0-9]+) ore and ([0-9]+) obsidian.\r?\n?/gm;

    const blueprints: Blueprint[] = [];
    for (const match of input.matchAll(regex)) {
        blueprints.push({
            id: Number.parseInt(match[1]),
            oreRobotRecipe: {
                ore: Number.parseInt(match[2]),
            },
            clayRobotRecipe: {
                ore: Number.parseInt(match[3]),
            },
            obsidianRobotRecipe: {
                ore: Number.parseInt(match[4]),
                clay: Number.parseInt(match[5]),
            },
            geodeRobotRecipe: {
                ore: Number.parseInt(match[6]),
                obsidian: Number.parseInt(match[7]),
            },
        });
    }

    return blueprints;
};

interface Resources {
    time: number;

    ore: number;
    clay: number;
    obsidian: number;
    geode: number;

    oreRobots: number;
    clayRobots: number;
    obsidianRobots: number;
    geodeRobots: number;
}

const blueprints = loadBlueprints();

const initialResource: Resources = {
    time: 1,

    ore: 0,
    clay: 0,
    obsidian: 0,
    geode: 0,

    oreRobots: 1,
    clayRobots: 0,
    obsidianRobots: 0,
    geodeRobots: 0,
};

const blueprint = blueprints[0];

let best: Resources = { ...initialResource };

for (let maxOreRobots = 0; maxOreRobots < 10; maxOreRobots++) {
    for (let maxClayRobots = 0; maxClayRobots < 10; maxClayRobots++) {
        for (let maxObsidianRobots = 0; maxObsidianRobots < 10; maxObsidianRobots++) {
            const resources = { ...initialResource };

            while (resources.time <= 24) {
                if (resources.ore >= blueprint.geodeRobotRecipe.ore && resources.obsidian >= blueprint.geodeRobotRecipe.obsidian) {
                    resources.ore -= blueprint.geodeRobotRecipe.ore;
                    resources.obsidian -= blueprint.geodeRobotRecipe.obsidian;
                    resources.geodeRobots++;
                    resources.geode--;
                } else if (
                    resources.obsidianRobots < maxObsidianRobots &&
                    resources.ore >= blueprint.obsidianRobotRecipe.ore &&
                    resources.clay >= blueprint.obsidianRobotRecipe.clay
                ) {
                    resources.ore -= blueprint.obsidianRobotRecipe.ore;
                    resources.clay -= blueprint.obsidianRobotRecipe.clay;
                    resources.obsidianRobots++;
                    resources.obsidian--;
                } else if (resources.clayRobots < maxClayRobots && resources.ore >= blueprint.clayRobotRecipe.ore) {
                    resources.ore -= blueprint.clayRobotRecipe.ore;
                    resources.clayRobots++;
                    resources.clay--;
                } else if (resources.oreRobots < maxOreRobots && resources.ore >= blueprint.oreRobotRecipe.ore) {
                    resources.ore -= blueprint.oreRobotRecipe.ore;
                    resources.oreRobots++;
                    resources.ore--;
                }

                resources.ore += resources.oreRobots;
                resources.clay += resources.clayRobots;
                resources.obsidian += resources.obsidianRobots;
                resources.geode += resources.geodeRobots;

                resources.time++;
            }

            if (best.geode < resources.geode) {
                best = resources;
            }

            log(`[${maxOreRobots},${maxClayRobots},${maxObsidianRobots}] ${resources.geode}`);
        }
    }
}

log(JSON.stringify(best, null, 2));
