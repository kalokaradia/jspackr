export interface User {
  id: number;
  name: string;
  age: number;
  active: boolean;
  scores: number[];
}

export class UserModel {
  constructor(private readonly user: User) { }

  get id(): number {
    return this.user.id;
  }

  get name(): string {
    return this.user.name;
  }

  get averageScore(): number {
    if (this.user.scores.length === 0) {
      return 0;
    }

    const total = this.user.scores.reduce((sum, score) => sum + score, 0);

    return total / this.user.scores.length;
  }

  isAdult(): boolean {
    return this.user.age >= 18;
  }

  isActive(): boolean {
    return this.user.active;
  }

  toJSON() {
    return {
      id: this.user.id,
      name: this.user.name,
      age: this.user.age,
      active: this.user.active,
      averageScore: Number(this.averageScore.toFixed(2)),
    };
  }
}