import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import LanguageDetector from "i18next-browser-languagedetector";
import HttpApi from 'i18next-http-backend';

i18n
  .use(HttpApi) // Carga traducciones desde un servidor (ej. /public/locales)
  .use(LanguageDetector) // Detecta el idioma del usuario
  .use(initReactI18next) // Pasa la instancia de i18n a react-i18next.
  .init({
    supportedLngs: ["en", "es"],
    fallbackLng: "es", // Idioma por defecto si el del navegador no está disponible
    debug: true, // Poner en false en producción
    interpolation: {
      escapeValue: false, // React ya protege contra XSS
    },
    backend: {
      loadPath: '/locales/{{lng}}/{{ns}}.json', // Ruta a los archivos de traducción
    },
  });

export default i18n;
