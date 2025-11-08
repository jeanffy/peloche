import { BrowserWindow } from 'electron';
import { SomeNotificationArgs } from 'interop/domain';
import { AppApiEventsHandler } from 'interop/infra';

import { EventsPort } from '../../domain/ports/events.port.js';
import { InjectedDependencies } from '../../types.js';

export const ELECTRON_APP_MAIN_WINDOW_TOKEN = Symbol.for('ElectronAppMainWindowToken');

export class ElectronEventsAdapter implements EventsPort {
  private mainWindow: BrowserWindow;

  public constructor(dependencies: InjectedDependencies) {
    this.mainWindow = dependencies[ELECTRON_APP_MAIN_WINDOW_TOKEN];
  }

  public async someNotification(args: SomeNotificationArgs): Promise<void> {
    this.mainWindow.webContents.send(AppApiEventsHandler.SomeNotification, args);
  }
}
