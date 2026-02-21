import { Command } from "commander";
import { logger } from "../lib/logger.js";

export function greetCommand(): Command {
  return new Command("greet")
    .description("Greet a user")
    .argument("<name>", "Name to greet")
    .option("-u, --uppercase", "Print name in uppercase")
    .action((name: string, opts: { uppercase?: boolean }) => {
      const displayName = opts.uppercase ? name.toUpperCase() : name;
      logger.success(`Hello, ${displayName}!`);
    });
}
