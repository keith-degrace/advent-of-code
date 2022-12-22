import * as fs from "fs";
import * as path from "path";
import { Position } from "../utils/position";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\n");

interface Sensor {
    position: Position;
    range: number;
}

const distance = (a: Position, b: Position): number => {
    return Math.abs(a.x - b.x) + Math.abs(a.y - b.y);
};

const sensors: Sensor[] = input.map((line) => {
    const match = line.match(/.*x=(.+).*y=(.+).*x=(.+).*y=(.+)/);

    const position: Position = {
        x: Number.parseInt(match[1]),
        y: Number.parseInt(match[2]),
    };

    const closestBeacon: Position = {
        x: Number.parseInt(match[3]),
        y: Number.parseInt(match[4]),
    };

    const range: number = distance(position, closestBeacon);

    return { position, range };
});

const locateSignal = (): Position => {
    interface Span {
        min: number;
        max: number;
    }

    for (let y = 0; y <= 4000000; y++) {
        const spans: Span[] = [];

        // Calculate the x-axis span of each sensor on this row.
        for (const sensor of sensors) {
            // Ignore any sensor that is out of range of this row.
            if (Math.abs(sensor.position.y - y) > sensor.range) {
                continue;
            }

            const halfWidth: number = sensor.range - Math.abs(sensor.position.y - y);
            spans.push({
                min: Math.max(sensor.position.x - halfWidth, 0),
                max: Math.min(sensor.position.x + halfWidth, 4000000),
            });
        }

        // Sort the spans by their min bound.
        spans.sort((a, b) => a.min - b.min);

        // Advance through spans looking for a gap.
        let currentSpanMax = spans[0].max;
        for (let i = 1; i < spans.length - 1; i++) {
            // If this span's min is higher than our current max, we found the gap!
            if (spans[i].min > currentSpanMax) {
                return { x: currentSpanMax + 1, y };
            }

            // If this span's max bound is higher than our current one, then this is our new max.
            currentSpanMax = Math.max(currentSpanMax, spans[i].max);
        }
    }
};

const signalPosition: Position = locateSignal();
const tuningFrequency: number = signalPosition.x * 4000000 + signalPosition.y;

console.log(tuningFrequency);
