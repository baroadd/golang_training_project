import { Component, OnInit } from '@angular/core';
import { BsDatepickerModule } from 'ngx-bootstrap/datepicker';
import { EventServiceService } from '../../services/event-service.service';

@Component({
  selector: 'app-create-event',
  templateUrl: './create-event.component.html',
  styleUrls: ['./create-event.component.css']
})
export class CreateEventComponent implements OnInit {

  constructor(
    private api: EventServiceService,
  ) { }

  ngOnInit() {
  }

  createEvent(){
    this.api.createEvent("ddd","ddd","ddd","ddd","ddd","ddd");
  }

}
