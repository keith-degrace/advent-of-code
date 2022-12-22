import * as fs from "fs";
import * as path from "path";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\n");

interface File {
    name: string;
    size: number;
}

interface Folder {
    parent: Folder;
    files: File[];
    folders: Record<string, Folder>;
}

const addFolder = (parent: Folder, name: string): Folder => {
    if (!parent) {
        return { parent: undefined, files: [], folders: {} };
    }

    if (!parent.folders[name]) {
        parent.folders[name] = { parent, files: [], folders: {} };
    }

    return parent.folders[name];
};

const root: Folder = {
    parent: undefined,
    files: [],
    folders: {},
};

let current: Folder = root;

for (const line of input) {
    if (line.startsWith("$ cd")) {
        const folderName: string = line.split(" ")[2];
        console.log(`[cd] ${folderName}`);
        if (folderName === "/") {
            current = root;
        } else if (folderName === "..") {
            current = current.parent;
        } else {
            current = addFolder(current, folderName);
        }
    } else if (line.startsWith("$ ls")) {
        console.log(`[ls] ${line}`);
    } else if (line.startsWith("dir")) {
        const folderName: string = line.split(" ")[1];

        console.log(`[dir] ${folderName}`);

        addFolder(current, folderName);
    } else {
        const size: number = Number.parseInt(line.split(" ")[0]);
        const name: string = line.split(" ")[1];

        console.log(`[file] ${line}`);

        current.files.push({ name, size });
    }
}

let sum: number = 0;

const getSize = (folder: Folder): number => {
    let size: number = 0;

    for (const file of folder.files) {
        size += file.size;
    }

    for (const subFodler of Object.entries(folder.folders).values()) {
        size += getSize(subFodler[1]);
    }

    if (size <= 100000) {
        sum += size;
    }

    return size;
};

getSize(root);

console.log(sum);
