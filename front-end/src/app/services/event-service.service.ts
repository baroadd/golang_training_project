import { Injectable } from '@angular/core';
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

  createEvent(title: string, description: string, avaliable: any, speaker: string, date: string, round: string) {

    const headers = new HttpHeaders({
      'Content-Type': 'application/json'
    });
    const payload = {
      'title': title,
      'description': description,
      'avaliable': parseInt(avaliable, 10),
      'speaker': speaker,
      'date': date,
      'round': round,
      'user': []
    };
    return this.http.post('http://' + this.config.hostname + ':' + this.config.port + '/events',
      JSON.stringify(payload), { headers: headers });
  }

  deleteEvent(events: any) {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json'
    });
    const payload = {
      'id': events.id,
      'title': events.title,
      'description': events.description,
      'avaliable': parseInt(events.avaliable, 10),
      'speaker': events.speaker,
      'date': events.date,
      'round': events.round,
      'user': events.user
    };
    return this.http.post('http://' + this.config.hostname + ':' + this.config.port + '/events/delete', JSON.stringify(payload), { headers: headers });
  }

  updateEvent(events: any) {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json'
    });

    const payload = {
      'id': events.id,
      'title': events.title,
      'description': events.description,
      'avaliable': parseInt(events.avaliable, 10) > 0 ? parseInt(events.avaliable, 10) - 1 : 0,
      'speaker': events.speaker,
      'date': events.date,
      'round': events.round,
      'user': events.user
    };
    return this.http.put('http://' + this.config.hostname + ':' + this.config.port + '/events', JSON.stringify(payload), { headers: headers });
  }
}
