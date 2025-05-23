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

import { PlayerAPI } from './player-api'
import { Player, CopyPlayerToPlayerAPI } from './player'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class PlayerService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  PlayerServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private playersUrl: string

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
    this.playersUrl = origin + '/api/github.com/fullstack-lang/gong/lib/tone/go/v1/players';
  }

  /** GET players from the server */
  // gets is more robust to refactoring
  gets(Name: string, frontRepo: FrontRepo): Observable<PlayerAPI[]> {
    return this.getPlayers(Name, frontRepo)
  }
  getPlayers(Name: string, frontRepo: FrontRepo): Observable<PlayerAPI[]> {

    let params = new HttpParams().set("Name", Name)

    return this.http.get<PlayerAPI[]>(this.playersUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<PlayerAPI[]>('getPlayers', []))
      );
  }

  /** GET player by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, Name: string, frontRepo: FrontRepo): Observable<PlayerAPI> {
    return this.getPlayer(id, Name, frontRepo)
  }
  getPlayer(id: number, Name: string, frontRepo: FrontRepo): Observable<PlayerAPI> {

    let params = new HttpParams().set("Name", Name)

    const url = `${this.playersUrl}/${id}`;
    return this.http.get<PlayerAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched player id=${id}`)),
      catchError(this.handleError<PlayerAPI>(`getPlayer id=${id}`))
    );
  }

  // postFront copy player to a version with encoded pointers and post to the back
  postFront(player: Player, Name: string): Observable<PlayerAPI> {
    let playerAPI = new PlayerAPI
    CopyPlayerToPlayerAPI(player, playerAPI)
    const id = typeof playerAPI === 'number' ? playerAPI : playerAPI.ID
    const url = `${this.playersUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<PlayerAPI>(url, playerAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<PlayerAPI>('postPlayer'))
    );
  }
  
  /** POST: add a new player to the server */
  post(playerdb: PlayerAPI, Name: string, frontRepo: FrontRepo): Observable<PlayerAPI> {
    return this.postPlayer(playerdb, Name, frontRepo)
  }
  postPlayer(playerdb: PlayerAPI, Name: string, frontRepo: FrontRepo): Observable<PlayerAPI> {

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<PlayerAPI>(this.playersUrl, playerdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted playerdb id=${playerdb.ID}`)
      }),
      catchError(this.handleError<PlayerAPI>('postPlayer'))
    );
  }

  /** DELETE: delete the playerdb from the server */
  delete(playerdb: PlayerAPI | number, Name: string): Observable<PlayerAPI> {
    return this.deletePlayer(playerdb, Name)
  }
  deletePlayer(playerdb: PlayerAPI | number, Name: string): Observable<PlayerAPI> {
    const id = typeof playerdb === 'number' ? playerdb : playerdb.ID;
    const url = `${this.playersUrl}/${id}`;

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<PlayerAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted playerdb id=${id}`)),
      catchError(this.handleError<PlayerAPI>('deletePlayer'))
    );
  }

  // updateFront copy player to a version with encoded pointers and update to the back
  updateFront(player: Player, Name: string): Observable<PlayerAPI> {
    let playerAPI = new PlayerAPI
    CopyPlayerToPlayerAPI(player, playerAPI)
    const id = typeof playerAPI === 'number' ? playerAPI : playerAPI.ID
    const url = `${this.playersUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<PlayerAPI>(url, playerAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<PlayerAPI>('updatePlayer'))
    );
  }

  /** PUT: update the playerdb on the server */
  update(playerdb: PlayerAPI, Name: string, frontRepo: FrontRepo): Observable<PlayerAPI> {
    return this.updatePlayer(playerdb, Name, frontRepo)
  }
  updatePlayer(playerdb: PlayerAPI, Name: string, frontRepo: FrontRepo): Observable<PlayerAPI> {
    const id = typeof playerdb === 'number' ? playerdb : playerdb.ID;
    const url = `${this.playersUrl}/${id}`;


    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<PlayerAPI>(url, playerdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated playerdb id=${playerdb.ID}`)
      }),
      catchError(this.handleError<PlayerAPI>('updatePlayer'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in PlayerService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("PlayerService" + error); // log to console instead

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
