import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { EventServiceService } from '../../services/event-service.service';
@Component({
  selector: 'app-submit-course',
  templateUrl: './submit-course.component.html',
  styleUrls: ['./submit-course.component.css']
})
export class SubmitCourseComponent implements OnInit {
   obj: any;
   id: any;
   fName: any;
   lName: any;
   tel: any;
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
    // this.obj.user.push(this.id + ' ' + this.fName + ' ' + this.lName + ' ' + this.tel);
    this.api.updateEvent(this.obj).subscribe(data => {
      console.log(data);
    });
  }

}
