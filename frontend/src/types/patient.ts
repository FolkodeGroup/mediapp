export interface Patient {
  id: number;
  firstName: string;
  lastName: string;
  dni: string;
  medicalRecordId: string;
  birthDate: string;
  gender: 'M' | 'F' | 'O';
  email: string;
}
