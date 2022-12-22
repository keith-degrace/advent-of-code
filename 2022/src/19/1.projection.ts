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

const getBestGeodeCount = (blueprint: Blueprint, resources: Resources, minute: number): Resources => {
    const newResources: Resources = {
        ...resources,
        ore: resources.ore + resources.oreRobots,
        clay: resources.clay + resources.clayRobots,
        obsidian: resources.obsidian + resources.obsidianRobots,
        geode: resources.geode + resources.geodeRobots,
    };

    if (minute == 24) {
        return newResources;
    }

    // If we can create a geode robot, it's a no-brainer.
    if (resources.ore >= blueprint.geodeRobotRecipe.ore && resources.obsidian >= blueprint.geodeRobotRecipe.obsidian) {
        return getBestGeodeCount(
            blueprint,
            {
                ...newResources,
                ore: newResources.ore - blueprint.geodeRobotRecipe.ore,
                obsidian: newResources.obsidian - blueprint.geodeRobotRecipe.obsidian,
                geodeRobots: newResources.geodeRobots + 1,
            },
            minute + 1
        );
    }

    // If we can't make a geode robot, we project how much ore/obsidian we will create and determine the best robot
    // to produce based on what is lacking.

    const projectedOre = resources.oreRobots * (24 - minute);
    const projectedClay = resources.clayRobots * (24 - minute);
    const projectedObsidian = resources.obsidianRobots * (24 - minute);

    const projectedGeodeRobotsFromOre = blueprint.geodeRobotRecipe.ore / projectedOre;
    const projectedGeodeRobotsFromObsidian = blueprint.geodeRobotRecipe.obsidian / projectedObsidian;

    if (projectedGeodeRobotsFromObsidian < projectedGeodeRobotsFromOre) {
        // Create an obsidian robot.
        if (resources.ore >= blueprint.obsidianRobotRecipe.ore && resources.clay >= blueprint.obsidianRobotRecipe.clay) {
            return getBestGeodeCount(
                blueprint,
                {
                    ...newResources,
                    ore: newResources.ore - blueprint.obsidianRobotRecipe.ore,
                    clay: newResources.clay - blueprint.obsidianRobotRecipe.clay,
                    obsidianRobots: newResources.obsidianRobots + 1,
                },
                minute + 1
            );
        }
    }

    // Create a clay robot.
    if (resources.ore >= blueprint.clayRobotRecipe.ore) {
        return getBestGeodeCount(
            blueprint,
            {
                ...newResources,
                ore: newResources.ore - blueprint.clayRobotRecipe.ore,
                clayRobots: newResources.clayRobots + 1,
            },
            minute + 1
        );
    }

    // Create an ore robot.
    if (resources.ore >= blueprint.oreRobotRecipe.ore) {
        return getBestGeodeCount(
            blueprint,
            {
                ...newResources,
                ore: newResources.ore - blueprint.oreRobotRecipe.ore,
                oreRobots: newResources.oreRobots + 1,
            },
            minute + 1
        );
    }

    // Don't create a robot
    return getBestGeodeCount(blueprint, newResources, minute + 1);
};

const blueprints = loadBlueprints();

const initialResource: Resources = {
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
    const resources: Resources = getBestGeodeCount(blueprint, initialResource, 1);
    log(JSON.stringify(resources, null, 2));

    const quality = blueprint.id * resources.geode;
    qualitySum += quality;
}

log(qualitySum);
