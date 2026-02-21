const cleanup: (() => void | Promise<void>)[] = [];

export function onExit(fn: () => void | Promise<void>): void {
  cleanup.push(fn);
}

export function handleExit(): void {
  const signals: NodeJS.Signals[] = ["SIGINT", "SIGTERM", "SIGHUP"];

  for (const signal of signals) {
    process.on(signal, async () => {
      process.stderr.write(`\nReceived ${signal}, shutting down...\n`);
      for (const fn of cleanup) {
        try {
          await fn();
        } catch {
          // best-effort cleanup
        }
      }
      process.exit(128 + (signal === "SIGINT" ? 2 : signal === "SIGTERM" ? 15 : 1));
    });
  }

  process.on("uncaughtException", (err) => {
    process.stderr.write(`Uncaught exception: ${err.message}\n`);
    process.exit(1);
  });

  process.on("unhandledRejection", (reason) => {
    process.stderr.write(`Unhandled rejection: ${reason}\n`);
    process.exit(1);
  });
}
