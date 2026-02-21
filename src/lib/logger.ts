import chalk from "chalk";

export const logger = {
  info(msg: string): void {
    process.stdout.write(`${chalk.blue("info")} ${msg}\n`);
  },
  success(msg: string): void {
    process.stdout.write(`${chalk.green("success")} ${msg}\n`);
  },
  warn(msg: string): void {
    process.stderr.write(`${chalk.yellow("warn")} ${msg}\n`);
  },
  error(msg: string): void {
    process.stderr.write(`${chalk.red("error")} ${msg}\n`);
  },
};
