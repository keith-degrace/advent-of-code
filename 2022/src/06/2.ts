import * as fs from "fs";
import * as path from "path";
import { getMarker } from "./1";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim();

console.log(getMarker(input, 14));
