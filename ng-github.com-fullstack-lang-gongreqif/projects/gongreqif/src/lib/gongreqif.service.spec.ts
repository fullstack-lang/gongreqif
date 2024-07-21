import { TestBed } from '@angular/core/testing';

import { GongreqifService } from './gongreqif.service';

describe('GongreqifService', () => {
  let service: GongreqifService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongreqifService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
