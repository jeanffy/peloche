import { asValue } from 'awilix';
import { BrowserWindow } from 'electron';

import { LOG_PORT_TOKEN, LogPort } from '../domain/ports/log.port.js';
import { InjectedDependencies } from '../types.js';
import { ELECTRON_APP_MAIN_WINDOW_TOKEN } from './adapters/electron-events.adapter.js';
import { ApiBackgroundTimer } from './api-background-timer.js';
import { initializeApiController } from './api-controller.js';
import { AppIoc } from './ioc.js';
import { AppPlatform, AppProtocol, INFRA_CONTEXT_PORT_TOKEN, InfraContextPort } from './ports/infra-context.port.js';

export class ElectronApp {
  private logPort: LogPort;
  private contextPort: InfraContextPort;

  private ioc?: AppIoc;
  private win?: BrowserWindow;

  private backgroundTimer?: ApiBackgroundTimer;

  public constructor(dependencies: InjectedDependencies) {
    this.logPort = dependencies[LOG_PORT_TOKEN];
    this.contextPort = dependencies[INFRA_CONTEXT_PORT_TOKEN];
  }

  public async init(app: Electron.App, ioc: AppIoc): Promise<void> {
    this.ioc = ioc;

    await this.ioc.initialize();
    await app.whenReady();
    initializeApiController(app, {
      logPort: this.logPort,
      contextPort: this.contextPort,
      ioc,
    });

    this.win = await this.createWindow();
    this.ioc.register<BrowserWindow>(ELECTRON_APP_MAIN_WINDOW_TOKEN, asValue(this.win));

    this.backgroundTimer = this.ioc.getItem(ApiBackgroundTimer);
    await this.backgroundTimer.initialize();

    app.on('activate', (): void => {
      if (BrowserWindow.getAllWindows().length === 0) {
        void this.createWindow();
      }
    });

    app.on('window-all-closed', (): void => {
      if (this.contextPort.platform !== AppPlatform.MacOS) {
        app.quit();
      }
    });

    let quitInProgress = false;
    app.on('before-quit', event => {
      if (!quitInProgress) {
        event.preventDefault();
        quitInProgress = true;
        setTimeout(() => {
          const quitPromises: Promise<unknown>[] = [];
          if (this.ioc !== undefined) {
            quitPromises.push(this.ioc.dispose());
          }
          if (this.backgroundTimer !== undefined) {
            quitPromises.push(this.backgroundTimer.terminate());
          }
          Promise.all(quitPromises)
            .catch(error => {
              console.error('error', error);
            })
            .finally(() => {
              app.quit();
            });
        });
      }
    });
  }

  private async createWindow(): Promise<BrowserWindow> {
    const rendererPath = this.contextPort.protocol === AppProtocol.Http ? this.contextPort.rendererUrl : this.contextPort.rendererIndexHtmlPath;

    this.logPort.info({
      module: ElectronApp.name,
      msg: 'createWindow',
      payload: {
        cwd: process.cwd(),
        preloadPath: this.contextPort.preloadJsPath,
        rendererPath: rendererPath,
      },
    });

    const win = new BrowserWindow({
      show: false,
      webPreferences: {
        preload: this.contextPort.preloadJsPath,
      },
    });

    if (this.contextPort.protocol === AppProtocol.Http) {
      await win.loadURL(this.contextPort.rendererUrl);
    } else {
      await win.loadFile(this.contextPort.rendererIndexHtmlPath);
    }

    win.show();

    return win;
  }
}
