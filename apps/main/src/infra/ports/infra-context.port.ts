import { ConsoleAndFileLoggerLevel } from 'interop/infra';

import { DomainContextPort } from '../../domain/ports/domain-context.port.js';

export const INFRA_CONTEXT_PORT_TOKEN = Symbol.for('InfraContextPort');

export enum AppPlatform {
  Windows,
  MacOS,
  Linux,
}

export enum AppProtocol {
  File,
  Http,
}

export interface InfraContextPort extends DomainContextPort {
  readonly platform: AppPlatform;
  readonly protocol: AppProtocol;
  readonly appSettingsPath: string;
  readonly forceLocale?: string;
  readonly logLevel: ConsoleAndFileLoggerLevel;
  readonly preloadJsPath: string;
  readonly rendererIndexHtmlPath: string;
  readonly rendererUrl: string;
}
