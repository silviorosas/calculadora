import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment'
import { Datos, Retorno } from './calculadora.models'

@Injectable({
  providedIn: 'root'
})
export class CalculadoraService {

  url: string = environment.apiURL;

  constructor(
    private http: HttpClient,
  ) { }


  getResultado(datos: Datos){

    return this.http.post<Retorno>(`${this.url}/calcular`, datos)

  }

}
