describe('Probando la página de Login', () => {

  it('testeando el url de la pagina', () => {

    cy.visit('/');


    cy.contains('h2', 'MediApp').should('be.visible');

    cy.contains('p', 'Iniciar sesión en tu cuenta').should('be.visible');

    cy.contains('button', 'Entrar').should('be.visible');
  });
});