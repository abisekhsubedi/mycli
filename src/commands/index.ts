import { Command } from "commander";
import { greetCommand } from "./greet.js";

export function createProgram(): Command {
  const program = new Command();

  program
    .name("mycli")
    .description("A CLI application")
    .version("0.1.0", "-v, --version")
    .configureOutput({
      writeOut: (str) => process.stdout.write(str),
      writeErr: (str) => process.stderr.write(str),
      outputError: (str, write) => write(str),
    });

  program.addCommand(greetCommand());

  return program;
}
