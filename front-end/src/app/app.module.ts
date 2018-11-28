import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { OverviewEventsComponent } from './modules/overview-events/overview-events.component';
import { FormsModule } from '@angular/forms';
import { NavbarComponent } from './modules/navbar/navbar.component';

import { HttpClientModule } from '@angular/common/http';
import { EventServiceService } from './services/event-service.service';
import { ConfigService } from './services/config.service';
import { ButtonsModule } from 'ngx-bootstrap/buttons';
import { SubmitCourseComponent } from './modules/submit-course/submit-course.component';
import { CreateEventComponent } from './modules/create-event/create-event.component';
import { BsDatepickerModule } from 'ngx-bootstrap/datepicker';
import { DeleteEventComponent } from './modules/delete-event/delete-event.component';
const appRoutes: Routes = [
  { path: '', component: OverviewEventsComponent },
  { path: 'submit', component: SubmitCourseComponent },
  { path: 'create', component: CreateEventComponent },
  { path: 'delete', component: DeleteEventComponent },
];

@NgModule({
  declarations: [
    AppComponent,
    OverviewEventsComponent,
    NavbarComponent,
    SubmitCourseComponent,
    CreateEventComponent,
    DeleteEventComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    RouterModule.forRoot(appRoutes),
    HttpClientModule,
    ButtonsModule.forRoot(),
    BsDatepickerModule.forRoot(),
  ],
  providers: [
    EventServiceService,
    ConfigService],
  bootstrap: [AppComponent]
})
export class AppModule { }
