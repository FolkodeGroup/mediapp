"use client";

import { forwardRef, type InputHTMLAttributes } from "react";

export interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  variant?: "primary" | "secondary";
}

const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ className = "", variant = "primary", ...props }, ref) => {
    const baseClasses = "px-4 py-2 rounded-lg border transition-colors bg";
    const variantClasses = {
      primary: "bg-surface border-default text-primary focus:border-primary",
      secondary: "bg-background border-subtle text-secondary focus:border-secondary",
    };

    return (
      <input
        ref={ref}
        className={`${baseClasses} ${variantClasses[variant]} ${className}`}
        placeholder="Nombre"
        {...props}
      />
    );
  }
);

Input.displayName = "Input"; // Para debugging en React DevTools

export default Input;