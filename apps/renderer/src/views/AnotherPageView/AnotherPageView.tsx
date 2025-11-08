import { ReactElement } from 'react';

import MainLayout from '../../layouts/MainLayout.js';

export default function AnotherPageView(): ReactElement {
  return (
    <MainLayout>
      <h3>Another page</h3>
      <p>Another page</p>
    </MainLayout>
  );
}
