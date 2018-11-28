import { Component, OnInit } from '@angular/core';
import { EventServiceService } from '../../services/event-service.service';
import { Router } from '@angular/router';
@Component({
  selector: 'app-delete-event',
  templateUrl: './delete-event.component.html',
  styleUrls: ['./delete-event.component.css']
})
export class DeleteEventComponent implements OnInit {
  listEvent: any;
  constructor(
    private api: EventServiceService,
    private router: Router,
  ) {
    this.listEvent = [];
  }

  ngOnInit() {
    this.api.getAllEvent().subscribe(data => {
      this.listEvent = data;
    });
  }

  onBtnClick(event: any) {
    this.api.deleteEvent(event).subscribe(data => {
      console.log(data)
    });
    this.api.getAllEvent().subscribe(data => {
      this.listEvent = data;
    });
  }

}
