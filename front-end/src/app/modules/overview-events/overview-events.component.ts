import { Component, OnInit } from '@angular/core';
import { EventServiceService } from '../../services/event-service.service';
import { Router } from '@angular/router';
@Component({
  selector: 'app-overview-events',
  templateUrl: './overview-events.component.html',
  styleUrls: ['./overview-events.component.css']
})
export class OverviewEventsComponent implements OnInit {
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

  onBtnClick() {
    this.router.navigateByUrl('/submit');
  }

}
