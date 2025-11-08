import dotenv from 'dotenv';
import { app, dialog } from 'electron';

import { ElectronApp } from './infra/electron-app.js';
import getAppIoc from './infra/ioc.js';

dotenv.config();

async function main(): Promise<void> {
  try {
    const ioc = getAppIoc();
    const electronApp = ioc.getContainer().build(ElectronApp);
    await electronApp.init(app, ioc);
  } catch (error) {
    const message = error instanceof Error ? error.message : JSON.stringify(error);
    dialog.showErrorBox('Startup error', message);
    process.exit();
  }
}

void main();
