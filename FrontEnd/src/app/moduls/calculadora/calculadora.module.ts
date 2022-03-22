import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { CalculadoraRoutingModule } from './calculadora-routing.module';
import { CalculadoraComponent } from './calculadora.component';

import {InputTextModule} from 'primeng/inputtext';
import {ButtonModule} from 'primeng/button';
import {CardModule} from 'primeng/card';



@NgModule({
  declarations: [
    CalculadoraComponent
  ],
  imports: [
    CommonModule,
    CalculadoraRoutingModule,
    InputTextModule,
    ButtonModule,
    CardModule
  ]
})
export class CalculadoraModule { }
