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

const Inicio = () => {
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
    <div className="p-6">
      <h2 className="text-2xl font-bold mb-4">Dashboard</h2>
      <p className="mb-4">Bienvenida al panel principal.</p>

      {/* Botón "Nuevo paciente" */}
      <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mb-4">
        Nuevo paciente
      </button>

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
  );
};

export default Inicio;