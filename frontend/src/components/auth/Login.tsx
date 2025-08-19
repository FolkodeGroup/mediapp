
"use client";

import { useState } from "react";
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
    try {
      await Promise.resolve(login(data));
      setMessage({ type: 'success', text: 'Inicio de sesión exitoso. Redirigiendo...' });
      setTimeout(() => {
        navigate('/dashboard');
      }, 1000);
    } catch (error: any) {
      let errorMsg = 'Error durante el login.';
      if (error?.response?.data?.message) {
        errorMsg = error.response.data.message;
      }
      setMessage({ type: 'error', text: errorMsg });
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
            {...register("username")}
            placeholder="Usuario"
            aria-invalid={!!errors.username}
            className="w-full bg-white border border-gray-300 focus:border-gray-900 text-gray-800 ph-login"
            autoComplete="username"
          />
          <div className="min-h-[20px]">
            {errors.username && (
              <p className="text-sm text-red-600">{errors.username.message}</p>
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
            {...register("password")}
            placeholder="Contraseña"
            aria-invalid={!!errors.password}
            className="w-full bg-white border border-gray-300 focus:border-gray-900 text-gray-800 ph-login"
            autoComplete="current-password"
          />
          <div className="min-h-[20px]">
            {errors.password && (
              <p className="text-sm text-red-600">{errors.password.message}</p>
            )}
          </div>
        </div>

        <div>
          {message && <Message type={message.type} text={message.text} />}
        </div>

        <div className="">
          <div className="">
            {/* Botón de enviar */}
            <button
              type="submit"
              className="btn-login"
            >
              Iniciar sesión
            </button>
            <button type="button" className="btn-login" onClick={() => navigate('/register')}>
              Registrarse
            </button>
          </div>
        </div>
      </div>
    </form>
  );
};

export default LoginForm;
