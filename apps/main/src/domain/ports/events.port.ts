import { SomeNotificationArgs } from 'interop/domain';

export const EVENTS_PORT_TOKEN = Symbol.for('');

export interface EventsPort {
  someNotification(args: SomeNotificationArgs): Promise<void>;
}
