export const numbers = [1, 2, 3, 4, 5];

export function sum(values) {
  return values.reduce((total, value) => total + value, 0);
}
