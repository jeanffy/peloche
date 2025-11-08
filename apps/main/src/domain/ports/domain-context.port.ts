export const DOMAIN_CONTEXT_PORT_TOKEN = Symbol.for('DomainContextPort');

export interface DomainContextPort {
  readonly appName: string;
  readonly appVersion: string;
  readonly dummiesDirPath: string;
}
