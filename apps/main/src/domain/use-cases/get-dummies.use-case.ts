import { DummyEntity } from 'interop/domain';

import { InjectedDependencies } from '../../types.js';
import { DOMAIN_CONTEXT_PORT_TOKEN, DomainContextPort } from '../ports/domain-context.port.js';
import { FS_PORT_TOKEN, FsPort } from '../ports/fs.port.js';

export class GetDummiesUseCase {
  private contextPort: DomainContextPort;
  private fsPort: FsPort;

  public constructor(dependencies: InjectedDependencies) {
    this.contextPort = dependencies[DOMAIN_CONTEXT_PORT_TOKEN];
    this.fsPort = dependencies[FS_PORT_TOKEN];
  }

  public async run(): Promise<DummyEntity[]> {
    const dummies: DummyEntity[] = [];

    const entries = await this.fsPort.readDirectory(this.contextPort.dummiesDirPath);
    for (const entry of entries) {
      const content = await this.fsPort.readFile(entry.fullPath);
      dummies.push({
        id: entry.fullPath,
        content,
      });
    }

    return dummies;
  }
}
