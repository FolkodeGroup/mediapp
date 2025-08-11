import '@testing-library/jest-dom';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import Button from './Button';

describe('Button', () => {
    test('renderiza el texto children correctamente', () => {
        render(<Button>Haz click</Button>);
        const buttonElement = screen.getByRole('button', { name: /haz click/i });
        expect(buttonElement).toBeInTheDocument();
    });

    test('llama a la función onClick cuando se hace clic en el botón', async () => {
        const mockOnClick = jest.fn();
        render(<Button onClick={mockOnClick}>Clickeame</Button>);
        const buttonElement = screen.getByRole('button', { name: /clickeame/i });
        await userEvent.click(buttonElement);
        expect(mockOnClick).toHaveBeenCalledTimes(1);
    });
});

// jest.config.cjs
module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'jsdom',
  setupFilesAfterEnv: ['<rootDir>/src/setupTests.ts'],
  moduleFileExtensions: ['ts', 'tsx', 'js', 'jsx'],
};