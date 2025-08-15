import { useEffect, useState } from "react";

const Patients = () => {
  const [patients, setPatients] = useState([]);
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
    <div className="p-4">
      <h2 className="text-2xl font-bold mb-4">Pacientes</h2>
      {patients.length === 0 ? (
        <p>No hay pacientes registrados.</p>
      ) : (
        <table className="min-w-full bg-white dark:bg-gray-800 rounded shadow">
          <thead>
            <tr>
              <th className="px-4 py-2">Nombre</th>
              <th className="px-4 py-2">DNI</th>
              <th className="px-4 py-2">Email</th>
              <th className="px-4 py-2">Nacimiento</th>
              <th className="px-4 py-2">Género</th>
            </tr>
          </thead>
          <tbody>
            {patients.map((p: any) => (
              <tr key={p.id}>
                <td className="border px-4 py-2">{p.firstName} {p.lastName}</td>
                <td className="border px-4 py-2">{p.dni}</td>
                <td className="border px-4 py-2">{p.email}</td>
                <td className="border px-4 py-2">{p.birthDate}</td>
                <td className="border px-4 py-2">{p.gender}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default Patients;
