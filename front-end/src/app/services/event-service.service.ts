import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { ConfigService } from './config.service';
import { map, filter, scan } from 'rxjs/operators';
import { Observable } from 'rxjs';
@Injectable()
export class EventServiceService {

  constructor(private http: Http, private config: ConfigService) { }

  getAllEvent() {
    return this.http.get('http://' + this.config.hostname + ':' + this.config.port + '/events')
      .pipe(map(data => data.json()));
  }

  createEvent(title: string, description: string, avaliable: string, speaker: string, date: string, round: string) {
    const payload = {
      title: title,
      description: description,
      avaliable: avaliable,
      speaker: speaker,
      date: date,
      round: round,
      user: []
    };
    this.http.post('http://' + this.config.hostname + ':' + this.config.port + '/events', payload);
  }

}
