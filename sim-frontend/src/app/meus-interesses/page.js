'use client';

import { useState, useEffect } from 'react';
import { fetchMeusInteresses } from '@/services/apiService';

export default function MeusInteressesPage() {
  const [interesses, setInteresses] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function carregarInteresses() {
      const dados = await fetchMeusInteresses();
      setInteresses(dados);
      setLoading(false);
    }
    carregarInteresses();
  }, []);

  if (loading) {
    return <p>A carregar os seus interesses...</p>;
  }

  return (
    <div>
      <h1 className="text-3xl font-bold text-unb-blue mb-6">Meus Interesses Registados</h1>
      {interesses.length === 0 ? (
        <p>Você ainda não registou interesse em nenhuma turma.</p>
      ) : (
        <div className="space-y-4">
          {interesses.map(interesse => (
            <div key={interesse.id_registro_interesse} className="bg-white p-6 rounded-lg shadow-md border-l-4 border-unb-green">
              <div className="flex justify-between items-start">
                <div>
                  <h2 className="text-xl font-bold text-gray-800">{interesse.nome_materia}</h2>
                  <p className="text-sm text-gray-500">{interesse.cod_disciplina} - {interesse.semestre_oferta}</p>
                </div>
                <span className="bg-blue-100 text-unb-blue text-sm font-semibold px-3 py-1 rounded-full">
                  Prioridade: {interesse.prioridade_interesse}
                </span>
              </div>
              <div className="mt-4 text-gray-700">
                <p><strong>Horário:</strong> {interesse.horario}</p>
                <p><strong>Status:</strong> {interesse.status_interesse}</p>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}