const Login = () => {
  return (
    <div className="max-w-sm mx-auto mt-10 p-6 bg-white dark:bg-gray-800 rounded shadow-md">
      <h2 className="text-lg font-semibold mb-4">Iniciar sesión</h2>
      <form className="space-y-4">
        <input
          type="email"
          placeholder="Correo"
          className="w-full px-3 py-2 border rounded bg-gray-50 dark:bg-gray-700"
        />
        <input
          type="password"
          placeholder="Contraseña"
          className="w-full px-3 py-2 border rounded bg-gray-50 dark:bg-gray-700"
        />
        <button
          type="submit"
          className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700"
        >
          Entrar
        </button>
      </form>
    </div>
  )
}

export default Login
