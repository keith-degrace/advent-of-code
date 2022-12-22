import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";
import { startTimer, stopTimer } from "../utils/timer";

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

const getUpperBoundOnGeode = (blueprint: Blueprint, resources: Resources): number => {
    let ore = resources.ore;
    let obsidian = resources.obsidian;
    let geode = resources.geode;

    let oreRobots = resources.oreRobots;
    let obsidianRobots = resources.obsidianRobots;
    let geodeRobots = resources.geodeRobots;

    let timeLeft = resources.timeLeft;

    while (timeLeft >= 0) {
        ore += oreRobots;
        obsidian += obsidianRobots;
        geode += geodeRobots;

        // Assume we can create one of each basic robot per minute
        oreRobots++;
        obsidianRobots++;

        // Let's actually build the geode robot though.
        if (ore >= blueprint.geodeRobotRecipe.ore && obsidian >= blueprint.geodeRobotRecipe.obsidian) {
            geodeRobots++;
            ore -= blueprint.geodeRobotRecipe.ore;
            obsidian -= blueprint.geodeRobotRecipe.obsidian;
        }

        timeLeft--;
    }

    return geode;
};

const isBetter = (resources1: Resources, resources2: Resources): boolean => {
    if (!resources2) {
        return true;
    }

    return resources1.geode > resources2.geode;
};

const getTimeToCollectResourcesForOreRobot = (blueprint: Blueprint, resources: Resources): number => {
    const missingOre: number = blueprint.oreRobotRecipe.ore - resources.ore;
    if (missingOre <= 0) {
        return undefined;
    }

    const minutesToCreate: number = Math.ceil(missingOre / resources.oreRobots);
    return minutesToCreate;
};

const getTimeToCollectResourcesForClayRobot = (blueprint: Blueprint, resources: Resources): number => {
    const missingOre: number = blueprint.clayRobotRecipe.ore - resources.ore;
    if (missingOre <= 0) {
        return undefined;
    }

    const minutesToCreate: number = Math.ceil(missingOre / resources.oreRobots);
    return minutesToCreate;
};

const getTimeToCollectResourcesForObsidianRobot = (blueprint: Blueprint, resources: Resources): number => {
    const missingOre: number = Math.max(0, blueprint.obsidianRobotRecipe.ore - resources.ore);
    const missingClay: number = Math.max(0, blueprint.obsidianRobotRecipe.clay - resources.clay);

    if (missingOre <= 0 && missingClay <= 0) {
        return undefined;
    }

    const minutesToCreateOre: number = Math.ceil(missingOre / resources.oreRobots);
    const minutesToCreateClay: number = Math.ceil(missingClay / resources.clayRobots);

    const minutesToCreate: number = Math.max(minutesToCreateOre, minutesToCreateClay);
    return minutesToCreate;
};

const getTimeToCollectResourcesForGeodeRobot = (blueprint: Blueprint, resources: Resources): number => {
    const missingOre: number = Math.max(0, blueprint.geodeRobotRecipe.ore - resources.ore);
    const missingObsidian: number = Math.max(0, blueprint.geodeRobotRecipe.obsidian - resources.obsidian);

    if (missingOre <= 0 && missingObsidian <= 0) {
        return undefined;
    }

    const minutesToCreateOre: number = Math.ceil(missingOre / resources.oreRobots);
    const minutesToCreateObsidian: number = Math.ceil(missingObsidian / resources.obsidianRobots);

    const minutesToCreate: number = Math.max(minutesToCreateOre, minutesToCreateObsidian);
    return minutesToCreate;
};

const isWorthCreatingGeodeRobot = (blueprint: Blueprint, resources: Resources): boolean => {
    //  - With 2 minutes left, create geode robot
    //  - With 1 minutes left, create geode
    if (resources.timeLeft < 2) {
        // log("Skipping obsidian robot... not enough time to create the robot, an obsidian, a new geode robot, and finally the geode ");
        return false;
    }

    return true;
};

const isWorthCreatingObsidianRobot = (blueprint: Blueprint, resources: Resources): boolean => {
    //  - With 3 minutes left, create obsidian robot
    //  - With 2 minutes left, create obsidian
    //  - With 1 minutes left, create geode robot
    //  - With 0 minutes left, create geode
    if (resources.timeLeft < 3) {
        // log("Skipping obsidian robot... not enough time to create the robot, an obsidian, a new geode robot, and finally the geode ");
        return false;
    }

    // Can I add enough obsidian to create a geode robot, and then a geode?
    const projectedObsidianFromNewRobot = resources.timeLeft - 1;
    if (projectedObsidianFromNewRobot === 0) {
        return false;
    }

    const projectedObsidianWithoutNewRobot = resources.obsidianRobots * resources.timeLeft;
    const projectedObsidianWithNewRobot = projectedObsidianWithoutNewRobot + projectedObsidianFromNewRobot;

    const projectedGeodeRobotWithoutNewRobot = Math.floor(projectedObsidianWithoutNewRobot / blueprint.geodeRobotRecipe.obsidian);
    const projectedGeodeRobotWithNewRobot = Math.floor(projectedObsidianWithNewRobot / blueprint.geodeRobotRecipe.obsidian);

    if (projectedGeodeRobotWithNewRobot < projectedGeodeRobotWithoutNewRobot) {
        log("Skipping obsidian robot... not enough new obsidian to create the new geode robot in time");
        return false;
    }

    return true;
};

const isWorthCreatingClayRobot = (blueprint: Blueprint, resources: Resources): boolean => {
    //  - With 5 minutes left, create clay robot
    //  - With 4 minutes left, create clay
    //  - With 3 minutes left, create obsidian robot
    //  - With 2 minutes left, create obsidian
    //  - With 1 minutes left, create geode robot
    //  - With 0 minutes left, create geode
    if (resources.timeLeft < 5) {
        // log("Skipping obsidian robot... not enough time to create the robot, an obsidian, a new geode robot, and finally the geode ");
        return false;
    }

    const projectedClayFromNewRobot = resources.timeLeft - 1;
    if (projectedClayFromNewRobot === 0) {
        return false;
    }

    const projectedClayWithoutNewRobot = resources.clayRobots * resources.timeLeft;
    const projectedClayWithNewRobot = projectedClayWithoutNewRobot + projectedClayFromNewRobot;

    const projectedObsidianRobotWithoutNewRobot = Math.floor(projectedClayWithoutNewRobot / blueprint.obsidianRobotRecipe.clay);
    const projectedObsidianRobotWithNewRobot = Math.floor(projectedClayWithNewRobot / blueprint.obsidianRobotRecipe.clay);

    if (projectedObsidianRobotWithNewRobot < projectedObsidianRobotWithoutNewRobot) {
        log("Skipping clay robot... not enough new clay to create the new obsidian robot in time");
        return false;
    }

    return true;
};

const isWorthCreatingOreRobot = (blueprint: Blueprint, resources: Resources): boolean => {
    //  - With 3 minutes left, create ore robot
    //  - With 2 minutes left, create ore
    //  - With 1 minute left, create geode robot
    //  - With 0 minute left, create geode
    if (resources.timeLeft < 2) {
        // log("Skipping obsidian robot... not enough time to create the robot, an obsidian, a new geode robot, and finally the geode ");
        return false;
    }

    return true;
};

const getBestGeodeCount = (blueprint: Blueprint, resources: Resources): Resources => {
    const memo: Map<string, Resources> = new Map();
    let bestSoFar: number = 0;

    const iterate = (resources: Resources) => {
        const memoKey = `${resources.timeLeft}-${resources.ore}-${resources.clay}-${resources.obsidian}-${resources.geode}-${resources.oreRobots}-${resources.clayRobots}-${resources.obsidianRobots}-${resources.geodeRobots}`;
        if (memo.has(memoKey)) {
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

        if (newResources.timeLeft === 0) {
            return newResources;
        }

        if (getUpperBoundOnGeode(blueprint, newResources) < bestSoFar) {
            return newResources;
        }

        let bestOption: Resources = newResources;

        // Option 1: Create an geode robot.
        if (resources.obsidianRobots > 0 && isWorthCreatingGeodeRobot(blueprint, resources)) {
            const optionResources = { ...newResources };

            const timeToCreateResources = getTimeToCollectResourcesForGeodeRobot(blueprint, resources);
            if (timeToCreateResources !== undefined) {
                optionResources.ore += timeToCreateResources * resources.oreRobots;
                optionResources.clay += timeToCreateResources * resources.clayRobots;
                optionResources.obsidian += timeToCreateResources * resources.obsidianRobots;
                optionResources.geode += timeToCreateResources * resources.geodeRobots;
                optionResources.timeLeft -= timeToCreateResources;
            }

            if (optionResources.timeLeft > 0) {
                optionResources.ore -= blueprint.geodeRobotRecipe.ore;
                optionResources.obsidian -= blueprint.geodeRobotRecipe.obsidian;
                optionResources.geodeRobots++;

                const candidateResources = iterate(optionResources);
                if (isBetter(candidateResources, bestOption)) {
                    bestOption = candidateResources;
                }
            }
        }

        // Option 2: Create an obsidian robot.
        if (resources.clayRobots > 0 && isWorthCreatingObsidianRobot(blueprint, resources)) {
            const optionResources = { ...newResources };

            const timeToCreateResources = getTimeToCollectResourcesForObsidianRobot(blueprint, resources);
            if (timeToCreateResources !== undefined) {
                optionResources.ore += timeToCreateResources * resources.oreRobots;
                optionResources.clay += timeToCreateResources * resources.clayRobots;
                optionResources.obsidian += timeToCreateResources * resources.obsidianRobots;
                optionResources.geode += timeToCreateResources * resources.geodeRobots;
                optionResources.timeLeft -= timeToCreateResources;
            }

            if (optionResources.timeLeft > 0) {
                optionResources.ore -= blueprint.obsidianRobotRecipe.ore;
                optionResources.clay -= blueprint.obsidianRobotRecipe.clay;
                optionResources.obsidianRobots++;

                const candidateResources = iterate(optionResources);
                if (isBetter(candidateResources, bestOption)) {
                    bestOption = candidateResources;
                }
            }
        }

        // Option 3: Create a clay robot.
        if (isWorthCreatingClayRobot(blueprint, resources)) {
            const optionResources = { ...newResources };

            const timeToCreateResources = getTimeToCollectResourcesForClayRobot(blueprint, resources);
            if (timeToCreateResources !== undefined) {
                optionResources.ore += timeToCreateResources * resources.oreRobots;
                optionResources.clay += timeToCreateResources * resources.clayRobots;
                optionResources.obsidian += timeToCreateResources * resources.obsidianRobots;
                optionResources.geode += timeToCreateResources * resources.geodeRobots;
                optionResources.timeLeft -= timeToCreateResources;
            }

            if (optionResources.timeLeft > 0) {
                optionResources.ore -= blueprint.clayRobotRecipe.ore;
                optionResources.clayRobots++;

                const candidateResources = iterate(optionResources);
                if (isBetter(candidateResources, bestOption)) {
                    bestOption = candidateResources;
                }
            }
        }

        // Option 4: Create an ore robot.
        if (isWorthCreatingOreRobot(blueprint, resources)) {
            const optionResources = { ...newResources };

            const timeToCreateResources = getTimeToCollectResourcesForOreRobot(blueprint, resources);
            if (timeToCreateResources !== undefined) {
                optionResources.ore += timeToCreateResources * resources.oreRobots;
                optionResources.clay += timeToCreateResources * resources.clayRobots;
                optionResources.obsidian += timeToCreateResources * resources.obsidianRobots;
                optionResources.geode += timeToCreateResources * resources.geodeRobots;
                optionResources.timeLeft -= timeToCreateResources;
            }

            if (optionResources.timeLeft > 0) {
                optionResources.ore -= blueprint.oreRobotRecipe.ore;
                optionResources.oreRobots++;

                const candidateResources = iterate(optionResources);
                if (isBetter(candidateResources, bestOption)) {
                    bestOption = candidateResources;
                }
            }
        }

        bestSoFar = Math.max(bestSoFar, bestOption.geode);

        memo[memoKey] = bestOption;
        return bestOption;
    };

    return iterate(resources);
};

const blueprints = loadBlueprints();

const initialResource: Resources = {
    timeLeft: 32,
    ore: 0,
    clay: 0,
    obsidian: 0,
    geode: 0,
    oreRobots: 1,
    clayRobots: 0,
    obsidianRobots: 0,
    geodeRobots: 0,
};

startTimer();

let qualitySum: number = 0;

for (const blueprint of blueprints) {
    const resources: Resources = getBestGeodeCount(blueprint, initialResource);
    const quality = blueprint.id * resources.geode;
    qualitySum += quality;

    log(`Blueprint #${blueprint.id} Max Geode: ${resources.geode}`);
}

log(qualitySum);

stopTimer();
