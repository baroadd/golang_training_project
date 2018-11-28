import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { ConfigService } from './config.service';
import { map } from 'rxjs/operators';
import { HttpHeaders } from '@angular/common/http';
import { HttpClient } from '@angular/common/http';

@Injectable()
export class EventServiceService {

  constructor(private http: HttpClient, private config: ConfigService) { }

  getAllEvent() {
    return this.http.get('http://' + this.config.hostname + ':' + this.config.port + '/events')
      .pipe(map(data => data));
  }

  createEvent(title: string, description: string, avaliable: string, speaker: string, date: string, round: string) {

    const headers = new HttpHeaders({
      'Accept': 'text/plain',
      'Content-Type': 'application/json',
    });
    const payload = {
      'title': title,
      'description': description,
      'avaliable': avaliable,
      'speaker': speaker,
      'date': date,
      'round': round,
      'user': []
    };
    return this.http.post('http://' + this.config.hostname + ':' + this.config.port + '/events',
      JSON.stringify(payload), { headers: headers });
  }

}
