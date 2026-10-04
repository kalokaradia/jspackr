import type { UserModel } from "../models/user";
import { average, percentage } from "../utils/math";

export interface UserStats {
  total: number;
  active: number;
  adults: number;
  averageScore: number;
  activePercentage: number;
}

export function calculateStats(users: UserModel[]): UserStats {
  const total = users.length;
  const active = users.filter((user) => user.isActive()).length;
  const adults = users.filter((user) => user.isAdult()).length;

  const scores = users.map((user) => user.averageScore);

  return {
    total,
    active,
    adults,
    averageScore: Number(average(scores).toFixed(2)),
    activePercentage: Number(percentage(active, total).toFixed(2)),
  };
}