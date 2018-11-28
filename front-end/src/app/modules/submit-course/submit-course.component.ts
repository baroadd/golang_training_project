import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from "@angular/router";
import { EventServiceService } from '../../services/event-service.service';
@Component({
  selector: 'app-submit-course',
  templateUrl: './submit-course.component.html',
  styleUrls: ['./submit-course.component.css']
})
export class SubmitCourseComponent implements OnInit {
  private obj: any;
  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private api: EventServiceService,
  ) { }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.obj = params;
    });
  }

  updateEvents() {
    this.api.updateEvent(this.obj).subscribe(data => {
      console.log(data)
    });
  }

}
