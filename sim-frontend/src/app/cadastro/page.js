'use client'; 

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import { fetchCursos } from '@/services/apiService'; 

export default function CadastroPage() {
  const [cursos, setCursos] = useState([]);
  
  const [formData, setFormData] = useState({
    matricula: '',
    cpf: '',
    nome_completo: '',
    data_nascimento: '',
    nacionalidade: 'Brasileira',
    semestre_ingresso: '',
    email_pessoal: '',
    email_institucional: '',
    senha: '',
    cod_curso: 0,
  });

  const [erro, setErro] = useState('');
  const router = useRouter();

  useEffect(() => {
    async function carregarCursos() {
      const cursosDaApi = await fetchCursos();
      if (cursosDaApi && cursosDaApi.length > 0) {
        setCursos(cursosDaApi);
        setFormData(prev => ({ ...prev, cod_curso: cursosDaApi[0].cod_curso }));
      }
    }
    carregarCursos();
  }, []); 

  const handleChange = (e) => {
    const { name, value } = e.target;
    const finalValue = name === 'cod_curso' ? parseInt(value, 10) : value;
    setFormData(prevState => ({
      ...prevState,
      [name]: finalValue,
    }));
  };

  // Função para lidar com a submissão do formulário
  const handleRegister = async (event) => {
    event.preventDefault();
    setErro('');

    try {
      const response = await fetch('http://localhost:8080/alunos', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        // Tenta ler a mensagem de erro da API
        const errorData = await response.text();
        throw new Error(errorData || 'Falha ao cadastrar aluno.');
      }

      alert('Aluno cadastrado com sucesso!');
      router.push('/'); // Redireciona para a página de login
      
    } catch (error) {
      setErro(error.message);
    }
  };

  return (
    <div>
      <h1>Página de Cadastro de Novo Aluno</h1>
      <hr />

      <form onSubmit={handleRegister}>
        {/* Usamos um único handler 'handleChange' para todos os inputs */}
        <div>
          <label htmlFor="nome_completo">Nome Completo:</label><br />
          <input type="text" id="nome_completo" name="nome_completo" value={formData.nome_completo} onChange={handleChange} required />
        </div>

        <div style={{ marginTop: '1rem' }}>
          <label htmlFor="matricula">Matrícula:</label><br />
          <input type="text" id="matricula" name="matricula" value={formData.matricula} onChange={handleChange} required />
        </div>
        
        <div style={{ marginTop: '1rem' }}>
          <label htmlFor="cpf">CPF:</label><br />
          <input type="text" id="cpf" name="cpf" value={formData.cpf} onChange={handleChange} required />
        </div>

        <div style={{ marginTop: '1rem' }}>
          <label htmlFor="data_nascimento">Data de Nascimento:</label><br />
          <input type="date" id="data_nascimento" name="data_nascimento" value={formData.data_nascimento} onChange={handleChange} required />
        </div>

        <div style={{ marginTop: '1rem' }}>
          <label htmlFor="semestre_ingresso">Semestre de Ingresso (ex: 2024.1):</label><br />
          <input type="text" id="semestre_ingresso" name="semestre_ingresso" value={formData.semestre_ingresso} onChange={handleChange} required />
        </div>

        <div style={{ marginTop: '1rem' }}>
          <label htmlFor="email_institucional">Email Institucional:</label><br />
          <input type="email" id="email_institucional" name="email_institucional" value={formData.email_institucional} onChange={handleChange} required />
        </div>

        <div style={{ marginTop: '1rem' }}>
          <label htmlFor="senha">Senha:</label><br />
          <input type="password" id="senha" name="senha" value={formData.senha} onChange={handleChange} required />
        </div>

        <div style={{ marginTop: '1rem' }}>
          <label htmlFor="cod_curso">Curso:</label><br />
          <select id="cod_curso" name="cod_curso" value={formData.cod_curso} onChange={handleChange} required>
            {cursos.map(curso => (
              <option key={curso.cod_curso} value={curso.cod_curso}>
                {curso.nome_curso}
              </option>
            ))}
          </select>
        </div>

        {erro && <p style={{ color: 'red', marginTop: '1rem' }}>Erro: {erro}</p>}

        <div style={{ marginTop: '1rem' }}>
          <button type="submit">Cadastrar</button>
        </div>
      </form>
      
      <div style={{ marginTop: '1rem' }}>
        <Link href="/">Voltar para Login</Link>
      </div>
    </div>
  );
}