import { InjectedDependencies } from '../../types.js';
import { EVENTS_PORT_TOKEN, EventsPort } from '../ports/events.port.js';
import { LOG_PORT_TOKEN, LogPort } from '../ports/log.port.js';

export class SendNotificationIfNeededUseCase {
  private eventsPort: EventsPort;
  private logPort: LogPort;

  public constructor(dependencies: InjectedDependencies) {
    this.eventsPort = dependencies[EVENTS_PORT_TOKEN];
    this.logPort = dependencies[LOG_PORT_TOKEN];
  }

  public async run(): Promise<void> {
    this.logPort.info({ module: SendNotificationIfNeededUseCase.name, msg: 'notification' });
    await this.eventsPort.someNotification({ message: 'hello' });
  }
}
