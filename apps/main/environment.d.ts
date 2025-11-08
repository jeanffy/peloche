declare global {
  namespace NodeJS {
    interface ProcessEnv {
      ELECTRON_RENDERER_URL?: string; // automatically set by electron-vite
      LOG_LEVEL?: string;
      FORCE_LOCALE?: string;
      DUMMIES_DIR_PATH: string;
    }
  }
}

export {};
