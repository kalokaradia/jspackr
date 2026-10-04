export interface AppConfig {
  name: string;
  version: string;
  environment: "development" | "production";
}

export const config: AppConfig = {
  name: "jspackr-ts-test",
  version: "1.0.0",
  environment: "development",
};