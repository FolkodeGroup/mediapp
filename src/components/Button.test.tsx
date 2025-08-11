import React from 'react';
import { render, screen } from '@testing-library/react';
import Button from './Button';

test('renders Button component', () => {
    render(<Button />);
    const buttonElement = screen.getByRole('button');
    expect(buttonElement).toBeInTheDocument();
});