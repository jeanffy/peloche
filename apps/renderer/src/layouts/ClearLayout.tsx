import { PropsWithChildren, ReactElement } from 'react';

import css from './ClearLayout.module.scss';

export default function ClearLayout({ children }: PropsWithChildren): ReactElement {
  return <div className={css.wrapper}>{children}</div>;
}
