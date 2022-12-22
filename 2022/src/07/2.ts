import * as fs from "fs";
import * as path from "path";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\n");

interface File {
    name: string;
    size: number;
}

interface Folder {
    parent: Folder;
    size: number;
    files: File[];
    folders: Record<string, Folder>;
}

const addFolder = (parent: Folder, name: string): Folder => {
    if (!parent) {
        return { parent: undefined, size: 0, files: [], folders: {} };
    }

    if (!parent.folders[name]) {
        parent.folders[name] = { parent, size: 0, files: [], folders: {} };
    }

    return parent.folders[name];
};

const root: Folder = {
    parent: undefined,
    size: 0,
    files: [],
    folders: {},
};

let current: Folder = root;

for (const line of input) {
    if (line.startsWith("$ cd")) {
        const folderName: string = line.split(" ")[2];
        if (folderName === "/") {
            current = root;
        } else if (folderName === "..") {
            current = current.parent;
        } else {
            current = addFolder(current, folderName);
        }
    } else if (line.startsWith("$ ls")) {
    } else if (line.startsWith("dir")) {
        const folderName: string = line.split(" ")[1];
        addFolder(current, folderName);
    } else {
        const size: number = Number.parseInt(line.split(" ")[0]);
        const name: string = line.split(" ")[1];
        current.files.push({ name, size });
    }
}

const calculateSizes = (folder: Folder): void => {
    for (const subFodler of Object.entries(folder.folders).values()) {
        calculateSizes(subFodler[1]);
    }

    for (const file of folder.files) {
        folder.size += file.size;
    }

    for (const subFodler of Object.entries(folder.folders).values()) {
        folder.size += subFodler[1].size;
    }
};

calculateSizes(root);

const TotalSpace = 70000000;
const TargetUnusedSpace = 30000000;
const CurrentUnusedSpace = TotalSpace - root.size;
const SpaceToDelete = TargetUnusedSpace - CurrentUnusedSpace;

let bestSize: number = TotalSpace;

const findFolder = (folder: Folder): void => {
    if (folder.size > SpaceToDelete && folder.size < bestSize) {
        bestSize = folder.size;
    }

    for (const subFodler of Object.entries(folder.folders).values()) {
        findFolder(subFodler[1]);
    }
};

findFolder(root);

console.log(bestSize);
