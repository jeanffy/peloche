//
// main to renderer communication
// messages sent from the main process to the renderer process
//

import { SomeNotificationArgs } from '../domain/events-params.js';

export const appApiEventsName = 'pelocheApiEvents';

export enum AppApiEventsHandler {
  SomeNotification = 'SomeNotification',
}

export interface AppApiEventListener0 {
  addListener: (listener: () => void) => void;
  removeListener: (listener: () => void) => void;
}

export interface AppApiEventListener1<T1> {
  addListener: (listener: (args1: T1) => void) => void;
  removeListener: (listener: (arg1: T1) => void) => void;
}

export interface AppApiEvents {
  someNotification: AppApiEventListener1<SomeNotificationArgs>;
}
