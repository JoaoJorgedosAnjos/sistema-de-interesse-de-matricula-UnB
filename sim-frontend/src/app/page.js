'use client'; 

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';

export default function LoginPage() {
  const [matricula, setMatricula] = useState('');
  const [senha, setSenha] = useState('');
  const [erro, setErro] = useState('');

  const router = useRouter();

  const handleLogin = async (event) => {
    event.preventDefault(); 
    setErro(''); 

    try {
        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ matricula: matricula, senha: senha }),
        });

        if (!response.ok) {
            const errorData = await response.text();
            throw new Error(errorData || 'Matrícula ou senha inválida');
        }

        const data = await response.json();
        
        localStorage.setItem('authToken', data.token);

        alert('Login realizado com sucesso!');
        router.push('/dashboard'); 
    
    } catch (error) {
        setErro(error.message);
    }
  };

  return (
    <div>
      <h1>Página de Login</h1>
      <hr />

      <form onSubmit={handleLogin}>
        <div>
          <label htmlFor="matricula">Matrícula:</label>
          <br />
          <input
            type="text"
            id="matricula"
            value={matricula}
            onChange={(e) => setMatricula(e.target.value)}
            required
          />
        </div>

        <div style={{ marginTop: '1rem' }}>
          <label htmlFor="senha">Senha:</label>
          <br />
          <input
            type="password"
            id="senha"
            value={senha}
            onChange={(e) => setSenha(e.target.value)}
            required
          />
        </div>

        {erro && <p style={{ color: 'red', marginTop: '1rem' }}>{erro}</p>}

        <div style={{ marginTop: '1rem' }}>
          <button type="submit">
            Logar
          </button>
        </div>
      </form>

      <div style={{ marginTop: '1rem' }}>
        <Link href="/cadastro">
          <button type="button">Ir para Cadastro</button>
        </Link>
      </div>
    </div>
  );
}