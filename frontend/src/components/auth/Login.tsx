"use client";

import { useState } from "react";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import Input from "../../ui/Input";

// Esquema de validación para login
const loginSchema = z.object({
  name: z.string().min(1, "El usuario es obligatorio"),
  contrasena: z.string().min(1, "La contraseña es obligatoria"),
});

const LoginForm = () => {
  const [isError, setIsError] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const [isSuccess, setIsSuccess] = useState(false);
  const [successMessage, setSuccessMessage] = useState("");

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<z.infer<typeof loginSchema>>({
    resolver: zodResolver(loginSchema),
    mode: "onChange",
  });

  const onSubmit = async (data: z.infer<typeof loginSchema>) => {
    setIsError(false);
    setErrorMessage("");
    setIsSuccess(false);
    setSuccessMessage("");

    try {
      // Simulación de login exitoso
      const loginExitoso = true;

      if (!loginExitoso) {
        throw new Error("Credenciales inválidas");
      }

      setIsSuccess(true);
      setSuccessMessage("Inicio de sesión exitoso");
      console.log("Datos validados:", data);
    } catch (error: any) {
      setIsError(true);
      setErrorMessage(error.message || "Error inesperado");
    }
  };

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="form-login"
      autoComplete="off"
      noValidate
    >
      <div className="">
        {/* Usuario */}
        <div className="space-y-1 div-login">
          <label htmlFor="name" className="campo-login">
            Usuario:
          </label>
          <Input
            id="name"
            {...register("name")}
            placeholder="Usuario"
            aria-invalid={!!errors.name}
            className="w-full bg-white border border-gray-300 focus:border-gray-900 text-gray-800 ph-login"
            autoComplete="username"
          />
          <div className="min-h-[20px]">
            {errors.name && (
              <p className="text-sm text-red-600">{errors.name.message}</p>
            )}
          </div>
        </div>

        {/* Contraseña */}
        <div className="space-y-1 div-login">
          <label htmlFor="contrasena" className="campo-login">
            Contraseña:
          </label>
          <Input
            type="password"
            id="contrasena"
            {...register("contrasena")}
            placeholder="Contraseña"
            aria-invalid={!!errors.contrasena}
            className="w-full bg-white border border-gray-300 focus:border-gray-900 text-gray-800 ph-login"
            autoComplete="current-password"
          />
          <div className="min-h-[20px]">
            {errors.contrasena && (
              <p className="text-sm text-red-600">{errors.contrasena.message}</p>
            )}
          </div>
        </div>

        {/* Botones */}
        <div className="space-y-2 mt-4">
          <button type="submit" className="btn-login">
            Iniciar sesión
          </button>
          <button type="submit" className="btn-login">
            Registrarse
          </button>

          {/* Mensaje de error */}
          {isError && (
            <p className="text-sm text-red-600 mt-2">{errorMessage}</p>
          )}

          {/* Mensaje de éxito */}
          {isSuccess && (
            <p className="text-sm text-green-600 mt-2">{successMessage}</p>
          )}
        </div>
      </div>
    </form>
  );
};

export default LoginForm;
