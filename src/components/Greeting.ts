export function Greeting(name: string) {
  const p = document.createElement('p');
  p.textContent = `Hola, ${name}!`;
  return p;
}
