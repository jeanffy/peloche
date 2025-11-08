import { createHashRouter } from 'react-router';

import AnotherPageView from '../views/AnotherPageView/AnotherPageView';
import HomeView from '../views/HomeView/HomeView';
import NotFoundView from '../views/NotFoundView/NotFoundView';
import { ROUTE_ANOTHER_PAGE, ROUTE_HOME } from './router-names';

export default function getRouter(): ReturnType<typeof createHashRouter> {
  return createHashRouter([
    {
      path: ROUTE_HOME,
      element: <HomeView />,
    },
    {
      path: ROUTE_ANOTHER_PAGE,
      element: <AnotherPageView />,
    },
    {
      path: '*',
      element: <NotFoundView />,
    },
  ]);
}
