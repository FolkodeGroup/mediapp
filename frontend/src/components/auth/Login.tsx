"use client";

import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import Input from "../../ui/Input";

// Esquema de validación para login
const loginSchema = z.object({
  name: z.string().min(1, "El usuario es obligatorio"),
  contrasena: z.string().min(1, "La contraseña es obligatoria"),
});

const LoginForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<z.infer<typeof loginSchema>>({
    resolver: zodResolver(loginSchema),
    mode: "onChange",
  });

  const onSubmit = (data: z.infer<typeof loginSchema>) => {
    console.log("Datos validados:", data);
    alert("Formulario enviado:\n" + JSON.stringify(data, null, 2));
  };

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="w-full max-w-md bg-[#transparent] rounded-xl shadow-lg p-8 flex flex-col justify-center h-full"
      autoComplete="off"
      noValidate
    >
      <div>
        {/* Usuario */}
        <div className="space-y-1 div-login">
          <label htmlFor="name" className="block text-sm font-semibold text-white mb-1">
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
          <label
            htmlFor="contrasena"
            className="block text-sm font-semibold text-white mb-1"
          >
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
        <div className="">
          {/* Botón de enviar */}
          <button
            type="submit"
            className="mt-6 w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition btn-login"
          >
            Iniciar sesión
          </button>
          {/* Registrarse  */}
          <button
            type="submit"
            className="mt-6 w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition btn-login"
          >
            Registrarse
          </button>
        </div>
      </div>
    </form>
  );
};

export default LoginForm;