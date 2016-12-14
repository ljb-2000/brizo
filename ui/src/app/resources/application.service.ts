import { Observable } from 'rxjs/Rx';
import { Component, Injectable } from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http';

import { Application } from './application.component'

@Injectable()
export class ApplicationService {

  private applicationsIndexUrl = '/api/v1/applications'
  private applicationsGetUrl = '/api/v1/applications/'

  constructor(private http: Http) { }

  getApplications(): Observable<Application[]> {
    return this.http.get(this.applicationsIndexUrl)
      .map((res: Response) => res.json())
      .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
  }

  getApplication(name: string): Observable<Application> {
    return this.http.get(this.applicationsGetUrl + `${name}`)
      .map((res: Response) => res.json())
      .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
  }
}
