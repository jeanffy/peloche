//
// renderer to main communication
// messages sent from the renderer process to the main process
//

import { DummyEntity } from '../domain/models/dummy.entity.js';
import { GetDummiesUseCaseParams } from '../domain/use-cases-params.js';

export const appApiName = 'pelocheApi';

export enum AppApiHandler {
  GetProgramInfo = 'GetProgramInfo',
  GetUIConfig = 'GetUIConfig',
  GetThrowing = 'GetThrowing',
  GetDummies = 'GetDummies',
}

export interface ProgramInfoDto {
  appVersion: string;
  nodeVersion: string;
  chromeVersion: string;
  electronVersion: string;
}

export interface UIConfigDto {
  locale: string;
  logLevel: number;
}

export interface AppApiResponse<T> {
  data?: T;
  error?: unknown;
}

export interface AppApi {
  getProgramInfo: () => Promise<ProgramInfoDto>;
  getUIConfig: () => Promise<UIConfigDto>;
  getThrowing: () => Promise<void>;
  getDummies: (params: GetDummiesUseCaseParams) => Promise<DummyEntity[]>;
}
