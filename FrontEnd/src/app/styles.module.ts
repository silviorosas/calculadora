import { NgModule } from '@angular/core';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule} from 'primeng/inputtext';
import { CalendarModule} from 'primeng/calendar';

import { registerLocaleData } from '@angular/common';
import locales from '@angular/common/locales/es-AR';
registerLocaleData(locales);

@NgModule({
  exports: [
    //Comun
    ButtonModule,
    DialogModule,
    InputTextModule,
    CalendarModule
  ]
})

export class StylesModule { }
