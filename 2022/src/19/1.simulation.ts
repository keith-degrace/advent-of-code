import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

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
    timeLeft: number;

    ore: number;
    clay: number;
    obsidian: number;
    geode: number;

    oreRobots: number;
    clayRobots: number;
    obsidianRobots: number;
    geodeRobots: number;
}

const memo: Record<string, Resources> = {};

const getBestGeodeCount = (blueprint: Blueprint, resources: Resources): Resources => {
    const memoKey = `${blueprint.id}-${JSON.stringify(resources)}`;
    // log(memoKey);
    if (memo[memoKey] !== undefined) {
        return memo[memoKey];
    }

    const newResources: Resources = {
        ...resources,
        timeLeft: resources.timeLeft - 1,
        ore: resources.ore + resources.oreRobots,
        clay: resources.clay + resources.clayRobots,
        obsidian: resources.obsidian + resources.obsidianRobots,
        geode: resources.geode + resources.geodeRobots,
    };

    // log(`[${24 - timeLeft + 1}] ${JSON.stringify(newResources)}`);

    let bestResources: Resources = newResources;

    if (newResources.timeLeft > 0) {
        // Create a geode robot.
        if (resources.ore >= blueprint.geodeRobotRecipe.ore && resources.obsidian >= blueprint.geodeRobotRecipe.obsidian) {
            // log(`[${24 - timeLeft + 1}] Creating geode robot`);
            const candidateResources = getBestGeodeCount(blueprint, {
                ...newResources,
                ore: newResources.ore - blueprint.geodeRobotRecipe.ore,
                obsidian: newResources.obsidian - blueprint.geodeRobotRecipe.obsidian,
                geodeRobots: newResources.geodeRobots + 1,
            });

            if (!bestResources || candidateResources.geode > bestResources.geode) {
                bestResources = candidateResources;
            }
        }

        // Create an obsidian robot.
        if (resources.ore >= blueprint.obsidianRobotRecipe.ore && resources.clay >= blueprint.obsidianRobotRecipe.clay) {
            // log(`[${24 - timeLeft + 1}] Creating obsidian robot`);
            const candidateResources = getBestGeodeCount(blueprint, {
                ...newResources,
                ore: newResources.ore - blueprint.obsidianRobotRecipe.ore,
                clay: newResources.clay - blueprint.obsidianRobotRecipe.clay,
                obsidianRobots: newResources.obsidianRobots + 1,
            });

            if (!bestResources || candidateResources.geode > bestResources.geode) {
                bestResources = candidateResources;
            }
        }

        // Create a clay robot.
        if (resources.ore >= blueprint.clayRobotRecipe.ore) {
            // log(`[${24 - timeLeft + 1}] Creating clay robot`);
            const candidateResources = getBestGeodeCount(blueprint, {
                ...newResources,
                ore: newResources.ore - blueprint.clayRobotRecipe.ore,
                clayRobots: newResources.clayRobots + 1,
            });

            if (!bestResources || candidateResources.geode > bestResources.geode) {
                bestResources = candidateResources;
            }
        }
        // Create an ore robot.
        if (resources.ore >= blueprint.oreRobotRecipe.ore) {
            // log(`[${24 - timeLeft + 1}] Creating ore robot`);
            const candidateResources = getBestGeodeCount(blueprint, {
                ...newResources,
                ore: newResources.ore - blueprint.oreRobotRecipe.ore,
                oreRobots: newResources.oreRobots + 1,
            });

            if (!bestResources || candidateResources.geode > bestResources.geode) {
                bestResources = candidateResources;
            }
        }

        // Don't create a robot
        {
            // log(`[${24 - timeLeft + 1}] Create nothing`);

            const candidateResources = getBestGeodeCount(blueprint, newResources);

            if (!bestResources || candidateResources.geode > bestResources.geode) {
                bestResources = candidateResources;
            }
        }
    }

    if (bestResources.geode > 10) {
        log(JSON.stringify(bestResources));
    }

    memo[memoKey] = bestResources;

    return bestResources;
};

const blueprints = loadBlueprints();

const initialResource: Resources = {
    timeLeft: 24,
    ore: 0,
    clay: 0,
    obsidian: 0,
    geode: 0,
    oreRobots: 1,
    clayRobots: 0,
    obsidianRobots: 0,
    geodeRobots: 0,
};

let qualitySum: number = 0;

for (const blueprint of blueprints) {
    const resources: Resources = getBestGeodeCount(blueprint, initialResource);
    log(JSON.stringify(resources, null, 2));

    const quality = blueprint.id * resources.geode;
    qualitySum += quality;
}

log(qualitySum);
