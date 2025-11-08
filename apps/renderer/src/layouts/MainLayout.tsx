import { PropsWithChildren, ReactElement } from 'react';

import MainNavigation from '../components/MainNavigation';
import css from './MainLayout.module.scss';

export default function MainLayout({ children }: PropsWithChildren): ReactElement {
  return (
    <div className={css.wrapper}>
      <div className={css.navigation}>
        <MainNavigation />
      </div>
      <div className={css.content}>{children}</div>
    </div>
  );
}
