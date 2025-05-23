// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { FormFieldIntAPI } from './formfieldint-api'
import { FormFieldInt, CopyFormFieldIntToFormFieldIntAPI } from './formfieldint'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class FormFieldIntService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  FormFieldIntServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private formfieldintsUrl: string

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
    this.formfieldintsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/table/go/v1/formfieldints';
  }

  /** GET formfieldints from the server */
  // gets is more robust to refactoring
  gets(Name: string, frontRepo: FrontRepo): Observable<FormFieldIntAPI[]> {
    return this.getFormFieldInts(Name, frontRepo)
  }
  getFormFieldInts(Name: string, frontRepo: FrontRepo): Observable<FormFieldIntAPI[]> {

    let params = new HttpParams().set("Name", Name)

    return this.http.get<FormFieldIntAPI[]>(this.formfieldintsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<FormFieldIntAPI[]>('getFormFieldInts', []))
      );
  }

  /** GET formfieldint by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, Name: string, frontRepo: FrontRepo): Observable<FormFieldIntAPI> {
    return this.getFormFieldInt(id, Name, frontRepo)
  }
  getFormFieldInt(id: number, Name: string, frontRepo: FrontRepo): Observable<FormFieldIntAPI> {

    let params = new HttpParams().set("Name", Name)

    const url = `${this.formfieldintsUrl}/${id}`;
    return this.http.get<FormFieldIntAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched formfieldint id=${id}`)),
      catchError(this.handleError<FormFieldIntAPI>(`getFormFieldInt id=${id}`))
    );
  }

  // postFront copy formfieldint to a version with encoded pointers and post to the back
  postFront(formfieldint: FormFieldInt, Name: string): Observable<FormFieldIntAPI> {
    let formfieldintAPI = new FormFieldIntAPI
    CopyFormFieldIntToFormFieldIntAPI(formfieldint, formfieldintAPI)
    const id = typeof formfieldintAPI === 'number' ? formfieldintAPI : formfieldintAPI.ID
    const url = `${this.formfieldintsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormFieldIntAPI>(url, formfieldintAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<FormFieldIntAPI>('postFormFieldInt'))
    );
  }
  
  /** POST: add a new formfieldint to the server */
  post(formfieldintdb: FormFieldIntAPI, Name: string, frontRepo: FrontRepo): Observable<FormFieldIntAPI> {
    return this.postFormFieldInt(formfieldintdb, Name, frontRepo)
  }
  postFormFieldInt(formfieldintdb: FormFieldIntAPI, Name: string, frontRepo: FrontRepo): Observable<FormFieldIntAPI> {

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormFieldIntAPI>(this.formfieldintsUrl, formfieldintdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted formfieldintdb id=${formfieldintdb.ID}`)
      }),
      catchError(this.handleError<FormFieldIntAPI>('postFormFieldInt'))
    );
  }

  /** DELETE: delete the formfieldintdb from the server */
  delete(formfieldintdb: FormFieldIntAPI | number, Name: string): Observable<FormFieldIntAPI> {
    return this.deleteFormFieldInt(formfieldintdb, Name)
  }
  deleteFormFieldInt(formfieldintdb: FormFieldIntAPI | number, Name: string): Observable<FormFieldIntAPI> {
    const id = typeof formfieldintdb === 'number' ? formfieldintdb : formfieldintdb.ID;
    const url = `${this.formfieldintsUrl}/${id}`;

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<FormFieldIntAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted formfieldintdb id=${id}`)),
      catchError(this.handleError<FormFieldIntAPI>('deleteFormFieldInt'))
    );
  }

  // updateFront copy formfieldint to a version with encoded pointers and update to the back
  updateFront(formfieldint: FormFieldInt, Name: string): Observable<FormFieldIntAPI> {
    let formfieldintAPI = new FormFieldIntAPI
    CopyFormFieldIntToFormFieldIntAPI(formfieldint, formfieldintAPI)
    const id = typeof formfieldintAPI === 'number' ? formfieldintAPI : formfieldintAPI.ID
    const url = `${this.formfieldintsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<FormFieldIntAPI>(url, formfieldintAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<FormFieldIntAPI>('updateFormFieldInt'))
    );
  }

  /** PUT: update the formfieldintdb on the server */
  update(formfieldintdb: FormFieldIntAPI, Name: string, frontRepo: FrontRepo): Observable<FormFieldIntAPI> {
    return this.updateFormFieldInt(formfieldintdb, Name, frontRepo)
  }
  updateFormFieldInt(formfieldintdb: FormFieldIntAPI, Name: string, frontRepo: FrontRepo): Observable<FormFieldIntAPI> {
    const id = typeof formfieldintdb === 'number' ? formfieldintdb : formfieldintdb.ID;
    const url = `${this.formfieldintsUrl}/${id}`;


    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<FormFieldIntAPI>(url, formfieldintdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated formfieldintdb id=${formfieldintdb.ID}`)
      }),
      catchError(this.handleError<FormFieldIntAPI>('updateFormFieldInt'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in FormFieldIntService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("FormFieldIntService" + error); // log to console instead

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
