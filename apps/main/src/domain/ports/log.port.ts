export const LOG_PORT_TOKEN = Symbol.for('LogPort');

export interface LogPortParams {
  msg: string;
  payload?: unknown;
  module: string;
}

export interface LogPort {
  trace(params: LogPortParams): void;
  debug(params: LogPortParams): void;
  info(params: LogPortParams): void;
  warning(params: LogPortParams): void;
  error(params: LogPortParams): void;
  exception(error: unknown, params: LogPortParams): void;
}
