export type LogLevel = "info" | "warn" | "error";

function timestamp(): string {
  return new Date().toISOString();
}

export function log(level: LogLevel, message: string): void {
  const prefix = `[${timestamp()}] [${level.toUpperCase()}]`;

  console.log(`${prefix} ${message}`);
}

export function info(message: string): void {
  log("info", message);
}

export function warn(message: string): void {
  log("warn", message);
}

export function error(message: string): void {
  log("error", message);
}