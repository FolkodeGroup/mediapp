import { useTranslation } from "react-i18next";

const LanguageSwitcher = () => {
  const { i18n } = useTranslation();

  const changeLanguage = (lng: string) => {
    i18n.changeLanguage(lng);
  };

  return (
    <div className="p-4">
      <button
        onClick={() => changeLanguage("en")}
        disabled={i18n.language === "en"}
        className="px-3 py-1 mr-2 text-sm font-medium text-white bg-blue-600 rounded-md disabled:bg-gray-400 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
      >
        English
      </button>
      <button
        onClick={() => changeLanguage("es")}
        disabled={i18n.language === "es"}
        className="px-3 py-1 text-sm font-medium text-white bg-blue-600 rounded-md disabled:bg-gray-400 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
      >
        Español
      </button>
    </div>
  );
};

export default LanguageSwitcher;
