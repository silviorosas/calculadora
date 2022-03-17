import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
 
 
  { path: 'calculadora', loadChildren: () => import('./moduls/calculadora/calculadora.module').then(m => m.CalculadoraModule) },
  { path: '', redirectTo: 'calculadora', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
