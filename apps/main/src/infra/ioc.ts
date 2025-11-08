import { asClass, asValue, AwilixContainer, ClassOrFunctionReturning, createContainer, Resolver } from 'awilix';

import { DOMAIN_CONTEXT_PORT_TOKEN } from '../domain/ports/domain-context.port.js';
import { EVENTS_PORT_TOKEN } from '../domain/ports/events.port.js';
import { FS_PORT_TOKEN } from '../domain/ports/fs.port.js';
import { LOG_PORT_TOKEN } from '../domain/ports/log.port.js';
import { ConsoleAndFileLogAdapter } from './adapters/console-and-file-log.adapter.js';
import { ELECTRON_APP_MAIN_WINDOW_TOKEN, ElectronEventsAdapter } from './adapters/electron-events.adapter.js';
import { EnvContextAdapter } from './adapters/env-context.adapter.js';
import { RealFsAdapter } from './adapters/real-fs.adapter.js';
import { INFRA_CONTEXT_PORT_TOKEN } from './ports/infra-context.port.js';

export class AppIoc {
  private iocContainer = createContainer();

  public constructor() {
    this.iocContainer.register({
      [FS_PORT_TOKEN]: asClass(RealFsAdapter).singleton(),
      [LOG_PORT_TOKEN]: asClass(ConsoleAndFileLogAdapter).singleton(),
      [DOMAIN_CONTEXT_PORT_TOKEN]: asClass(EnvContextAdapter).singleton(), // TODO: share instance with INFRA_CONTEXT_PORT_TOKEN
      [INFRA_CONTEXT_PORT_TOKEN]: asClass(EnvContextAdapter).singleton(),
      [EVENTS_PORT_TOKEN]: asClass(ElectronEventsAdapter).singleton(),
      [ELECTRON_APP_MAIN_WINDOW_TOKEN]: asValue(undefined),
    });
  }

  public async initialize(): Promise<void> {
    const domainConfigPort = this.getItemFromToken(DOMAIN_CONTEXT_PORT_TOKEN) as EnvContextAdapter;
    await domainConfigPort.initialize();
    const infraConfigPort = this.getItemFromToken(INFRA_CONTEXT_PORT_TOKEN) as EnvContextAdapter;
    await infraConfigPort.initialize();
  }

  public getContainer(): AwilixContainer {
    return this.iocContainer;
  }

  public getItemFromToken<T>(token: symbol): T {
    return this.iocContainer.resolve<T>(token);
  }

  public getItem<T>(targetOrResolver: ClassOrFunctionReturning<T> | Resolver<T>): T {
    return this.iocContainer.build(targetOrResolver);
  }

  public register<T>(token: symbol, registration: Resolver<T>): void {
    this.iocContainer.register(token, registration);
  }

  public async dispose(): Promise<void> {
    await this.iocContainer.dispose();
  }
}

let appIocInstance: AppIoc | undefined;

export default function getAppIoc(): AppIoc {
  if (appIocInstance === undefined) {
    appIocInstance = new AppIoc();
  }
  return appIocInstance;
}
