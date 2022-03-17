
export interface Datos {
  num1: number,
  num2: number,
  op: string
}

export interface Retorno {
  error: string | null,
  datos: number | null,
  status: boolean,
  tag: string
}
