import i18n from 'i18next';
import Backend, { HttpBackendOptions } from 'i18next-http-backend';
import { initReactI18next } from 'react-i18next';
import yaml from 'yaml';

void i18n
  .use(Backend)
  .use(initReactI18next)
  .init<HttpBackendOptions>({
    backend: {
      loadPath: (languages, _namespaces) => {
        if (languages.length === 1) {
          const language = languages[0];
          if (language === 'en' || language === 'fr') {
            return `./locales/${language}.yaml`;
          }
        }
        return '';
      },
      parse: data => {
        return yaml.parse(data);
      },
    },
    fallbackLng: 'en',
    debug: false,
    interpolation: {
      escapeValue: false,
    },
  });

export default i18n;
