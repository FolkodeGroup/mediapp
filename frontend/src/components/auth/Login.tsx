
"use client";

import { useEffect, useState } from "react";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import Input from "../../ui/Input";
import { useAuth } from "../../auth/AuthContext";
import Message from "./Message";
import { useNavigate } from "react-router-dom";

const loginSchema = z.object({
  username: z.string().min(1, "El usuario es obligatorio"),
  password: z.string().min(1, "La contraseña es obligatoria"),
});


const LoginForm = () => {
  const { login } = useAuth();
  const [message, setMessage] = useState<{ type: 'success' | 'error'; text: string } | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const navigate = useNavigate();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<z.infer<typeof loginSchema>>({
    resolver: zodResolver(loginSchema),
    mode: "onChange",
  });

  const onSubmit = async (data: z.infer<typeof loginSchema>) => {


    setIsLoading(true);
    try {
      await Promise.resolve(login(data));
      setMessage({ type: 'success', text: 'Inicio de sesión exitoso. Redirigiendo...' });
      setTimeout(() => {
        navigate('/dashboard');
      }, 1000);
    } catch (error: any) {
      let errorMsg = 'Error durante el login.';
      if (error?.response) {
        const status = error.response.status;
        if (status === 400) {
          errorMsg = 'Solicitud inválida. Verifica los datos ingresados.';
        } else if (status === 401) {
          errorMsg = 'Credenciales incorrectas o usuario bloqueado.';
        } else if (status === 500) {
          errorMsg = 'Error interno del servidor. Intenta más tarde.';
        } else if (status === 403) {
          errorMsg = 'Usuario bloqueado. Contactate con el administrador del software.';
        }
        else {
          // Si el backend envía un mensaje específico
          if (error.response.data?.message) {
            errorMsg = error.response.data.message;
          } else if (error.response.data?.error) {
            errorMsg = error.response.data.error;
          }
        }
      } else if (error?.message === 'Network Error') {
        errorMsg = 'No se pudo conectar con el servidor. Verifica tu conexión.';
      }
      setMessage({ type: 'error', text: errorMsg });
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    if (message) {
      setTimeout(() => {
        setMessage(null);
      }, 3000);
    }
  }, [message])

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="space-y-4 form-login"
      autoComplete="off"
      noValidate
    >
      {/* Usuario */}
      <div>
        <label
          htmlFor="name"
          className="block text-sm font-medium text-gray-700 mb-1 campo-login"
        >
          Usuario:
        </label>
        <Input
          id="name"
          {...register("username")}
          placeholder="Usuario"
          aria-invalid={!!errors.username}
          className="ph-login"
          autoComplete="username"
        />
        {errors.username && (
          <p className="text-sm text-red-600">{errors.username.message}</p>
        )}
      </div>

      <div>
        <label
          htmlFor="contrasena"
          className="block text-sm font-medium text-gray-700 mb-1 campo-login"
        >
          Contraseña:
        </label>
        <Input
          type="password"
          id="contrasena"
          {...register("password")}
          placeholder="Contraseña"
          aria-invalid={!!errors.password}
          className="ph-login"
          autoComplete="current-password"
        />
        {errors.password && (
          <p className="text-sm text-red-600">{errors.password.message}</p>
        )}
      </div>

      {message && (
        <Message type={message.type} text={message.text} />
      )}

      {/* Botones */}
      <div className="flex gap-2 justify-center pt-2">
        <button type="submit" className="btn-login" disabled={isLoading}>
          {isLoading ? "Cargando..." : "Iniciar sesión"}
        </button>
        <button
          type="button"
          className="btn-login"
          onClick={() => navigate("/register")}
        >
          Registrarse
        </button>
      </div>
    </form>
  );
};

export default LoginForm;
