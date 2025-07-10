export async function fetchCursos() {
  try {
    const response = await fetch('http://localhost:8080/cursos', { cache: 'no-store' });

    if (!response.ok) {
      throw new Error('Falha ao buscar dados da API');
    }

    const cursos = await response.json();
    return cursos;
  } catch (error) {
    console.error("Erro no serviço da API:", error);
    return []; 
  }
}

export async function fetchMeusInteresses() {
  const token = localStorage.getItem('authToken');
  if (!token) {
    return [];
  }

  try {
    const response = await fetch('http://localhost:8080/meus-interesses', {
      headers: {
        'Authorization': `Bearer ${token}`
      }
    });

    if (!response.ok) {
      throw new Error('Falha ao buscar interesses');
    }
    return await response.json();
  } catch (error) {
    console.error("Erro ao buscar interesses:", error);
    return [];
  }
}