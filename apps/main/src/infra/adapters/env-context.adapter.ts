import fs from 'node:fs';
import path from 'node:path';

import electron from 'electron';
import { ConsoleAndFileLoggerLevel } from 'interop/infra';

import { DomainContextPort } from '../../domain/ports/domain-context.port.js';
import { AppPlatform, AppProtocol, InfraContextPort } from '../ports/infra-context.port.js';

export class EnvContextAdapter implements InfraContextPort, DomainContextPort {
  // -- infra --
  //
  public readonly platform: AppPlatform;
  public readonly protocol: AppProtocol;
  public readonly appSettingsPath: string;
  public readonly forceLocale?: string;
  public readonly logLevel: ConsoleAndFileLoggerLevel;
  public readonly preloadJsPath: string;
  public readonly rendererIndexHtmlPath: string;
  public readonly rendererUrl: string;

  // -- domain --
  //
  public readonly appName: string;
  public readonly appVersion: string;
  public readonly dummiesDirPath: string;

  public constructor() {
    this.appName = electron.app.getName();
    this.appVersion = '1';

    if (process.platform === 'win32') {
      this.platform = AppPlatform.Windows;
    } else if (process.platform === 'darwin') {
      this.platform = AppPlatform.MacOS;
    } else {
      this.platform = AppPlatform.Linux;
    }

    // ELECTRON_RENDERER_URL is active when debugging with VSCode because in that case
    // renderer is loaded with http protocol rather than with file protocol
    // should be false in production mode (no debugging)
    this.protocol = process.env.ELECTRON_RENDERER_URL !== undefined ? AppProtocol.Http : AppProtocol.File;

    const userDataPath = electron.app.getPath('userData');

    switch (this.platform) {
      case AppPlatform.Windows:
        this.appSettingsPath = `C:\\ProgramData\\${this.appName}\\settings.json`;
        break;
      case AppPlatform.MacOS:
        this.appSettingsPath = path.join(userDataPath, 'App', 'settings.json');
        break;
      case AppPlatform.Linux:
        this.appSettingsPath = '/tmp/settings.json';
        break;
    }

    this.forceLocale = process.env.FORCE_LOCALE;

    this.logLevel = ConsoleAndFileLoggerLevel.Debug;
    switch (process.env.LOG_LEVEL) {
      case `${ConsoleAndFileLoggerLevel.Trace}`:
        this.logLevel = ConsoleAndFileLoggerLevel.Trace;
        break;
      case `${ConsoleAndFileLoggerLevel.Info}`:
        this.logLevel = ConsoleAndFileLoggerLevel.Info;
        break;
      case `${ConsoleAndFileLoggerLevel.Warning}`:
        this.logLevel = ConsoleAndFileLoggerLevel.Warning;
        break;
      case `${ConsoleAndFileLoggerLevel.Error}`:
        this.logLevel = ConsoleAndFileLoggerLevel.Error;
        break;
      case `${ConsoleAndFileLoggerLevel.Exception}`:
        this.logLevel = ConsoleAndFileLoggerLevel.Exception;
        break;
    }

    // some dummy default values
    this.rendererUrl = 'http://localhost';
    this.rendererIndexHtmlPath = 'index.html';

    switch (this.protocol) {
      case AppProtocol.Http: // when debugging in VSCode or running through electron-vite in dev mode
        this.preloadJsPath = path.resolve('output/preload/preload.cjs');
        this.rendererUrl = process.env.ELECTRON_RENDERER_URL ?? 'http://localhost';
        break;
      case AppProtocol.File: // when running through electron-vite in preview mode or directly through electron in production
        // when running the packaged app, the '__dirname' variable points to the parent directory
        // of the 'main.cjs' file, that is 'resources/app/output/main'
        // ok to use sync functions here because we are at application boot
        this.preloadJsPath = path.resolve('output/preload/preload.cjs');
        if (!fs.existsSync(this.preloadJsPath)) {
          this.preloadJsPath = path.join(__dirname, '../preload/preload.cjs');
        }
        this.rendererIndexHtmlPath = path.resolve('output/renderer/index.html');
        if (!fs.existsSync(this.rendererIndexHtmlPath)) {
          this.rendererIndexHtmlPath = path.join(__dirname, '../renderer/index.html');
        }
        break;
    }

    this.dummiesDirPath = process.env.DUMMIES_DIR_PATH;
  }

  public async initialize(): Promise<void> {}
}
