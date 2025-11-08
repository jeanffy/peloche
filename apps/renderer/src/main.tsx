import './styles/index.scss';
import './i18n';

import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';

import App from './App.tsx';
import AppContext from './helpers/app-context.ts';

async function main(): Promise<void> {
  await AppContext.instance().initialize();

  createRoot(document.getElementById('root')!).render(
    <StrictMode>
      <App />
    </StrictMode>,
  );
}

void main();
