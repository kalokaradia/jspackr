export function clamp(value: number, min: number, max: number): number {
  return Math.min(Math.max(value, min), max);
}

export function average(values: number[]): number {
  if (values.length === 0) {
    return 0;
  }

  return values.reduce((sum, value) => sum + value, 0) / values.length;
}

export function percentage(value: number, total: number): number {
  if (total === 0) {
    return 0;
  }

  return (value / total) * 100;
}

export function randomInt(min: number, max: number): number {
  const range = max - min + 1;

  return Math.floor(Math.random() * range) + min;
}