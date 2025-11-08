import { SendNotificationIfNeededUseCase } from '../domain/use-cases/send-notification-if-needed.use-case.js';
import { InjectedDependencies } from '../types.js';
import getAppIoc from './ioc.js';

const BACKGROUND_INTERVAL_MS = 15_000; // every 15 seconds

export class ApiBackgroundTimer {
  private interval?: NodeJS.Timeout;
  private sendNotificationIfNeededUseCase: SendNotificationIfNeededUseCase;

  public constructor(_dependencies: InjectedDependencies) {
    this.sendNotificationIfNeededUseCase = getAppIoc().getItem(SendNotificationIfNeededUseCase);

    if (this.interval === undefined) {
      this.interval = setInterval(this.onBackgroundTimer.bind(this), BACKGROUND_INTERVAL_MS);
    }
  }

  public async initialize(): Promise<void> {}

  public async terminate(): Promise<void> {
    if (this.interval === undefined) {
      return;
    }
    clearInterval(this.interval);
    this.interval = undefined;
  }

  private async onBackgroundTimer(): Promise<void> {
    await this.sendNotificationIfNeededUseCase.run();
  }
}
