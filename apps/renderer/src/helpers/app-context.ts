import log from 'electron-log/renderer';
import { AppApi, AppApiEvents, appApiEventsName, appApiName, ConsoleAndFileLogger } from 'interop/infra';

export default class AppContext {
  private api: AppApi;
  private apiEvents: AppApiEvents;
  private locale: string;
  private logger: ConsoleAndFileLogger;

  public static get api(): AppApi {
    return this.instance().api;
  }

  public static get apiEvents(): AppApiEvents {
    return this.instance().apiEvents;
  }

  public static get locale(): string {
    return this.instance().locale;
  }

  public static get logger(): ConsoleAndFileLogger {
    return this.instance().logger;
  }

  private constructor() {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    this.api = (window as any)[appApiName] as AppApi;

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    this.apiEvents = (window as any)[appApiEventsName] as AppApiEvents;

    this.locale = 'fr';
    this.logger = new ConsoleAndFileLogger({
      logLevel: 0,
      log,
      process: 'renderer',
    });
  }

  public async initialize(): Promise<void> {
    // api is not mounted into window when not running inside Electron
    if (this.api === undefined) {
      this.logger.error({ module: 'AppContext', msg: 'No api in window object', payload: appApiName });
    }

    // apiEvents is not mounted into window when not running inside Electron
    if (this.apiEvents === undefined) {
      this.logger.error({ module: 'AppContext', msg: 'No apiEvents in window object', payload: appApiEventsName });
    }

    const uiConfig = await this.api.getUIConfig();

    this.locale = uiConfig.locale;

    this.logger = new ConsoleAndFileLogger({
      logLevel: uiConfig.logLevel,
      log,
      process: 'renderer',
    });
  }

  public static instance(): AppContext {
    if (appContextInstance === undefined) {
      appContextInstance = new AppContext();
    }
    return appContextInstance;
  }
}

let appContextInstance: AppContext | undefined;
