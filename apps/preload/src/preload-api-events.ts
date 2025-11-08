import { contextBridge, ipcRenderer } from 'electron';
import { AppApiEventListener1, AppApiEvents, AppApiEventsHandler, appApiEventsName } from 'interop/infra';

const appApiEventsHandlers: AppApiEvents = {
  someNotification: eventHandler1(AppApiEventsHandler.SomeNotification),
};

// function eventHandler0(channel: string): AppApiEventListener0 {
//   return {
//     addListener: (listener: () => void): void => {
//       ipcRenderer.on(channel, _event => listener());
//     },
//     removeListener: (listener: () => void): void => {
//       ipcRenderer.off(channel, _event => listener());
//     },
//   };
// }

function eventHandler1<T>(channel: string): AppApiEventListener1<T> {
  return {
    addListener: (listener: (param: T) => void): void => {
      ipcRenderer.on(channel, (_event, param) => listener(param));
    },
    removeListener: (listener: (param: T) => void): void => {
      ipcRenderer.off(channel, (_event, param) => listener(param));
    },
  };
}

contextBridge.exposeInMainWorld(appApiEventsName, appApiEventsHandlers);
