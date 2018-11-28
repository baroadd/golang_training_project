import { Component, OnInit } from '@angular/core';
import { EventServiceService } from '../../services/event-service.service';
import { Router } from '@angular/router';


@Component({
  selector: 'app-create-event',
  templateUrl: './create-event.component.html',
  styleUrls: ['./create-event.component.css']
})
export class CreateEventComponent implements OnInit {

  constructor(
    private api: EventServiceService,
    private router: Router,
  ) { }
  private title: any;
  private desc: any
  private limit: number;
  private speaker: any;
  private date: any;
  private round: any;

  ngOnInit() {
    this.title = "";
    this.desc = "";
    this.limit = 0;
    this.speaker ="";
    this.date = "";
    this.round = "";
  }


  createEvent() {
    this.api.createEvent(this.title, this.desc, this.limit, this.speaker, this.date, this.round).subscribe(data => { });
    this.router.navigateByUrl('/');
  }

}
