// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { RectAnchoredTextAPI } from './rectanchoredtext-api'
import { RectAnchoredText, CopyRectAnchoredTextToRectAnchoredTextAPI } from './rectanchoredtext'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { AnimateAPI } from './animate-api'

@Injectable({
  providedIn: 'root'
})
export class RectAnchoredTextService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  RectAnchoredTextServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private rectanchoredtextsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.rectanchoredtextsUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/rectanchoredtexts';
  }

  /** GET rectanchoredtexts from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredTextAPI[]> {
    return this.getRectAnchoredTexts(GONG__StackPath, frontRepo)
  }
  getRectAnchoredTexts(GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredTextAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<RectAnchoredTextAPI[]>(this.rectanchoredtextsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<RectAnchoredTextAPI[]>('getRectAnchoredTexts', []))
      );
  }

  /** GET rectanchoredtext by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredTextAPI> {
    return this.getRectAnchoredText(id, GONG__StackPath, frontRepo)
  }
  getRectAnchoredText(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredTextAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.rectanchoredtextsUrl}/${id}`;
    return this.http.get<RectAnchoredTextAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched rectanchoredtext id=${id}`)),
      catchError(this.handleError<RectAnchoredTextAPI>(`getRectAnchoredText id=${id}`))
    );
  }

  // postFront copy rectanchoredtext to a version with encoded pointers and post to the back
  postFront(rectanchoredtext: RectAnchoredText, GONG__StackPath: string): Observable<RectAnchoredTextAPI> {
    let rectanchoredtextAPI = new RectAnchoredTextAPI
    CopyRectAnchoredTextToRectAnchoredTextAPI(rectanchoredtext, rectanchoredtextAPI)
    const id = typeof rectanchoredtextAPI === 'number' ? rectanchoredtextAPI : rectanchoredtextAPI.ID
    const url = `${this.rectanchoredtextsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<RectAnchoredTextAPI>(url, rectanchoredtextAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<RectAnchoredTextAPI>('postRectAnchoredText'))
    );
  }
  
  /** POST: add a new rectanchoredtext to the server */
  post(rectanchoredtextdb: RectAnchoredTextAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredTextAPI> {
    return this.postRectAnchoredText(rectanchoredtextdb, GONG__StackPath, frontRepo)
  }
  postRectAnchoredText(rectanchoredtextdb: RectAnchoredTextAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredTextAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<RectAnchoredTextAPI>(this.rectanchoredtextsUrl, rectanchoredtextdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted rectanchoredtextdb id=${rectanchoredtextdb.ID}`)
      }),
      catchError(this.handleError<RectAnchoredTextAPI>('postRectAnchoredText'))
    );
  }

  /** DELETE: delete the rectanchoredtextdb from the server */
  delete(rectanchoredtextdb: RectAnchoredTextAPI | number, GONG__StackPath: string): Observable<RectAnchoredTextAPI> {
    return this.deleteRectAnchoredText(rectanchoredtextdb, GONG__StackPath)
  }
  deleteRectAnchoredText(rectanchoredtextdb: RectAnchoredTextAPI | number, GONG__StackPath: string): Observable<RectAnchoredTextAPI> {
    const id = typeof rectanchoredtextdb === 'number' ? rectanchoredtextdb : rectanchoredtextdb.ID;
    const url = `${this.rectanchoredtextsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<RectAnchoredTextAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted rectanchoredtextdb id=${id}`)),
      catchError(this.handleError<RectAnchoredTextAPI>('deleteRectAnchoredText'))
    );
  }

  // updateFront copy rectanchoredtext to a version with encoded pointers and update to the back
  updateFront(rectanchoredtext: RectAnchoredText, GONG__StackPath: string): Observable<RectAnchoredTextAPI> {
    let rectanchoredtextAPI = new RectAnchoredTextAPI
    CopyRectAnchoredTextToRectAnchoredTextAPI(rectanchoredtext, rectanchoredtextAPI)
    const id = typeof rectanchoredtextAPI === 'number' ? rectanchoredtextAPI : rectanchoredtextAPI.ID
    const url = `${this.rectanchoredtextsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<RectAnchoredTextAPI>(url, rectanchoredtextAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<RectAnchoredTextAPI>('updateRectAnchoredText'))
    );
  }

  /** PUT: update the rectanchoredtextdb on the server */
  update(rectanchoredtextdb: RectAnchoredTextAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredTextAPI> {
    return this.updateRectAnchoredText(rectanchoredtextdb, GONG__StackPath, frontRepo)
  }
  updateRectAnchoredText(rectanchoredtextdb: RectAnchoredTextAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RectAnchoredTextAPI> {
    const id = typeof rectanchoredtextdb === 'number' ? rectanchoredtextdb : rectanchoredtextdb.ID;
    const url = `${this.rectanchoredtextsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<RectAnchoredTextAPI>(url, rectanchoredtextdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated rectanchoredtextdb id=${rectanchoredtextdb.ID}`)
      }),
      catchError(this.handleError<RectAnchoredTextAPI>('updateRectAnchoredText'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in RectAnchoredTextService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("RectAnchoredTextService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}