import { useQuery } from '@tanstack/react-query';
import { ReactElement } from 'react';

import { apiQuery } from '../../helpers/api-query';
import AppContext from '../../helpers/app-context';
import MainLayout from '../../layouts/MainLayout';
import css from './HomeView.module.scss';

export default function HomeView(): ReactElement {
  const { data, isPending, error } = useQuery({ queryKey: ['getProgramInfo'], queryFn: apiQuery(() => AppContext.api.getProgramInfo()) });
  if (isPending || error !== null) {
    return <></>;
  }

  return (
    <MainLayout>
      <h3>Home</h3>
      <p className={css.helloWorld}>Hello World!</p>
      <p>Version {data.appVersion}</p>
    </MainLayout>
  );
}
