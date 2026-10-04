import type { User } from "../models/user";
import { UserModel } from "../models/user";
import { clamp } from "../utils/math";
import { info } from "../utils/logger";

export class UserService {
  private readonly users: User[] = [
    {
      id: 1,
      name: "Kaloka",
      age: 15,
      active: true,
      scores: [90, 85, 95, 88],
    },
    {
      id: 2,
      name: "Nanda",
      age: 21,
      active: true,
      scores: [76, 82, 79, 88],
    },
    {
      id: 3,
      name: "Radia",
      age: 17,
      active: false,
      scores: [65, 70, 72, 68],
    },
  ];

  async findById(id: number): Promise<UserModel | null> {
    info(`Searching for user ${id}`);

    await new Promise((resolve) => setTimeout(resolve, 10));

    const user = this.users.find((item) => item.id === id);

    if (!user) {
      return null;
    }

    return new UserModel(user);
  }

  getActiveUsers(): UserModel[] {
    return this.users
      .filter((user) => user.active)
      .map((user) => new UserModel(user));
  }

  calculateScore(user: UserModel): number {
    const score = user.averageScore;

    return clamp(score, 0, 100);
  }
}