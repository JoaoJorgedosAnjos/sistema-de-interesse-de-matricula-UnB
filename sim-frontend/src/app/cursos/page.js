import { fetchCursos } from "@/services/apiService";

export default async function CursosPage() {
  const cursos = await fetchCursos();

  return (
    <div>
      <h1 className="text-3xl font-bold text-unb-blue mb-6">Cursos Ofertados</h1>
      
      {cursos.length === 0 ? (
        <p>Nenhum curso encontrado ou falha ao carregar. Verifique se a API backend está em execução na porta 8080.</p>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {cursos.map((curso) => (
            <div key={curso.cod_curso} className="bg-white p-6 rounded-lg shadow-md hover:shadow-xl transition-shadow duration-300 border-l-4 border-unb-green">
              <h2 className="text-xl font-semibold text-gray-800">{curso.nome_curso}</h2>
              <p className="text-gray-600 mt-2">{curso.nome_departamento}</p>
              <span className="inline-block bg-blue-100 text-unb-blue text-sm font-semibold mt-4 px-3 py-1 rounded-full">
                {curso.nivel_curso}
              </span>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}