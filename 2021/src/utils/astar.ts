import { getManhattanDistance, isPositionsEqual, Position } from "./position";

export interface Node {
    parent?: Node;
    position: Position;
    g: number;
    h: number;
    f: number;
}

export const getShortestPath = (start: Position, end: Position, getNeighbors: (position: Position) => Position[]): Position[] => {
    const open: Node[] = [{ position: start, f: 0, h: getManhattanDistance(start, end), g: 0 }];
    const closed: Node[] = [];

    while (open.length > 0) {
        open.sort((a, b) => b.f - a.f);
        const current: Node = open.pop();
        closed.push(current);

        if (isPositionsEqual(current.position, end)) {
            return buildPath(current);
        }

        for (const neighbor of getNeighbors(current.position)) {
            if (closed.find((node) => isPositionsEqual(node.position, neighbor))) {
                continue;
            }

            const g: number = current.g + 1;
            const h: number = getManhattanDistance(neighbor, end);
            const f: number = g + h;

            const neighborInOpen = open.find((node) => isPositionsEqual(node.position, neighbor));
            if (neighborInOpen && neighborInOpen.f <= f) {
                continue;
            }

            open.push({ parent: current, position: neighbor, g, h, f });
        }
    }
};

const buildPath = (end: Node): Position[] => {
    const path: Position[] = [];

    for (let current = end; current !== undefined; current = current.parent) {
        path.push(current.position);
    }

    return path.reverse();
};
