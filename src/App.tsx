import { useState } from 'react';
import './App.css';
import Button from './components/Button';
import MainLayout from './components/layout/MainLayout';
import Login from './components/auth/Login';
function App() {
  const [count, setCount] = useState<number>(0);

  return (
    <MainLayout>
      <div className="card">
        <Button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </Button>
      </div>
      <Login />
    </MainLayout>
  );
}

export default App;
