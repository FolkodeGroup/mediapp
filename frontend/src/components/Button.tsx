import React from 'react';

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  children: React.ReactNode;
  variant?: "primary" | "secondary" | "danger";
  size?: "sm" | "md" | "lg";
}

export default function Button({ children, variant = "primary", size = "md", ...props }: ButtonProps) {
  // Clases base del botón
  let baseClasses = "inline-flex items-center justify-center rounded font-medium transition focus:outline-none";

  // Variantes de estilo
  const variantClasses = {
    primary: "bg-blue-600 text-white hover:bg-blue-700 focus:ring-2 focus:ring-blue-500",
    secondary: "bg-gray-200 text-gray-800 hover:bg-gray-300 focus:ring-2 focus:ring-gray-500",
    danger: "bg-red-600 text-white hover:bg-red-700 focus:ring-2 focus:ring-red-500",
  };

  // Tamaños
  const sizeClasses = {
    sm: "px-3 py-1.5 text-sm",
    md: "px-4 py-2 text-base",
    lg: "px-6 py-3 text-lg",
  };

  // Combinar todas las clases
  const className = [
    baseClasses,
    variantClasses[variant],
    sizeClasses[size],
    props.className // para permitir personalización extra
  ].join(" ");

  return (
    <button className={className} {...props}>
      {children}
    </button>
  );
}
