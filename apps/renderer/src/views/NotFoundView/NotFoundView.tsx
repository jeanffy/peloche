import { ReactElement } from 'react';
import { Link } from 'react-router';

import ClearLayout from '../../layouts/ClearLayout';
import { ROUTE_HOME } from '../../router/router-names';
import css from './NotFoundView.module.scss';

export default function NotFoundView(): ReactElement {
  return (
    <ClearLayout>
      <h1 className={css.notFound}>Not found</h1>
      <p>{window.location.href}</p>
      <Link to={ROUTE_HOME}>Go to home</Link>
    </ClearLayout>
  );
}
