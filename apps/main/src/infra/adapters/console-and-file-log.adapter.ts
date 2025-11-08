import fs from 'node:fs';
import path from 'node:path';

import log from 'electron-log/main.js';
import { ConsoleAndFileLogger } from 'interop/infra';

import { LogPort } from '../../domain/ports/log.port.js';
import { InjectedDependencies } from '../../types.js';
import { INFRA_CONTEXT_PORT_TOKEN, InfraContextPort } from '../ports/infra-context.port.js';

export class ConsoleAndFileLogAdapter extends ConsoleAndFileLogger implements LogPort {
  public constructor(dependencies: InjectedDependencies) {
    const contextPort: InfraContextPort = dependencies[INFRA_CONTEXT_PORT_TOKEN];

    super({
      logLevel: contextPort.logLevel,
      log,
      process: 'main',
    });

    log.initialize();

    log.transports.file.maxSize = 1024 * 1024 * 1024;
    log.transports.file.archiveLogFn = (oldLogFile): void => {
      const parsed = path.parse(oldLogFile.path);
      try {
        const now = new Date();
        const isoNow = now.toISOString().replaceAll(':', '-').replaceAll('.', '-');
        fs.renameSync(oldLogFile.path, path.join(parsed.dir, `${parsed.name}.${isoNow}${parsed.ext}`));
      } catch (error) {
        console.warn(`Could not rotate log file '${oldLogFile}'`, error);
      }
    };
  }
}
