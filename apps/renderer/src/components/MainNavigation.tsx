import { ReactElement } from 'react';
import { Link } from 'react-router';

import { ROUTE_ANOTHER_PAGE, ROUTE_HOME } from '../router/router-names.js';
import css from './MainNavigation.module.scss';

export default function MainNavigation(): ReactElement {
  return (
    <ul className={css.nav}>
      <li>
        <Link to={ROUTE_HOME}>Home</Link>
      </li>
      <li>
        <Link to={ROUTE_ANOTHER_PAGE}>Another</Link>
      </li>
      <li>
        <Link to="/foobar">Nonexistent link</Link>
      </li>
    </ul>
  );
}
