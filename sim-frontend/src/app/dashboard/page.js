import Link from 'next/link';

export default function DashboardPage() {
  return (
    <div>
      <h1 className="text-3xl font-bold text-unb-blue mb-6">Meu Painel</h1>
      <p className="mb-8">Bem-vindo! Selecione uma das opções abaixo para continuar.</p>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <Link href="/meus-interesses" className="bg-white p-6 rounded-lg shadow-md hover:shadow-xl transition-shadow">
          <h2 className="text-xl font-semibold text-unb-green mb-2">Meus Interesses</h2>
          <p className="text-gray-600">Ver e gerir as matérias em que você registou interesse.</p>
        </Link>

        <Link href="/turmas" className="bg-white p-6 rounded-lg shadow-md hover:shadow-xl transition-shadow">
          <h2 className="text-xl font-semibold text-unb-green mb-2">Ver Todas as Turmas</h2>
          <p className="text-gray-600">Explorar todas as turmas ofertadas no semestre.</p>
        </Link>

        <Link href="/perfil" className="bg-white p-6 rounded-lg shadow-md hover:shadow-xl transition-shadow">
          <h2 className="text-xl font-semibold text-unb-green mb-2">Meu Perfil</h2>
          <p className="text-gray-600">Ver e editar os seus dados de cadastro.</p>
        </Link>
      </div>
    </div>
  );
}