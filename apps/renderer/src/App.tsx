import './App.scss';

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ReactElement, useEffect } from 'react';
import { Toaster } from 'react-hot-toast';
import { useTranslation } from 'react-i18next';
import { RouterProvider } from 'react-router';

import AppContext from './helpers/app-context';
import getRouter from './router/router';

function App(): ReactElement {
  const { i18n } = useTranslation();

  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  });

  useEffect(() => {
    AppContext.logger.debug({ module: 'renderer App', msg: 'changeLanguage', payload: AppContext.locale });
    void i18n.changeLanguage(AppContext.locale);
  }, [i18n]);

  return (
    <>
      <Toaster />
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={getRouter()} />
      </QueryClientProvider>
    </>
  );
}

export default App;
