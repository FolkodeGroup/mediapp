import { useEffect, useState } from "react";

type Patient = {
  id: number;
  firstName: string;
  lastName: string;
  dni: string;
  medicalRecordId: string;
  birthDate: string;
  gender: string;
  email: string;
};

const Dashboard = () => {
  const [patients, setPatients] = useState<Patient[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch("/api/patients")
      .then((res) => res.json())
      .then((data) => setPatients(data.patients || []))
      .catch(() => setError("Error al cargar pacientes"))
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <div className="p-4">Cargando pacientes...</div>;
  if (error) return <div className="p-4 text-red-600">{error}</div>;

  return (
    <section className="">
      <div className="section-dashboard">
      {/* Sidebar */}
      <aside className="bg-white flex flex-col items-center py-6 shadow-md">
        {/* Logo / Nombre */}
        <h1 className="text-blue-600 text-4xl font-bold mb-10">MediApp</h1>

        {/* Avatar */}
        <div className="w-24 h-24 bg-blue-500 rounded-full flex items-center justify-center text-white text-3xl mb-10 align-self-center">
          <i className="fas fa-user"></i>
        </div>

        {/* Botones */}
        <nav className="flex flex-col gap-4 w-full px-4 align-self-center">
          <button className="bg-gray-400 text-white font-bold py-3 mb-5 rounded-md hover:bg-gray-500">
            Perfil
          </button>
          <button className="bg-gray-400 text-white font-bold py-3 mb-5 rounded-md hover:bg-gray-500">
            Historia
          </button>
          <button className="bg-gray-400 text-white font-bold py-3 mb-5 rounded-md hover:bg-gray-500">
            Consulta
          </button>
        </nav>
      </aside>

      {/* Contenido principal */}
      <main className="bg-blue-600 flex flex-col items-center justify-center text-white text-3xl font-bold p-6">
        <div className="flex gap-4 mt-6">
          {/* Tabla de pacientes */}
          <div className="bg-white dark:bg-gray-800 rounded shadow overflow-hidden">
            {patients.length === 0 ? (
              <p className="p-4">No hay pacientes registrados.</p>
            ) : (
              <table className="min-w-full">
                <thead className="bg-gray-50 dark:bg-gray-700">
                  <tr>
                    <th className="px-4 py-2 text-left">Nombre</th>
                    <th className="px-4 py-2 text-left">DNI</th>
                    <th className="px-4 py-2 text-left">Email</th>
                    <th className="px-4 py-2 text-left">Nacimiento</th>
                    <th className="px-4 py-2 text-left">Género</th>
                  </tr>
                </thead>
                <tbody>
                  {patients.map((patient) => (
                    <tr key={patient.id} className="border-t dark:border-gray-700">
                      <td className="px-4 py-2">{patient.firstName} {patient.lastName}</td>
                      <td className="px-4 py-2">{patient.dni}</td>
                      <td className="px-4 py-2">{patient.email}</td>
                      <td className="px-4 py-2">{patient.birthDate}</td>
                      <td className="px-4 py-2">{patient.gender}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            )}
          </div>
        </div>
      </main>
    </div>
    </section>
  );
};

export default Dashboard;