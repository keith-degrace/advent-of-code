import * as fs from "fs";
import * as path from "path";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\n");

interface Position {
    x: number;
    y: number;
}

interface Sensor {
    position: Position;
    closestBeacon: Position;
}

const sensors: Sensor[] = [];

for (const line of input) {
    const match = line.match(/.*x=(.+).*y=(.+).*x=(.+).*y=(.+)/);

    sensors.push({
        position: {
            x: Number.parseInt(match[1]),
            y: Number.parseInt(match[2]),
        },
        closestBeacon: {
            x: Number.parseInt(match[3]),
            y: Number.parseInt(match[4]),
        },
    });
}

const distance = (a: Position, b: Position): number => {
    return Math.abs(a.x - b.x) + Math.abs(a.y - b.y);
};

const min: Position = { x: Infinity, y: Infinity };
const max: Position = { x: -Infinity, y: -Infinity };
for (const sensor of sensors) {
    const distanceToBeacon: number = distance(sensor.position, sensor.closestBeacon);

    min.x = Math.min(min.x, sensor.position.x - distanceToBeacon);
    min.y = Math.min(min.y, sensor.position.y - distanceToBeacon);
    max.x = Math.max(max.x, sensor.position.x + distanceToBeacon);
    max.y = Math.max(max.y, sensor.position.y + distanceToBeacon);
}

const canContainBeacon = (position: Position): boolean => {
    for (const sensor of sensors) {
        if (sensor.closestBeacon.x === position.x && sensor.closestBeacon.y === position.y) {
            return false;
        }

        const distanceToBeacon: number = distance(sensor.position, sensor.closestBeacon);
        const distanceToLocation: number = distance(sensor.position, position);

        if (distanceToLocation <= distanceToBeacon) {
            return false;
        }
    }

    return true;
};

let count = 0;

for (let x = min.x; x <= max.x; x++) {
    if (!canContainBeacon({ x, y: 2000000 })) {
        count++;
    }
}

console.log(count);
