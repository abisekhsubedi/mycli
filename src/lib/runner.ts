import { createProgram } from "../commands/index.js";
import { handleExit } from "./exit.js";

export async function run(): Promise<void> {
  handleExit();

  const program = createProgram();

  try {
    await program.parseAsync(process.argv);
  } catch (err) {
    if (err instanceof Error) {
      process.stderr.write(`Error: ${err.message}\n`);
    }
    process.exit(1);
  }
}
