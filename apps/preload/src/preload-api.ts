import { contextBridge, ipcRenderer } from 'electron';
import { DummyEntity, GetDummiesUseCaseParams } from 'interop/domain';
import { AppApi, AppApiHandler, appApiName, AppApiResponse, ProgramInfoDto, UIConfigDto } from 'interop/infra';

const appApiHandlers: AppApi = {
  getProgramInfo: () => invokeApi<ProgramInfoDto>(AppApiHandler.GetProgramInfo),
  getUIConfig: () => invokeApi<UIConfigDto>(AppApiHandler.GetUIConfig),
  getThrowing: () => invokeApiVoid(AppApiHandler.GetThrowing),
  getDummies: (params: GetDummiesUseCaseParams) => invokeApi<DummyEntity[]>(AppApiHandler.GetDummies, params),
};

async function invokeApiVoid(channel: AppApiHandler, ...args: unknown[]): Promise<void> {
  const response = (await ipcRenderer.invoke(channel, ...args)) as AppApiResponse<void>;
  if (response.error !== undefined) {
    throw response.error;
  }
}

async function invokeApi<T>(channel: AppApiHandler, ...args: unknown[]): Promise<T> {
  const response = (await ipcRenderer.invoke(channel, ...args)) as AppApiResponse<T>;
  if (response.error !== undefined) {
    throw response.error;
  }
  if (response.data === undefined) {
    throw new Error(`No data returned from invoke ${channel}`);
  }
  return response.data;
}

contextBridge.exposeInMainWorld(appApiName, appApiHandlers);
