import { initReactI18next } from 'react-i18next';
import i18n from 'i18next';
import Backend from 'i18next-http-backend';
import LanguageDetector from 'i18next-browser-languagedetector';

const supportedLngs = ['en'];

i18n
  .use(initReactI18next)
  .use(Backend)
  .init({
    //debug: true,
    lng: 'en',
    fallbackLng: 'en',
    interpolation: { escapeValue: false },
    react: { useSuspense: false },
    supportedLngs,
    backend: {
      loadPath: '/i18n-{{lng}}.json',
    },
  });

export default i18n;
