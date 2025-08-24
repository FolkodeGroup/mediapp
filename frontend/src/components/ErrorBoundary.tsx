import { Component, ErrorInfo, ReactNode } from "react";

interface Props {
  children: ReactNode;
}

interface State {
  hasError: boolean;
}

class ErrorBoundary extends Component<Props, State> {
  public state: State = {
    hasError: false,
  };

  public static getDerivedStateFromError(_: Error): State {
    // Actualiza el estado para que el siguiente renderizado muestre la UI de fallback.
    return { hasError: true };
  }

  public componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    // También puedes registrar el error en un servicio de reporte de errores
    console.error("Uncaught error:", error, errorInfo);
  }

  public render() {
    if (this.state.hasError) {
      // Puedes renderizar cualquier UI de fallback personalizada
      return (
        <div className="flex flex-col items-center justify-center h-screen bg-gray-100 text-center p-4">
          <h1 className="text-4xl font-bold text-red-600 mb-4">Algo salió mal.</h1>
          <p className="text-lg text-gray-700 mb-8">
            Lo sentimos, la aplicación ha encontrado un error inesperado.
          </p>
          <button
            onClick={() => window.location.reload()}
            className="px-6 py-3 bg-blue-600 text-white font-semibold rounded-lg shadow-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-75"
          >
            Refrescar la página
          </button>
        </div>
      );
    }

    return this.props.children;
  }
}

export default ErrorBoundary;
