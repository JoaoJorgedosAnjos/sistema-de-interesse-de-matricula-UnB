import { Inter } from "next/font/google";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "SIM - Sistema de Interesse de Matrícula",
  description: "Projeto de Banco de Dados - UnB",
};

export default function RootLayout({ children }) {
  return (
    <html lang="pt-br">
      <body className={`${inter.className} bg-gray-100`}>
        {/* Usando a nossa cor customizada do Tailwind! */}
        <header className="bg-unb-blue text-white shadow-lg">
          <nav className="container mx-auto px-6 py-4">
            <h1 className="text-2xl font-bold">
              Sistema de Interesse de Matrícula
            </h1>
          </nav>
        </header>

        <main className="container mx-auto px-6 py-8">
          {children}
        </main>

        <footer className="bg-gray-200 text-center text-sm p-4 mt-8">
          <p>Projeto de Banco de Dados - UnB © 2025</p>
        </footer>
      </body>
    </html>
  );
}