import { ipcMain } from 'electron';
import { AppApiHandler, AppApiResponse, ProgramInfoDto, UIConfigDto } from 'interop/infra';

import { LogPort } from '../domain/ports/log.port.js';
import { AppIoc } from './ioc.js';
import { InfraContextPort } from './ports/infra-context.port.js';

export interface ApiControllerParams {
  logPort: LogPort;
  contextPort: InfraContextPort;
  ioc: AppIoc;
}

export function initializeApiController(app: Electron.App, params: ApiControllerParams): void {
  handleApi(AppApiHandler.GetProgramInfo, params.logPort, async (): Promise<ProgramInfoDto> => {
    return Promise.resolve({
      appVersion: `${params.contextPort.appVersion}b`,
      nodeVersion: process.versions.node,
      chromeVersion: process.versions.chrome,
      electronVersion: process.versions.electron,
    });
  });

  handleApi(AppApiHandler.GetUIConfig, params.logPort, async (): Promise<UIConfigDto> => {
    return {
      locale: params.contextPort.forceLocale ?? app.getLocaleCountryCode().toLowerCase(),
      logLevel: params.contextPort.logLevel,
    };
  });

  handleApi(AppApiHandler.GetThrowing, params.logPort, async (): Promise<void> => {
    throw new Error('Throwing error from API');
  });
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
function handleApi<T = void>(channel: AppApiHandler, logPort: LogPort, handler: (...args: any[]) => Promise<T>): void {
  ipcMain.handle(channel, async (_event, ...args): Promise<AppApiResponse<T>> => {
    try {
      const result = await handler(...args);
      return { data: result, error: undefined };
    } catch (error) {
      logPort.exception(error, { module: initializeApiController.name, msg: `Handler error for ${channel}` });
      return { data: undefined, error };
    }
  });
}
