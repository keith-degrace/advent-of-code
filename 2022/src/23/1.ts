import * as fs from "fs";
import * as path from "path";
import { log } from "../utils";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");
