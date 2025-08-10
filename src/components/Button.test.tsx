// src/components/Button.test.tsx
import { render, screen} from '@testing-library/react';
import Button from './Button';
import userEvent from '@testing-library/user-event';

// aca es en donde agregamos  varios teste al componente buttom, tambin podemos agregar test separados
describe('Button', () => {

    // (1): Renderizamos el contenido
    test('renderizando el texto children correctament', () => {
        render(<Button>Haz clic</Button>);
        const buttonElement = screen.getByRole('button', {name: /haz clic/i});
        expect(buttonElement).toBeInTheDocument();
    })

    // (2): manejando el evento onClick
    test('Llamamos a la función onClick cuando se hace clic en el botón', async () =>{
        const mockOnClick = jest.fn();
        render(<Button onClick={mockOnClick}>Clickeame</Button>);

        const buttonElement = screen.getByRole('button', {name: /clickeame/i});
        await userEvent.click(buttonElement);
        expect(mockOnClick).toHaveBeenCalledTimes(1);
    });

})